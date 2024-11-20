package main

import (
	"fmt"
	"github.com/akbaralishaikh/denti/cmd/server"
	"github.com/akbaralishaikh/denti/pkg/di"
	"github.com/akbaralishaikh/denti/pkg/logger"
	"github.com/pphee/test/vendor/github.com/akbaralishaikh/denti/pkg/config"
	"os"

	"github.com/gin-gonic/gin"
)

type QdrantConfig struct {
	Host           string
	ApiKey         string
	CollectionName string
	Port           int
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
}

func run() error {

	cfg, err := config.NewConfig()
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	g := gin.Default()
	d := di.BuildContainer()

	var l logger.LogInfoFormat
	di.Invoke(func(log logger.LogInfoFormat) {
		l = log
	})

	svr := server.NewServer(g, d, l, cfg)

	svr.MapRoutes()
	if err := svr.SetupDB(); err != nil {
		return err
	}
	if err := svr.SetupQdrant(); err != nil {
		return err
	}
	return svr.Start()
}
