package reset_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/josuedeavila/supreme-palm-tree/internal/interface/controller/http/reset"
	"github.com/josuedeavila/supreme-palm-tree/internal/usecase/reset/fake"
	"github.com/matryer/is"
)

func TestReset(t *testing.T) {
	t.Run("reset success", func(t *testing.T) {
		is := is.New(t)
		req := httptest.NewRequest(http.MethodPost, "/reset", nil)

		w := httptest.NewRecorder()
		_, engine := gin.CreateTestContext(w)

		useCase := fake.NewUseCases(func() error {
			return nil
		})
		handler := reset.NewHandler(useCase)
		handler.RegisterResetRoutes(engine.Group("/"))

		engine.ServeHTTP(w, req)
		is.Equal(http.StatusOK, w.Code)
		is.Equal(w.Body.String(), "OK")
	})

	t.Run("reset error", func(t *testing.T) {
		is := is.New(t)
		req := httptest.NewRequest(http.MethodPost, "/reset", nil)

		w := httptest.NewRecorder()
		_, engine := gin.CreateTestContext(w)

		useCase := fake.NewUseCases(func() error {
			return fmt.Errorf("internal error")
		})
		handler := reset.NewHandler(useCase)
		handler.RegisterResetRoutes(engine.Group("/"))

		engine.ServeHTTP(w, req)
		is.Equal(http.StatusInternalServerError, w.Code)
		is.Equal(w.Body.String(), "internal error")
	})
}
