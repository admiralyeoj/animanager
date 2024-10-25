package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/admiralyeoj/anime-announcements/internal/logger"
	"github.com/bluesky-social/indigo/api/atproto"
	appbsky "github.com/bluesky-social/indigo/api/bsky"
	"github.com/bluesky-social/indigo/atproto/syntax"
	"github.com/bluesky-social/indigo/lex/util"
	"github.com/bluesky-social/indigo/xrpc"
)

// BlueSkyRepository defines the interface for BlySky repository actions
type BlueSkyRepository interface {
	CreateRecord(text string) error
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

func (repo blueSkyRepository) CreateRecord(text string) error {

	post := appbsky.FeedPost{
		Text:      text,
		CreatedAt: syntax.DatetimeNow().String(),
	}

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
