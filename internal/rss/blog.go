package rss

import (
	"fmt"
	"time"

	"github.com/gorilla/feeds"
	"github.com/lazyIoad/site/internal/content"
)

type Feeds struct {
	Atom string
	Rss  string
	Json string
}

func GetBlogFeeds(ps []*content.BlogPost, title string, desc string, origin string) (*Feeds, error) {
	f := &feeds.Feed{
		Title:       title,
		Link:        &feeds.Link{Href: origin},
		Description: desc,
		Created:     time.Now(),
	}

	fis := make([]*feeds.Item, len(ps))
	for i, p := range ps {
		fi := &feeds.Item{
			Title:   p.Title,
			Created: p.PublishedAt,
			Link:    &feeds.Link{Href: fmt.Sprintf("%s/%s", origin, p.Slug)},
			Content: p.Doc.Body,
		}

		fis[i] = fi
	}

	f.Items = fis
	return buildFeeds(f)
}

func buildFeeds(f *feeds.Feed) (*Feeds, error) {
	atom, err := f.ToAtom()
	if err != nil {
		return nil, err
	}

	rss, err := f.ToRss()
	if err != nil {
		return nil, err
	}

	json, err := f.ToJSON()
	if err != nil {
		return nil, err
	}

	rssFeeds := &Feeds{Atom: atom, Rss: rss, Json: json}
	return rssFeeds, nil
}
