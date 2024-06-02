package balance

import (
	"net/http"
	"strconv"

	"github.com/josuedeavila/supreme-palm-tree/internal/interface/controller/dto"

	"github.com/gin-gonic/gin"
)

// GetHandler handle post request
func (eh *balaceHandler) GetHandler(ctx *gin.Context) {
	queryParams := ctx.Request.URL.Query()
	accountID := queryParams.Get("account_id")
	if accountID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "account_id is required"})
		return
	}

	accountIDInt, err := strconv.Atoi(accountID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	result, err := eh.useCases.Get(accountIDInt)
	if err != nil {
		if err.Error() == "balance not found" {
			ctx.String(http.StatusNotFound, "0")
			return
		}
		ctx.String(http.StatusInternalServerError, "0")
		return
	}

	output := dto.Balance{
		AccountID: result.AccountID,
		Amount:    result.Amount,
	}

	ctx.String(http.StatusOK, strconv.Itoa(output.Amount))
}
