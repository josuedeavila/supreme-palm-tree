package reset

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (eh *resetHandler) PostHandler(c *gin.Context) {
	err := eh.useCases.Reset()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, "OK")
}
