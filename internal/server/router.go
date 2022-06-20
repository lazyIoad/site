package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lazyIoad/site/internal/content"
	"github.com/lazyIoad/site/internal/logging"
	"github.com/lazyIoad/site/internal/template"
)

type navLink struct {
	Title string
	Path  string
}

type siteData struct {
	NavLinks []*navLink
}

type indexData struct {
	BlogPosts []*content.BlogPost
	SiteData  *siteData
}

func InitRoutes(r *httprouter.Router, p []*content.BlogPost) {
	r.GET("/", index(p))
}

func index(p []*content.BlogPost) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		err := template.Render(w, "index", indexData{BlogPosts: p})
		if err != nil {
			logging.ErrorLogger.Println(err)
			http.Error(w, "Error rendering page", http.StatusInternalServerError)
		}
	}
}
