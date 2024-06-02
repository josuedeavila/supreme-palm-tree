package event_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/josuedeavila/supreme-palm-tree/internal/interface/controller/http/event"
	usecaseEvent "github.com/josuedeavila/supreme-palm-tree/internal/usecase/event"
	"github.com/josuedeavila/supreme-palm-tree/internal/usecase/event/fake"

	"github.com/matryer/is"
)

func TestCreateEvent(t *testing.T) {
	t.Run("deposit success", func(t *testing.T) {
		is := is.New(t)

		eventBody := `{
			"type": "deposit",
			"amount": 100,
			"destination": "0"
		}`
		req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/event"), bytes.NewReader([]byte(eventBody)))

		w := httptest.NewRecorder()
		_, engine := gin.CreateTestContext(w)

		useCase := fake.NewUseCases(func(e *usecaseEvent.Event) (*usecaseEvent.EventOutput, error) {
			return &usecaseEvent.EventOutput{
				Destination: &usecaseEvent.TransactionResult{
					ID:     "0",
					Amount: 100,
				},
			}, nil
		})
		handler := event.NewHandler(useCase)
		handler.RegisterEventRoutes(engine.Group("/"))

		engine.ServeHTTP(w, req)
		is.Equal(http.StatusCreated, w.Code)

		output := new(bytes.Buffer)
		json.Compact(output, w.Body.Bytes())
		is.Equal(output.String(), `{"destination":{"id":"0","amount":100}}`)
	})

	t.Run("withdraw success", func(t *testing.T) {
		is := is.New(t)

		eventBody := `{
			"type": "withdraw",
			"amount": 100,
			"origin": "0"
		}`
		req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/event"), bytes.NewReader([]byte(eventBody)))

		w := httptest.NewRecorder()
		_, engine := gin.CreateTestContext(w)

		useCase := fake.NewUseCases(func(e *usecaseEvent.Event) (*usecaseEvent.EventOutput, error) {
			return &usecaseEvent.EventOutput{
				Origin: &usecaseEvent.TransactionResult{
					ID:     "0",
					Amount: 100,
				},
			}, nil
		})
		handler := event.NewHandler(useCase)
		handler.RegisterEventRoutes(engine.Group("/"))

		engine.ServeHTTP(w, req)
		is.Equal(http.StatusCreated, w.Code)

		output := new(bytes.Buffer)
		json.Compact(output, w.Body.Bytes())
		is.Equal(output.String(), `{"origin":{"id":"0","amount":100}}`)
	})

	t.Run("transfer success", func(t *testing.T) {
		is := is.New(t)

		eventBody := `{
			"type": "transfer",
			"amount": 100,
			"origin": "0",
			"destination": "1"
		}`
		req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/event"), bytes.NewReader([]byte(eventBody)))

		w := httptest.NewRecorder()
		_, engine := gin.CreateTestContext(w)

		useCase := fake.NewUseCases(func(e *usecaseEvent.Event) (*usecaseEvent.EventOutput, error) {
			return &usecaseEvent.EventOutput{
				Origin: &usecaseEvent.TransactionResult{
					ID:     "0",
					Amount: 100,
				},
				Destination: &usecaseEvent.TransactionResult{
					ID:     "1",
					Amount: 100,
				},
			}, nil
		})
		handler := event.NewHandler(useCase)
		handler.RegisterEventRoutes(engine.Group("/"))

		engine.ServeHTTP(w, req)
		is.Equal(http.StatusCreated, w.Code)

		output := new(bytes.Buffer)
		json.Compact(output, w.Body.Bytes())
		is.Equal(output.String(), `{"origin":{"id":"0","amount":100},"destination":{"id":"1","amount":100}}`)
	})

	t.Run("invalid event", func(t *testing.T) {
		is := is.New(t)

		eventBody := `{
			"type": "invalid",
			"amount": 100,
			"origin": "0",
			"destination": "1"
		}`
		req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/event"), bytes.NewReader([]byte(eventBody)))

		w := httptest.NewRecorder()
		_, engine := gin.CreateTestContext(w)

		useCase := fake.NewUseCases(func(e *usecaseEvent.Event) (*usecaseEvent.EventOutput, error) {
			return nil, fmt.Errorf("invalid transaction type")
		})
		handler := event.NewHandler(useCase)
		handler.RegisterEventRoutes(engine.Group("/"))

		engine.ServeHTTP(w, req)
		is.Equal(http.StatusInternalServerError, w.Code)
	})

	t.Run("invalid payload", func(t *testing.T) {
		is := is.New(t)

		eventBody := `{
			"type": "deposit",
			"amount": 100,
			"destination": "0"
		`
		req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/event"), bytes.NewReader([]byte(eventBody)))

		w := httptest.NewRecorder()
		_, engine := gin.CreateTestContext(w)

		useCase := fake.NewUseCases(func(e *usecaseEvent.Event) (*usecaseEvent.EventOutput, error) {
			return nil, fmt.Errorf("invalid transaction type")
		})
		handler := event.NewHandler(useCase)
		handler.RegisterEventRoutes(engine.Group("/"))

		engine.ServeHTTP(w, req)
		is.Equal(http.StatusBadRequest, w.Code)
	})
}
