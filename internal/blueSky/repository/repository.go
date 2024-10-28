package repository

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/admiralyeoj/animanager/internal/logger"
	"github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/api/bsky"
	"github.com/bluesky-social/indigo/atproto/syntax"
	"github.com/bluesky-social/indigo/lex/util"
	"github.com/bluesky-social/indigo/xrpc"
)

// BlueSkyRepository defines the interface for BlySky repository actions
type BlueSkyRepository interface {
	CreateRecord(text *string, images *[]string) (*string, error)
}

// blueSkyRepository is a concrete implementation of BlueSkyRepository
type blueSkyRepository struct {
	Client  *xrpc.Client
	Session *atproto.ServerCreateSession_Output
}

// NewAniListRepository creates and returns a new instance of aniListRepository
func NewBlueSkyRepositories() BlueSkyRepository {
	blueskyHandle := os.Getenv("BLUESKY_HANDLE")
	blueskyPass := os.Getenv("BLUESKY_APP_PASSWORD")

	client := &xrpc.Client{
		Host: "https://bsky.social",
	}

	ctx := context.Background()
	input := &atproto.ServerCreateSession_Input{
		Identifier: blueskyHandle,
		Password:   blueskyPass,
	}

	session, err := atproto.ServerCreateSession(ctx, client, input)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if !(*session.Active) {
		logger.Logger.DPanicf("Account " + blueskyHandle + " has been " + *session.Status + "!")
	}

	auth := &xrpc.AuthInfo{
		AccessJwt:  session.AccessJwt,
		RefreshJwt: session.RefreshJwt,
		Handle:     session.Handle,
		Did:        session.Did,
	}
	client.Auth = auth

	return blueSkyRepository{
		Client:  client,
		Session: session,
	}
}

func (repo blueSkyRepository) CreateRecord(text *string, images *[]string) (*string, error) {

	var linksTo []LinkInfo
	text, linksTo = extractLinks(*text)

	var facets []*bsky.RichtextFacet
	if len(linksTo) > 0 {
		for _, link := range linksTo {
			facet := &bsky.RichtextFacet{
				Features: []*bsky.RichtextFacet_Features_Elem{
					{
						RichtextFacet_Link: &bsky.RichtextFacet_Link{
							Uri: link.URL,
						},
					},
				},
				Index: &bsky.RichtextFacet_ByteSlice{
					ByteStart: int64(link.Indices[0]),
					ByteEnd:   int64(link.Indices[1]),
				},
			}

			facets = append(facets, facet)
		}
	}

	post := bsky.FeedPost{
		Text:      *text,
		CreatedAt: syntax.DatetimeNow().String(),
		Facets:    facets,
	}

	if images != nil {
		embeds := createImageEmbeds(repo.Client, images)

		if len(embeds) > 0 {
			post.Embed = &bsky.FeedPost_Embed{
				EmbedImages: &bsky.EmbedImages{
					Images: embeds,
				},
			}
		}
	}

	ctx := context.Background()
	input := &atproto.RepoCreateRecord_Input{
		Repo:       repo.Session.Handle,
		Collection: "app.bsky.feed.post",
		Record:     &util.LexiconTypeDecoder{Val: &post},
	}

	response, err := atproto.RepoCreateRecord(ctx, repo.Client, input)
	if err != nil {
		return nil, err
	}

	return extractSegment(response.Uri)
}

// createImageEmbeds downloads images, uploads them as blobs, and creates embedded images for a post.
func createImageEmbeds(client *xrpc.Client, images *[]string) []*bsky.EmbedImages_Image {
	var embedImages []*bsky.EmbedImages_Image // Slice to hold embedded images

	// Loop through each image URL
	for _, imageURL := range *images {
		// Send an HTTP GET request to retrieve the image
		response, err := http.Get(imageURL)
		if err != nil {
			logger.Logger.Panicln("failed to fetch image:", err)
			continue // Skip to the next image on error
		}

		// Ensure the response body is closed after usage
		imgReader := response.Body
		defer imgReader.Close()

		// Check if the request was successful (HTTP 200)
		if response.StatusCode != http.StatusOK {
			logger.Logger.Panicln("failed to fetch image, status code:", response.StatusCode)
			continue // Skip to the next image on error
		}

		// Upload the image blob to BlueSky and get the CID
		ctx := context.Background()
		blobCID, err := atproto.RepoUploadBlob(ctx, client, imgReader)
		if err != nil {
			logger.Logger.Panicln("Error uploading blob:", err)
			continue // Skip to the next image on error
		}

		// Create the embedded image for this blob and append it to the slice
		embedImages = append(embedImages, &bsky.EmbedImages_Image{
			Image: blobCID.Blob, // Use the blob returned by the upload
		})
	}

	return embedImages // Return the slice of embedded images
}

type LinkInfo struct {
	URL     string
	Indices [2]int // Start and end indices of the filtered text
}

// extractLinks processes the input string to remove <a> tags while extracting the URL and keeping the text.
func extractLinks(input string) (*string, []LinkInfo) {
	// Define a regex pattern to find <a href='...'>...</a>
	pattern := `<a[^>]*href=['"]([^'"]+)['"][^>]*>(.*?)<\/a>`
	re := regexp.MustCompile(pattern)

	// Create a slice to hold the LinkInfo structs
	var links []LinkInfo

	// Track cumulative offset adjustment as we remove <a> tags
	offset := 0

	// Iterate over all matches
	for _, match := range re.FindAllStringSubmatchIndex(input, -1) {
		// Ensure match has expected submatches to avoid slice out-of-range errors
		if len(match) < 6 {
			continue
		}

		startTagIndex := match[0] - offset  // Start index of <a> adjusted by offset
		endTagIndex := match[1] - offset    // End index of </a> adjusted by offset
		urlStartIndex := match[2] - offset  // Start index of the URL adjusted by offset
		urlEndIndex := match[3] - offset    // End index of the URL adjusted by offset
		textStartIndex := match[4] - offset // Start index of the text adjusted by offset
		textEndIndex := match[5] - offset   // End index of the text adjusted by offset

		// Calculate positions relative to the cleaned text
		newStartIndex := startTagIndex
		textLength := textEndIndex - textStartIndex
		newEndIndex := newStartIndex + textLength

		// Append the URL and new indices to the slice
		links = append(links, LinkInfo{
			URL:     input[urlStartIndex:urlEndIndex], // Extract the URL
			Indices: [2]int{newStartIndex, newEndIndex},
		})

		// Remove the <a>...</a> tags from the input string while keeping the text in between
		input = input[:newStartIndex] + input[textStartIndex:textEndIndex] + input[endTagIndex:]

		// Update the cumulative offset based on the difference in length after tag removal
		offset += (endTagIndex - startTagIndex) - textLength
	}

	return &input, links
}

func extractSegment(input string) (*string, error) {
	// Find the last index of the '/' character
	lastSlashIndex := strings.LastIndex(input, "/")
	if lastSlashIndex == -1 {
		return nil, errors.New("no segment found")
	}

	// Extract the substring after the last '/' character
	lastSegment := input[lastSlashIndex+1:]
	return &lastSegment, nil
}
