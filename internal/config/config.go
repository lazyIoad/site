package config

import (
	"github.com/lazyIoad/site/internal/logging"
	"github.com/philandstuff/dhall-golang/v6"
)

type SiteConfig struct {
	Port   uint16
	Domain string
}

func ReadSiteConfig() *SiteConfig {
	var sc SiteConfig
	err := dhall.UnmarshalFile("site.dhall", &sc)
	if err != nil {
		logging.PanicLogger.Fatalf("Failed to read site config %v", err)
	}
	return &sc
}
