package controller

import (
	"github.com/gin-gonic/gin"
	"main/db"
	"main/request"
	"main/response"
	"net/http"
)

type AccountController struct {
	DB *db.AccountDB
}

func (receiver AccountController) Create(ctx *gin.Context) {
	var account request.Account
	if err := ctx.ShouldBindJSON(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Error: err.Error()})
		return
	}

	err := receiver.DB.Create(account)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: err.Error()})
		return
	}
	ctx.Status(http.StatusCreated)
}

func (receiver AccountController) Get(ctx *gin.Context) {
	acc, err := receiver.DB.GetAccount()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, acc)
}
