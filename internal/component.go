package internal

import (
	"github.com/gin-gonic/gin"

	eventHandler "github.com/josuedeavila/supreme-palm-tree/internal/interface/controller/http/event"
	eventUseCases "github.com/josuedeavila/supreme-palm-tree/internal/usecase/event"

	"github.com/josuedeavila/supreme-palm-tree/internal/entity"
	balanceHandler "github.com/josuedeavila/supreme-palm-tree/internal/interface/controller/http/balance"
	resetHandler "github.com/josuedeavila/supreme-palm-tree/internal/interface/controller/http/reset"
	resetUseCases "github.com/josuedeavila/supreme-palm-tree/internal/usecase/reset"

	balanceUseCases "github.com/josuedeavila/supreme-palm-tree/internal/usecase/balance"
)

// Component expose component interface
type Component interface {
	RegisterEventRoutes(group *gin.RouterGroup)
	RegisterBalanceRoutes(group *gin.RouterGroup)
	RegisterResetRoutes(group *gin.RouterGroup)
}

type component struct {
	balanceHandler balanceHandler.Handler
	resetHandler   resetHandler.Handler
	eventHandler   eventHandler.Handler
}

// New allocates component
func New(balance entity.BalanceRepository, event entity.EventRepository) Component {
	buc := balanceUseCases.New(balance)
	euc := eventUseCases.New(event, balance)
	ruc := resetUseCases.New(event, balance)
	resetHdl := resetHandler.NewHandler(ruc)
	balanceHdl := balanceHandler.NewHandler(buc)
	eventHdl := eventHandler.NewHandler(euc)

	return &component{
		balanceHandler: balanceHdl,
		resetHandler:   resetHdl,
		eventHandler:   eventHdl,
	}
}

func (c *component) RegisterBalanceRoutes(group *gin.RouterGroup) {
	c.balanceHandler.RegisterBalanceRoutes(group)
}

func (c *component) RegisterEventRoutes(group *gin.RouterGroup) {
	c.eventHandler.RegisterEventRoutes(group)
}

func (c *component) RegisterResetRoutes(group *gin.RouterGroup) {
	c.resetHandler.RegisterResetRoutes(group)
}
