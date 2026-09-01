package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/kaszta1274/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 0 && len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s [limit] - limit defaults to 2", cmd.Name)
	}

	limit := 2
	if len(cmd.Args) == 1 {
		var err error
		limit, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("couldn't convert limit to int: %w", err)
		}
		if limit < 0 {
			return fmt.Errorf("limit should be greater than 0")
		}
	}

	posts, err := s.db.GetPostsForUser(
		context.Background(),
		database.GetPostsForUserParams{
			UserID: user.ID,
			Limit:  int32(limit),
		},
	)
	if err != nil {
		return fmt.Errorf("couldn't get posts for user: %w", err)
	}

	if len(posts) == 0 {
		fmt.Printf("No posts found for %s\n", user.Name)
		return nil
	}

	fmt.Printf("Found %d posts for %s:\n", len(posts), user.Name)
	for _, post := range posts {
		printPost(post)
		fmt.Println()
	}
	return nil
}

func printPost(post database.Post) {
	fmt.Printf("* Title: %s\n", post.Title)
	fmt.Printf("* URL: %s\n", post.Url)
	fmt.Printf("* Published at: %v\n", post.PublishedAt.Time)
	fmt.Printf("* Description: %s\n", post.Description.String)
}
