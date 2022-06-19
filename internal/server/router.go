package server

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func InitRoutes(r *httprouter.Router) {
	r.GET("/", index)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello!\n")
}
