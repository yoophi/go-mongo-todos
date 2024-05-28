package webserver

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func NewRouter(injector *do.Injector) (*gin.Engine, error) {
	r := gin.Default()
	return r, nil
}
