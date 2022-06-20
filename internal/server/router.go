package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lazyIoad/site/internal/config"
	"github.com/lazyIoad/site/internal/content"
	"github.com/lazyIoad/site/internal/logging"
	"github.com/lazyIoad/site/internal/template"
)

type siteData struct {
	NavLinks []*config.NavLink
}

type indexData struct {
	BlogPosts []*content.BlogPost
	SiteData  *siteData
}

func InitRoutes(r *httprouter.Router, p []*content.BlogPost, c *config.SiteConfig) {
	r.ServeFiles("/static/*filepath", http.Dir("build/static/"))
	r.GET("/", index(p, c.NavLinks))
}

func index(p []*content.BlogPost, n []*config.NavLink) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		err := template.Render(w, "index", indexData{BlogPosts: p, SiteData: &siteData{NavLinks: n}})
		if err != nil {
			logging.ErrorLogger.Println(err)
			http.Error(w, "Error rendering page", http.StatusInternalServerError)
		}
	}
}
