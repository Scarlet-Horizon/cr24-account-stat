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

func (receiver StatController) LastEndpoint(ctx *gin.Context) {
	last, err := receiver.DB.LastEndpoint()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, last)
}

func (receiver StatController) MostCalled(ctx *gin.Context) {
	most, err := receiver.DB.MostCalled()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, most)
}
