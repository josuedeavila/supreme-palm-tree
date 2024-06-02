package balance

import (
	"github.com/gin-gonic/gin"

	"github.com/josuedeavila/supreme-palm-tree/internal/usecase/balance"
)

// Handler exposes interface for handlers of elected product
type Handler interface {
	RegisterBalanceRoutes(c *gin.RouterGroup)
}

type balaceHandler struct {
	useCases balance.UseCases
}

// NewHandler allocate elected handler
func NewHandler(useCases balance.UseCases) Handler {
	return &balaceHandler{
		useCases,
	}
}

// RegisterBalanceRoutes register routes
func (eh *balaceHandler) RegisterBalanceRoutes(routes *gin.RouterGroup) {
	electedRoutes := routes.Group("/balance")
	electedRoutes.GET("", eh.GetHandler)
}
