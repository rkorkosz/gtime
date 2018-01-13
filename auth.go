package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/gob"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

func init() {
	gob.Register(&oauth2.Token{})
}

func login(ac *appContext) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := ac.SessionStore.Get(r, "session")
		if err != nil {
			http.Error(w, http.StatusText(403), 403)
			return
		}
		dat := make([]byte, 9)
		rand.Read(dat)
		state := base64.StdEncoding.EncodeToString(dat)
		session.Values["auth-state"] = state
		session.Save(r, w)
		url := ac.OAuthConf.AuthCodeURL(state, oauth2.AccessTypeOffline)
		http.Redirect(w, r, url, 302)
	})
}

func callback(ac *appContext) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := ac.SessionStore.Get(r, "session")
		if err != nil {
			http.Error(w, http.StatusText(403), 403)
			return
		}
		state := r.URL.Query().Get("state")
		if session.Values["auth-state"] != state {
			http.Error(w, http.StatusText(401), 401)
			return
		}
		delete(session.Values, "auth-state")
		ctx := context.Background()
		code := r.URL.Query().Get("code")
		tok, err := ac.OAuthConf.Exchange(ctx, code)
		if err != nil {
			log.Fatal(err)
			return
		}
		session.Values["token"] = tok
		session.Save(r, w)
		http.Redirect(w, r, "/", 302)
	})
}
