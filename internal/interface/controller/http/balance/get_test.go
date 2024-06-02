package balance_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/josuedeavila/supreme-palm-tree/internal/interface/controller/http/balance"
	usecaseBalance "github.com/josuedeavila/supreme-palm-tree/internal/usecase/balance"
	"github.com/josuedeavila/supreme-palm-tree/internal/usecase/balance/fake"

	"github.com/matryer/is"
)

func TestGetBalance(t *testing.T) {
	t.Run("balance not found", func(t *testing.T) {
		is := is.New(t)
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/balance?account_id=%s", "0"), nil)

		w := httptest.NewRecorder()
		_, engine := gin.CreateTestContext(w)

		useCase := fake.NewUseCases(func(accountID int) (*usecaseBalance.Balance, error) {
			return &usecaseBalance.Balance{
				AccountID: 0,
				Amount:    100,
			}, nil
		})
		handler := balance.NewHandler(useCase)
		handler.RegisterBalanceRoutes(engine.Group("/"))

		engine.ServeHTTP(w, req)
		is.Equal(http.StatusOK, w.Code)
		is.Equal(w.Body.String(), "100")
	})

	t.Run("balance found", func(t *testing.T) {
		is := is.New(t)
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/balance?account_id=%s", "0"), nil)

		w := httptest.NewRecorder()
		_, engine := gin.CreateTestContext(w)

		useCase := fake.NewUseCases(func(accountID int) (*usecaseBalance.Balance, error) {
			return nil, fmt.Errorf("balance not found")
		})
		handler := balance.NewHandler(useCase)
		handler.RegisterBalanceRoutes(engine.Group("/"))

		engine.ServeHTTP(w, req)
		is.Equal(http.StatusNotFound, w.Code)
		is.Equal(w.Body.String(), "0")
	})

	t.Run("internal error", func(t *testing.T) {
		is := is.New(t)
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/balance?account_id=%s", "0"), nil)

		w := httptest.NewRecorder()
		_, engine := gin.CreateTestContext(w)

		useCase := fake.NewUseCases(func(accountID int) (*usecaseBalance.Balance, error) {
			return nil, fmt.Errorf("error")
		})
		handler := balance.NewHandler(useCase)
		handler.RegisterBalanceRoutes(engine.Group("/"))

		engine.ServeHTTP(w, req)
		is.Equal(http.StatusInternalServerError, w.Code)
		is.Equal(w.Body.String(), "0")
	})

	t.Run("invalid account id", func(t *testing.T) {
		is := is.New(t)
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/balance?account_id=%s", "invalid"), nil)

		w := httptest.NewRecorder()
		_, engine := gin.CreateTestContext(w)

		useCase := fake.NewUseCases(nil)
		handler := balance.NewHandler(useCase)
		handler.RegisterBalanceRoutes(engine.Group("/"))

		engine.ServeHTTP(w, req)
		is.Equal(http.StatusBadRequest, w.Code)
		is.Equal(w.Body.String(), `{"error":"Invalid account ID"}`)
	})

}
