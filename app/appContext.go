package app

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	oauthGh "golang.org/x/oauth2/github"
)

type appContext struct {
	OAuthConf *oauth2.Config
}

func NewAppContext() *appContext {
	return &appContext{
		OAuthConf: &oauth2.Config{
			ClientID:     viper.GetString("github-client-id"),
			ClientSecret: viper.GetString("github-client-secret"),
			Scopes:       []string{},
			Endpoint:     oauthGh.Endpoint,
		},
	}
}

func (ac *appContext) GitHub(token *oauth2.Token) *github.Client {
	ctx := context.Background()
	return github.NewClient(
		oauth2.NewClient(ctx, ac.OAuthConf.TokenSource(ctx, token)))
}
