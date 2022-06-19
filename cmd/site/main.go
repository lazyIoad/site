package main

import (
	"github.com/lazyIoad/site/internal/config"
	"github.com/lazyIoad/site/internal/logging"
	"github.com/lazyIoad/site/internal/server"
)

func main() {
	logging.InitLogging()
	c := config.ReadSiteConfig()
	server.StartServer(c)
}
