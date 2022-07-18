package models

import (
	"context"
	"html"
	"log"
	"strings"

	"guessr.net/pkg/database"

	"golang.org/x/crypto/bcrypt"
)

func (u *User) SaveUser() (*User, error) {
	var err error
	queries := New(database.DB)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Password = string(hashedPassword)
	params := CreateUserParams{
		Username: u.Username,
		Password: u.Password,
	}

	user, err := queries.CreateUser(context.Background(), params)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByUsername(username string) (int64, error) {
	queries := New(database.DB)
	user, err := queries.GetUserByUsername(context.Background(), username)
	if err != nil {
		log.Println(err)
	}
	return user.ID, nil
}

func GetUserByID(id int64) (User, error) {
	queries := New(database.DB)
	user, err := queries.GetUser(context.Background(), id)
	if err != nil {
		log.Println(err)
	}

	return user, nil
}
