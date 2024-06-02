package event

import (
	"net/http"

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
		// TODO: deal with different kind of errors
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result == nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	output := dto.EventOutput{}
	if result.Origin != nil {
		output.Origin = &dto.TransactionResult{
			ID:     result.Origin.ID,
			Amount: result.Origin.Amount,
		}
	}

	if result.Destination != nil {
		output.Destination = &dto.TransactionResult{
			ID:     result.Destination.ID,
			Amount: result.Destination.Amount,
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
