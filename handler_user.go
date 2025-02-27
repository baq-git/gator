package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/baq-git/blogaggregator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}

	ctx := context.Background()
	name := cmd.Args[0]

	_, err := s.db.GetUserByName(ctx, name)
	if err != nil {
		fmt.Printf("%s username doesn't exist in database\n", name)
		os.Exit(1)
	}

	err = s.config.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	ctx := context.Background()
	name := cmd.Args[0]

	_, err := s.db.GetUserByName(ctx, name)
	if err == nil {
		fmt.Printf("%s already existed!\n", name)
		os.Exit(1)
	}

	user, err := s.db.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	})
	if err != nil {
		return err
	}

	err = s.config.SetUser(user.Name)
	if err != nil {
		return err
	}
	fmt.Printf("User %s created successfully\n", name)
	fmt.Printf("User details: Name=%s, CreatedAt=%s, UpdatedAt=%s\n",
		user.Name, user.CreatedAt, user.UpdatedAt)

	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	currentUserName := s.config.CurrentUserName

	ctx := context.Background()
	users, err := s.db.GetUsers(ctx)
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Name == currentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}

func handlerReset(s *state, cmd command) error {
	ctx := context.Background()
	err := s.db.DeleteAllUsers(ctx)
	if err != nil {
		return fmt.Errorf("Delete users fail, error: %v", err)
	}

	err = s.db.DeleteAllFeeds(ctx)
	if err != nil {
		return err
	}

	return nil
}
