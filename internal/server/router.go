package server

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lazyIoad/site/internal/config"
	"github.com/lazyIoad/site/internal/content"
	"github.com/lazyIoad/site/internal/logging"
	"github.com/lazyIoad/site/internal/rss"
	"github.com/lazyIoad/site/internal/template"
)

type siteData struct {
	NavLinks []*config.NavLink
}

type indexData struct {
	BlogPosts []*content.BlogPost
	SiteData  *siteData
}

type blogPostData struct {
	Post     *content.BlogPost
	SiteData *siteData
}

type tagData struct {
	BlogPosts []*content.BlogPost
	Tag       string
	SiteData *siteData
}

func InitRoutes(r *httprouter.Router, ps []*content.BlogPost, c *config.SiteConfig, fs *rss.Feeds) {
	r.ServeFiles("/static/*filepath", http.Dir("build/static/"))
	r.GET("/", index(ps, c.NavLinks))
	r.GET("/blog/:slug", blog(ps, c.NavLinks))
	r.GET("/feeds/blog/:type", blogRss(fs))
	r.GET("/tags/:tag", tags(ps, c.NavLinks))
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

func blog(p []*content.BlogPost, n []*config.NavLink) httprouter.Handle {
	s := make(map[string]*content.BlogPost, len(p))
	for _, po := range p {
		if val, ok := s[po.Slug]; ok {
			logging.PanicLogger.Fatalf("posts %s and %s have conflicting slugs", po.Doc.Path, val.Doc.Path)
		}

		s[po.Slug] = po
	}

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		rp, ok := s[ps.ByName("slug")]
		if !ok {
			http.Error(w, "Page not found", http.StatusNotFound)
			return
		}

		err := template.Render(w, "post", blogPostData{Post: rp, SiteData: &siteData{NavLinks: n}})
		if err != nil {
			logging.ErrorLogger.Println(err)
			http.Error(w, "Error rendering page", http.StatusInternalServerError)
		}
	}
}

func blogRss(fs *rss.Feeds) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ft := ps.ByName("type")
		switch ft {
		case "atom":
			fmt.Fprint(w, fs.Atom)
		case "rss":
			fmt.Fprint(w, fs.Rss)
		case "json":
			fmt.Fprint(w, fs.Json)
		default:
			http.Error(w, "Feed type not supported", http.StatusNotFound)
		}
	}
}

func tags(p []*content.BlogPost, n []*config.NavLink) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		tag := ps.ByName("tag")
		tps := content.FilterBlogPostsByTag(p, tag)
		err := template.Render(w, "tag", tagData{BlogPosts: tps, Tag: tag, SiteData: &siteData{NavLinks: n}})
		if err != nil {
			logging.ErrorLogger.Println(err)
			http.Error(w, "Error rendering page", http.StatusInternalServerError)
		}
	}
}
