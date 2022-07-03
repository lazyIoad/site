package server

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lazyIoad/site/internal/config"
	"github.com/lazyIoad/site/internal/content"
	"github.com/lazyIoad/site/internal/logging"
	"github.com/lazyIoad/site/internal/rss"
)

func StartServer(sc *config.SiteConfig) {
	r := httprouter.New()
	p, err := content.GetBlogPosts("blog")
	if err != nil {
		logging.PanicLogger.Fatalf("Failed to parse blog posts\n%v", err)
	}

	rss, err := rss.GetBlogFeeds(p, sc.Title, sc.Description, sc.Origin)
	if err != nil {
		logging.PanicLogger.Fatalf("Failed to generate blog rss feeds\n%v", err)
	}

	InitRoutes(r, p, sc, rss)
	logging.InfoLogger.Printf("Starting server on port %d", sc.Port)
	logging.PanicLogger.Fatal("Server returned error\n", http.ListenAndServe(fmt.Sprintf(":%d", sc.Port), r))
}
