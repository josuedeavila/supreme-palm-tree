package event_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/josuedeavila/supreme-palm-tree/internal/interface/controller/http/event"
	"github.com/josuedeavila/supreme-palm-tree/internal/usecase/event/fake"
	"github.com/matryer/is"
)

func TestRegisterRoutes(t *testing.T) {
	w := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(w)
	is := is.New(t)

	useCase := fake.NewUseCases(nil)
	handler := event.NewHandler(useCase)
	handler.RegisterEventRoutes(engine.Group("/"))
	routesInfo := engine.Routes()
	routesMethodAndPath := make([][]string, 0, len(routesInfo))
	for _, routeInfo := range routesInfo {
		routesMethodAndPath = append(routesMethodAndPath, []string{routeInfo.Method, routeInfo.Path})
	}

	expectedRoutesMethodAndPath := [][]string{
		{"POST", "/event"},
	}

	is.Equal(expectedRoutesMethodAndPath, routesMethodAndPath)
}
