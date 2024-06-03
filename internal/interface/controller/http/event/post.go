package event

import (
	"net/http"
	"strings"

	"github.com/josuedeavila/supreme-palm-tree/internal/interface/controller/dto"
	"github.com/josuedeavila/supreme-palm-tree/internal/usecase/event"

	"github.com/gin-gonic/gin"
)

// PostHandler handle post request
func (eh *eventHandler) PostHandler(ctx *gin.Context) {
	event, err := buildEventPayload(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := eh.useCases.Create(event)
	if err != nil {
		if strings.Contains(err.Error(), "balance not found") {
			ctx.String(http.StatusNotFound, "0")
			return
		}
		ctx.String(http.StatusInternalServerError, "0")
		return
	}

	if result == nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	output := dto.EventOutput{}
	if result.Origin != nil {
		output.Origin = &dto.TransactionResult{
			ID:      result.Origin.ID,
			Balance: result.Origin.Balance,
		}
	}

	if result.Destination != nil {
		output.Destination = &dto.TransactionResult{
			ID:      result.Destination.ID,
			Balance: result.Destination.Balance,
		}
	}

	ctx.JSON(http.StatusCreated, output)
}

func buildEventPayload(ctx *gin.Context) (*event.Event, error) {
	input := &dto.Event{}
	if err := ctx.BindJSON(input); err != nil {
		return nil, err
	}
	return &event.Event{
		Type:        event.TransactionType(input.Type),
		Origin:      input.Origin,
		Amount:      input.Amount,
		Destination: input.Destination,
		ID:          input.ID,
	}, nil
}
