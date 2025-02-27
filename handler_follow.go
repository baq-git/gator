package main

import (
	"context"
	"fmt"
	"os"

	"github.com/baq-git/blogaggregator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		fmt.Print("follow require 1 url argument\n ")
		os.Exit(1)
	}

	fmt.Print(cmd)

	ctx := context.Background()
	url := cmd.Args[0]

	feed, err := s.db.GetFeedByUrl(ctx, url)

	feedFollows, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("%s", feedFollows.FeedName)
	fmt.Printf("%s", feedFollows.UserName)

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.Args) > 0 {
		fmt.Print("following command does not require any arguments\n ")
		os.Exit(1)
	}

	ctx := context.Background()

	feedFollows, err := s.db.GetFeedsByUser(ctx, user.ID)
	if err != nil {
		return err
	}

	for _, ff := range feedFollows {
		fmt.Println(ff.FeedName)
	}

	return nil
}

func handlerUnFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		fmt.Print("following command does not require any arguments\n ")
		os.Exit(1)
	}

	ctx := context.Background()

	feed, err := s.db.GetFeedByUrl(ctx, cmd.Args[0])
	if err != nil {
		return err
	}

	err = s.db.DeleteFeedFollow(ctx, database.DeleteFeedFollowParams{UserID: user.ID, FeedID: feed.ID})
	if err != nil {
		return err
	}

	return nil
}
