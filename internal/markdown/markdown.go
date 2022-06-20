package markdown

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

type Doc struct {
	Path string
	Body string
	Meta map[string]interface{}
}

var md = goldmark.New(
	goldmark.WithExtensions(
		meta.Meta,
		highlighting.NewHighlighting(
			highlighting.WithStyle("dracula"),
		),
	),
)

func ParseDir(path string) ([]*Doc, error) {
	files, err := filepath.Glob(filepath.Join(path, "*.md"))
	if err != nil {
		return nil, err
	}

	docs := make([]*Doc, len(files))
	for i, f := range files {
		d, err := ParseFile(f)
		if err != nil {
			return nil, err
		}

		docs[i] = d
	}

	return docs, nil
}

func ParseFile(file string) (*Doc, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var buf strings.Builder
	ctx := parser.NewContext()
	err = md.Convert(b, &buf, parser.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	m := meta.Get(ctx)
	return &Doc{Body: buf.String(), Meta: m, Path: file}, nil
}
