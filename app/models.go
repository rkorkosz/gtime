package app

import (
	"context"
	"fmt"
	"net/http"
	"net/mail"
	"strings"

	"golang.org/x/oauth2"
)

type User struct {
	Username    string        `json:"username"`
	Email       *mail.Address `json:"email"`
	SyncedRepos []string      `json:"synced_repos"`
	Token       *oauth2.Token `json:"-"`
}

func GetUser(ac *appContext, r *http.Request) (*User, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("No authorization header")
	}
	splitted := strings.Split(authHeader, " ")
	tok := &oauth2.Token{
		AccessToken: splitted[1],
		TokenType:   splitted[0],
	}
	cli := ac.GitHub(tok)
	user, _, err := cli.Users.Get(context.Background(), "")
	if err != nil {
		return nil, err
	}
	return &User{
		Username: user.GetLogin(),
		Token:    tok,
	}, nil
}
