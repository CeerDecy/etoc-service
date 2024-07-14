package handler

import (
	"etoc-service/internal/app/http/handler/docker"
	"etoc-service/internal/app/http/router"
	"etoc-service/internal/app/http/svc"
)

func Handler() func(router *router.Router) {
	return func(router *router.Router) {
		common := router.Group("/common")
		common.GET("/health-check", func(ctx *svc.Context) error {
			ctx.Success(nil)
			return nil
		})

		dockerGroup := common.Group("/docker")
		dockerGroup.POST("/try/connection", docker.TryConnect)
	}
}
