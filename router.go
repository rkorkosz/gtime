package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(ac *appContext) http.Handler {
	r := mux.NewRouter()
	r.Handle("/login/", login(ac))
	r.Handle("/login/callback/", callback(ac))
	lh := handlers.LoggingHandler(os.Stdout, r)
	return lh
}
