package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kaszta1274/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get user: %w", err)
	}

	feedName := cmd.Args[0]
	feedUrl := cmd.Args[1]

	feed, err := s.db.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      feedName,
			Url:       feedUrl,
			UserID:    user.ID,
		},
	)
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	fmt.Printf("Feed added successfully: %+v\n", feed)
	return nil
}

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	feeds, err := s.db.GetFeedsWithUserName(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	fmt.Printf("Found %d feeds:\n", len(feeds))
	for _, feed := range feeds {
		fmt.Printf("* ID: %v\n", feed.ID)
		fmt.Printf("* Name: %v\n", feed.Name)
		fmt.Printf("* URL: %v\n", feed.Url)
		fmt.Printf("* Created by: %v\n", feed.CreatedBy)
		fmt.Println()
	}
	return nil
}
