package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/rassulmagauin/simplebank/db/sqlc"
)

type createTransactionRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,min=1"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req createTransactionRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}
	if !server.validateAccount(ctx, req.FromAccountID, req.Currency) {
		return
	}
	if !server.validateAccount(ctx, req.ToAccountID, req.Currency) {
		return
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, result)

}

func (server *Server) validateAccount(ctx *gin.Context, accountID int64, currency string) bool {
	if account, err := server.store.GetAccount(ctx, accountID); err == nil {
		if account.Currency != currency {
			err := fmt.Errorf("account [%d] currency mismatch, %s vs %s", accountID, account.Currency, currency)
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return false
		}
		return true
	} else {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

}
