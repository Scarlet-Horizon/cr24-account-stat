package controller

import (
	"github.com/gin-gonic/gin"
	"main/db"
	"main/request"
	"main/response"
	"net/http"
)

type StatController struct {
	DB *db.StatDB
}

func (receiver StatController) CreateStat(ctx *gin.Context) {
	var req request.StatRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{Error: err.Error()})
		return
	}

	err := receiver.DB.CreateStat(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: err.Error()})
		return
	}
	ctx.Status(http.StatusCreated)
}
