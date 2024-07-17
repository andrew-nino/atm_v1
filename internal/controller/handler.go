package controller

import (
	"github.com/andrew-nino/atm_v1/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.New()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/accounts", h.addAccount)
		v1.POST("/accounts/:id/deposit", h.deposit)
		v1.POST("/accounts/:id/withdraw", h.withdraw)
		v1.GET("/accounts/:id/balance", h.getBalance)

	}
	return router
}
