package event

import (
	"github.com/gin-gonic/gin"

	"github.com/josuedeavila/supreme-palm-tree/internal/usecase/event"
)

// Handler exposes interface for handlers of elected product
type Handler interface {
	RegisterEventRoutes(c *gin.RouterGroup)
}

type eventHandler struct {
	useCases event.UseCases
}

// NewHandler allocate elected handler
func NewHandler(useCases event.UseCases) Handler {
	return &eventHandler{
		useCases,
	}
}

// RegisterEventRoutes register routes
func (eh *eventHandler) RegisterEventRoutes(routes *gin.RouterGroup) {
	electedRoutes := routes.Group("/event")
	electedRoutes.POST("", eh.PostHandler)
}
