package main

import (
	"context"
	"fmt"

	"github.com/kaszta1274/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	feedUrl := cmd.Args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	err = s.db.DeleteFeedFollowByUserIdAndFeedId(
		context.Background(),
		database.DeleteFeedFollowByUserIdAndFeedIdParams{
			UserID: user.ID,
			FeedID: feed.ID,
		},
	)
	if err != nil {
		return fmt.Errorf("couldn't delete feed follow: %w", err)
	}

	fmt.Printf("Successfully unfollowed %s!\n", feed.Name)
	return nil
}
