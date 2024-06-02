package balance_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/josuedeavila/supreme-palm-tree/internal/interface/controller/http/balance"
	"github.com/josuedeavila/supreme-palm-tree/internal/usecase/balance/fake"
	"github.com/matryer/is"
)

func TestRegisterRoutes(t *testing.T) {
	w := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(w)
	is := is.New(t)

	useCase := fake.NewUseCases(nil)
	handler := balance.NewHandler(useCase)
	handler.RegisterBalanceRoutes(engine.Group("/"))
	routesInfo := engine.Routes()
	routesMethodAndPath := make([][]string, 0, len(routesInfo))
	for _, routeInfo := range routesInfo {
		routesMethodAndPath = append(routesMethodAndPath, []string{routeInfo.Method, routeInfo.Path})
	}

	expectedRoutesMethodAndPath := [][]string{
		{"GET", "/balance"},
	}

	is.Equal(expectedRoutesMethodAndPath, routesMethodAndPath)
}
