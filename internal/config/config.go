package config

import (
	"github.com/lazyIoad/site/internal/logging"
	"github.com/philandstuff/dhall-golang/v6"
)

type NavLink struct {
	Title  string
	Target string
}

type SiteConfig struct {
	Port        uint16
	Origin      string
	Title       string
	Description string
	NavLinks    []*NavLink
}

func ReadSiteConfig() *SiteConfig {
	var sc SiteConfig
	err := dhall.UnmarshalFile("site.dhall", &sc)
	if err != nil {
		logging.PanicLogger.Fatalf("Failed to read site config %v", err)
	}
	return &sc
}
