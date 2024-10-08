package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/miggoxyz/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	user, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("error finding user %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("could not set user: %w", err)
	}

	fmt.Println("User switched successfully")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: name,})
	if err != nil {
		return fmt.Errorf("could not create user: %w", err) 
	}
	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("could not set user: %w", err)
	}
	
	fmt.Printf("User was created successfully: %v\n", user)
	return nil
}

func handlerDeleteAll(s *state, cmd command) error {
	err := s.db.DelUsers(context.Background())
	if err != nil {
		return fmt.Errorf("could not delete users: %w", err)
	}
	fmt.Println("Deletion of all users successful.")
	
	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error getting users: %w", err)	
	}

	for _, user := range users {
		fmt.Printf("* %s ", user.Name)
		if user.Name == s.cfg.User {
			fmt.Printf("(current)")
		}
		fmt.Printf("\n")
	}

	return nil
}