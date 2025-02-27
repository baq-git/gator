package main

import (
	"context"
	"fmt"
	"os"

	"github.com/baq-git/blogaggregator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		fmt.Print("Addfeed require 2 arguments\n")
		os.Exit(1)
	}

	ctx := context.Background()
	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := s.db.CreateFeed(ctx, database.CreateFeedParams{
		Name:   name,
		Url:    url,
		UserID: user.ID,
	})
	if err != nil {
		return err
	}

	fmt.Print(feed.ID)

	_, err = s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println(feed)

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.Args) > 0 {
		fmt.Print("Expect 0 argument but get %i", len(cmd.Args))
		os.Exit(1)
	}
	ctx := context.Background()

	feeds, err := s.db.ListAllFeeds(ctx)
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Println(feed)
	}

	return nil
}
