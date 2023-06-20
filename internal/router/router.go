package router

import (
	"github.com/gin-gonic/gin"
	"github.com/renfy96/yy-layout-v1/internal/application/admin/service"
	"github.com/renfy96/yy-layout-v1/pkg/helper/resp"
	"github.com/renfy96/yy-layout-v1/pkg/log"
)

func NewServerHTTP(
	logger *log.Logger,
	adminHandler service.Service,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		resp.HandleSuccess(ctx, map[string]interface{}{
			"say": "Hi Y!",
		})
	})
	r.POST("/admin/login", adminHandler.Login)

	return r
}
