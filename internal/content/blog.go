package content

import (
	"fmt"
	"sort"
	"time"

	"github.com/lazyIoad/site/internal/markdown"
)

var REQUIRED_METADATA = [...]string{"title", "publishedAt", "slug"}
var ERROR_MSG_FAILED_CAST = "failed to cast metadata %s from file %s"
var METADATA_DATE_FORMAT = "2006-01-02"

type BlogPost struct {
	Title       string
	PublishedAt time.Time
	UpdatedAt   time.Time
	Slug        string
	Doc         *markdown.Doc
}

func GetBlogPosts(path string) ([]*BlogPost, error) {
	docs, err := markdown.ParseDir(path)
	if err != nil {
		return nil, err
	}

	posts := make([]*BlogPost, len(docs))
	for i, d := range docs {
		p, err := buildPost(d)
		if err != nil {
			return nil, err
		}

		posts[i] = p
	}

	sort.Slice(posts, func(i, j int) bool { return posts[i].PublishedAt.Unix() > posts[j].PublishedAt.Unix() })
	return posts, nil
}

func buildPost(doc *markdown.Doc) (*BlogPost, error) {
	for _, m := range REQUIRED_METADATA {
		if _, ok := doc.Meta[m]; !ok {
			return nil, fmt.Errorf("file %s is missing meta field %s", doc.Path, m)
		}
	}

	title, ok := doc.Meta["title"].(string)
	if !ok {
		return nil, fmt.Errorf(ERROR_MSG_FAILED_CAST, "title", doc.Path)
	}

	pubRaw, ok := doc.Meta["publishedAt"].(string)
	if !ok {
		return nil, fmt.Errorf(ERROR_MSG_FAILED_CAST, "publishedAt", doc.Path)
	}

	pub, err := time.Parse(METADATA_DATE_FORMAT, pubRaw)
	if err != nil {
		return nil, fmt.Errorf(ERROR_MSG_FAILED_CAST, "publishedAt", doc.Path)
	}

	var upd time.Time
	updRaw, ok := doc.Meta["publishedAt"].(string)
	if ok {
		upd, err = time.Parse(METADATA_DATE_FORMAT, updRaw)
		if err != nil {
			return nil, fmt.Errorf(ERROR_MSG_FAILED_CAST, "updatedAt", doc.Path)
		}
	}

	slug, ok := doc.Meta["slug"].(string)
	if !ok {
		return nil, fmt.Errorf(ERROR_MSG_FAILED_CAST, "slug", doc.Path)
	}

	return &BlogPost{Title: title, PublishedAt: pub, UpdatedAt: upd, Slug: slug}, nil
}
