package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"main/db"
	"main/request"
	"main/response"
	"net/http"
)

type StatController struct {
	DB *db.StatDB
}

//	@description	Create stat.
//	@summary		Create stat
//	@accept			json
//	@produce		json
//	@tags			stat
//	@param			requestBody	body	request.StatRequest	true	"Endpoint path"
//	@success		201			"No Content"
//	@failure		400			{object}	response.ErrorResponse
//	@failure		500			{object}	response.ErrorResponse
//	@security		JWT
//	@param			Authorization	header	string	true	"Authorization"
//	@router			/stat [POST]
func (receiver StatController) CreateStat(ctx *gin.Context) {
	var req request.StatRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Error: err.Error()})
		return
	}

	err := receiver.DB.CreateStat(req)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Error: "invalid endpoint"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: err.Error()})
		return
	}
	ctx.Status(http.StatusCreated)
}

//	@description	Get last stat.
//	@summary		Get last stat
//	@produce		json
//	@tags			stat
//	@success		200	{object}	[]response.EndpointLast	"An array of response.EndpointLast"
//	@failure		500	{object}	response.ErrorResponse
//	@security		JWT
//	@param			Authorization	header	string	true	"Authorization"
//	@router			/last [GET]
func (receiver StatController) LastEndpoint(ctx *gin.Context) {
	last, err := receiver.DB.LastEndpoint()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, last)
}

//	@description	Get most frequently called endpoints.
//	@summary		Get most frequently called endpoints
//	@produce		json
//	@tags			stat
//	@success		200	{object}	[]response.EndpointStat	"An array of response.EndpointStat"
//	@failure		500	{object}	response.ErrorResponse
//	@security		JWT
//	@param			Authorization	header	string	true	"Authorization"
//	@router			/most [GET]
func (receiver StatController) MostCalled(ctx *gin.Context) {
	most, err := receiver.DB.MostCalled()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, most)
}

//	@description	Get count for all endpoints.
//	@summary		Get count for all endpoints
//	@produce		json
//	@tags			stat
//	@success		200	{object}	[]response.EndpointStat	"An array of response.EndpointStat"
//	@failure		500	{object}	response.ErrorResponse
//	@security		JWT
//	@param			Authorization	header	string	true	"Authorization"
//	@router			/all [GET]
func (receiver StatController) All(ctx *gin.Context) {
	all, err := receiver.DB.EndpointAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, all)
}
