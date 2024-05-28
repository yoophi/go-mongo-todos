package main

import (
	"go-mongo-todos/pkg/app/delivery/webserver"

	"github.com/samber/do"
	"go.uber.org/zap"
)

var injector *do.Injector

func init() {
	injector = do.New()

	do.Provide(injector, webserver.NewRouter)
	do.Provide(injector, webserver.NewWebServer)

	zap.NewDevelopmentConfig()
}

func main() {
	runWebServer()
}

func runWebServer() {
	server := do.MustInvoke[*webserver.WebServer](injector)
	err := server.ListenAndServe()
	if err != nil {
		zap.L().Error("closed", zap.Error(err))
	}
}
