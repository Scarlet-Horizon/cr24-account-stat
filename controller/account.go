package controller

import (
	"github.com/gin-gonic/gin"
	"main/db"
	"main/model"
	_ "main/model"
	"main/response"
	"net/http"
)

type AccountController struct {
	DB *db.AccountDB
}

//	@description	Store newly created account.
//	@summary		Store newly created account
//	@accept			json
//	@produce		json
//	@tags			account
//	@param			requestBody	body	model.Account	true	"Account data"
//	@success		204			"No Content"
//	@failure		400			{object}	response.ErrorResponse
//	@failure		500			{object}	response.ErrorResponse
//	@security		JWT
//	@param			Authorization	header	string	true	"Authorization"
//	@router			/account [POST]
func (receiver AccountController) Create(ctx *gin.Context) {
	var account model.Account
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

//	@description	Get stored account.
//	@summary		Get stored account
//	@produce		json
//	@tags			account
//	@success		200	{object}	model.Account
//	@failure		500	{object}	response.ErrorResponse
//	@security		JWT
//	@param			Authorization	header	string	true	"Authorization"
//	@router			/account [GET]
func (receiver AccountController) Get(ctx *gin.Context) {
	acc, err := receiver.DB.GetAccount()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, acc)
}
