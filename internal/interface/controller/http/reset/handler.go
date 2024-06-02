package reset

import (
	"github.com/gin-gonic/gin"

	"github.com/josuedeavila/supreme-palm-tree/internal/usecase/reset"
)

// Handler exposes interface for handlers of elected product
type Handler interface {
	RegisterResetRoutes(c *gin.RouterGroup)
}

type resetHandler struct {
	useCases reset.UseCases
}

// NewHandler allocate elected handler
func NewHandler(useCases reset.UseCases) Handler {
	return &resetHandler{
		useCases,
	}
}

// RegisterResetRoutes register routes
func (eh *resetHandler) RegisterResetRoutes(routes *gin.RouterGroup) {
	electedRoutes := routes.Group("/reset")
	electedRoutes.POST("", eh.PostHandler)
}
