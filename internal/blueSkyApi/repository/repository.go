package repository

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/admiralyeoj/anime-announcements/internal/logger"
	"github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/api/bsky"
	"github.com/bluesky-social/indigo/atproto/syntax"
	"github.com/bluesky-social/indigo/lex/util"
	"github.com/bluesky-social/indigo/xrpc"
)

// BlueSkyRepository defines the interface for BlySky repository actions
type BlueSkyRepository interface {
	CreateRecord(text string, images *[]string) error
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

	return &blueSkyRepository{
		Client:  client,
		Session: session,
	}
}

func (repo blueSkyRepository) CreateRecord(text string, images *[]string) error {

	post := bsky.FeedPost{
		Text:      text,
		CreatedAt: syntax.DatetimeNow().String(),
	}

	if images != nil {
		post.Embed = &bsky.FeedPost_Embed{
			EmbedImages: &bsky.EmbedImages{
				Images: createImageEmbeds(repo.Client, images),
			},
		}
	}

	// createPostEmbed(),

	ctx := context.Background()
	input := &atproto.RepoCreateRecord_Input{
		Repo:       repo.Session.Handle,
		Collection: "app.bsky.feed.post",
		Record:     &util.LexiconTypeDecoder{Val: &post},
	}

	response, err := atproto.RepoCreateRecord(ctx, repo.Client, input)
	fmt.Printf("Record created successfully: %+v\n", response)

	return err
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
