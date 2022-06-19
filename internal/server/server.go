package server

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lazyIoad/site/internal/config"
	"github.com/lazyIoad/site/internal/logging"
)

func StartServer(sc *config.SiteConfig) {
	r := httprouter.New()
	InitRoutes(r)
	logging.InfoLogger.Printf("Starting server on port %d", sc.Port)
	logging.PanicLogger.Fatal("Server returned error\n", http.ListenAndServe(fmt.Sprintf(":%d", sc.Port), r))
}
