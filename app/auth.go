package app

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

func login(ac *appContext) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code == "" {
			url := ac.OAuthConf.AuthCodeURL("state", oauth2.AccessTypeOffline)
			http.Redirect(w, r, url, 302)
			return
		}
		ctx := context.Background()
		tok, err := ac.OAuthConf.Exchange(ctx, code)
		if err != nil {
			http.Error(w, http.StatusText(401), 401)
			return
		}
		json.NewEncoder(w).Encode(tok)
	})
}
