package handler

import (
	"github.com/gin-gonic/gin"
	"wallet/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/send", h.send)
		api.GET("/transactions", h.transaction)
		api.GET("/wallet/:uid/balance", h.balance)
	}

	return router
}
