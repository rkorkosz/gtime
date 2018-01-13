package main

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

func loginRequired(ac *appContext, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := ac.SessionStore.Get(r, "session")
		if err != nil {
			http.Error(w, http.StatusText(403), 403)
			return
		}
		tok := session.Values["token"].(*oauth2.Token)
		if tok.Valid() {
			next.ServeHTTP(w, r)
			return
		}
		ac.OAuthConf.Client(context.Background(), tok)
		session.Values["token"] = tok
		session.Save(r, w)
		next.ServeHTTP(w, r)
	})
}
