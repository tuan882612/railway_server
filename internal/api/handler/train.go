package handler

import (
	"main/internal/domain/train"
	"main/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TrainController struct {
	Repo train.TrainRepository
}

func (tc TrainController) GetAll(ctx *gin.Context) {
	data := tc.Repo.GetAllTrains()

	if data == nil {
		ctx.JSON(http.StatusNoContent, response.Empty())
		return
	}

	res := response.Custom(&data, http.StatusOK, "")
	ctx.JSON(http.StatusOK, res)
} 

func (tc TrainController) GetInfo(ctx *gin.Context) {
	name := ctx.Query("name")

	if name == "" {
		res := response.Error("Empty parameter.")
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	data := tc.Repo.GetTrainInfoSum(name)

	if data == nil {
		ctx.JSON(http.StatusNoContent, response.Empty())
		return
	}

	res := response.Custom(&data, http.StatusOK, "")
	ctx.JSON(http.StatusOK, res)
}

func (tc TrainController) GetInfoAge(ctx *gin.Context) {
	start := ctx.Query("start")
	end := ctx.Query("end")

	if start == "" || end == "" {
		res := response.Error("Empty parameter.")
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	data := tc.Repo.GetInfoAgeRange(start, end)

	if data == nil {
		ctx.JSON(http.StatusNoContent, response.Empty())
		return
	}
	res := response.Custom(&data, http.StatusOK, "")
	ctx.JSON(http.StatusOK, res)
}