package webserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

type WebServer struct {
	router *gin.Engine
}

func (s *WebServer) ListenAndServe() error {
	srv := &http.Server{
		Addr:    ":8000",
		Handler: s.router,
	}
	return srv.ListenAndServe()
}

func NewWebServer(injector *do.Injector) (*WebServer, error) {
	return &WebServer{
		router: do.MustInvoke[*gin.Engine](injector),
	}, nil
}
