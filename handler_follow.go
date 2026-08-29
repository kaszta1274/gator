package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kaszta1274/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			UserID:    user.ID,
			FeedID:    feed.ID,
		},
	)
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Printf("Successfully followed %s\n", cmd.Args[0])
	printCreateFeedFollow(feedFollow)
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get followed feeds: %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("No followed feeds found.")
		return nil
	}

	fmt.Printf("Found %d followed feeds:\n", len(feedFollows))
	for _, feedFollow := range feedFollows {
		printGetFeedFollow(feedFollow)
		fmt.Println()
	}
	return nil
}

func printCreateFeedFollow(feedFollow database.CreateFeedFollowRow) {
	fmt.Printf("* ID: %v\n", feedFollow.ID)
	fmt.Printf("* Created at: %v\n", feedFollow.CreatedAt)
	fmt.Printf("* Updated at: %v\n", feedFollow.UpdatedAt)
	fmt.Printf("* UserID: %v\n", feedFollow.UserID)
	fmt.Printf("* FeedID: %v\n", feedFollow.FeedID)
	fmt.Printf("* User name: %v\n", feedFollow.UserName)
	fmt.Printf("* Feed name: %v\n", feedFollow.FeedName)
}

func printGetFeedFollow(feedFollow database.GetFeedFollowsForUserRow) {
	fmt.Printf("* ID: %v\n", feedFollow.ID)
	fmt.Printf("* Created at: %v\n", feedFollow.CreatedAt)
	fmt.Printf("* Updated at: %v\n", feedFollow.UpdatedAt)
	fmt.Printf("* UserID: %v\n", feedFollow.UserID)
	fmt.Printf("* FeedID: %v\n", feedFollow.FeedID)
	fmt.Printf("* User name: %v\n", feedFollow.UserName)
	fmt.Printf("* Feed name: %v\n", feedFollow.FeedName)
}
