package main

import (
	"os"

	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type appContext struct {
	OAuthConf    *oauth2.Config
	SessionStore *sessions.CookieStore
}

func NewAppContext() *appContext {
	return &appContext{
		OAuthConf: &oauth2.Config{
			ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
			ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
			Scopes:       []string{},
			Endpoint:     github.Endpoint,
		},
		SessionStore: sessions.NewCookieStore([]byte(os.Getenv("SECRET_KEY"))),
	}
}
