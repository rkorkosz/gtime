package app

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(ac *appContext) http.Handler {
	r := mux.NewRouter()
	r.Handle("/login/", login(ac))
	r.Handle("/repos/", repos(ac))

	ch := handlers.CompressHandler(jsonResponse(r))
	lh := handlers.LoggingHandler(os.Stdout, ch)
	return lh
}
