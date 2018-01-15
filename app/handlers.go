package app

import (
	"context"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func repos(ac *appContext) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := GetUser(ac, r)
		if err != nil {
			http.Error(w, http.StatusText(401), 401)
			return
		}
		cli := ac.GitHub(user.Token)
		ctx := context.Background()
		repos, _, err := cli.Repositories.List(ctx, user.Username, nil)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		json.NewEncoder(w).Encode(repos)
	})
}

func syncRepo(ac *appContext) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}
