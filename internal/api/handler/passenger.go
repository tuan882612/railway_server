package handler

import (
	"main/internal/domain/passenger"
	"main/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PassengerController struct {
	Repo passenger.PassengerRepository
}

func (pc PassengerController) GetAll(ctx *gin.Context) {
	data := pc.Repo.GetAllPassengers()

	if data == nil {
		ctx.JSON(http.StatusNoContent, response.Empty())
		return
	}

	res := response.Custom(&data, http.StatusOK, "")
	ctx.JSON(http.StatusOK, res)
}

func (pc PassengerController) GetWaiting(ctx *gin.Context) {
	data := pc.Repo.GetWaiting()

	if data == nil {
		ctx.JSON(http.StatusNoContent, response.Empty())
		return
	}

	res := response.Custom(&data, http.StatusOK, "")
	ctx.JSON(http.StatusOK, res)
}

func (pc PassengerController) GetArea(ctx *gin.Context) {
	code := ctx.Query("area-code")

	if code == "" {
		res := response.Error("Empty parameter.")
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	data := pc.Repo.GetAllPsgAreaCode(code)

	if data == nil {
		ctx.JSON(http.StatusNoContent, response.Empty())
		return
	}

	res := response.Custom(&data, http.StatusOK, "")
	ctx.JSON(http.StatusOK, res)
}

func (pc PassengerController) GetTrainsRiding(ctx *gin.Context) {
	firstName := ctx.Query("first-name")
	lastName := ctx.Query("last-name")

	if lastName == "" || firstName == "" {
		res := response.Error("Empty parameter.")
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	
	data := pc.Repo.GetTrainsPsgRiding(firstName, lastName)

	if data == nil {
		ctx.JSON(http.StatusNoContent, response.Empty())
		return
	}

	res := response.Custom(&data, http.StatusOK, "")
	ctx.JSON(http.StatusOK, res)
}

func (pc PassengerController) GetAllwTrainName(ctx *gin.Context) {
	trainName := ctx.Query("train-name")

	if trainName == "" {
		res := response.Error("Empty parameter.")
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	data := pc.Repo.GetPsgConfirmedTrainName(trainName)

	if data == nil {
		ctx.JSON(http.StatusNoContent, response.Empty())
		return
	}

	res := response.Custom(data, http.StatusOK, "")
	ctx.JSON(http.StatusOK, res)
}

func (pc PassengerController) GetTraveling(ctx *gin.Context) {
	day := ctx.Query("day")

	if day == "" {
		res := response.Error("Empty parameter.")
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	data := pc.Repo.GetPsgRiding(day)

	if data == nil {
		ctx.JSON(http.StatusNoContent, response.Empty())
		return
	}

	res := response.Custom(data, http.StatusOK, "")
	ctx.JSON(http.StatusOK, res)
}

func (pc PassengerController) GetTravelingConfirmed(ctx *gin.Context) {
	day := ctx.Query("day")

	if day == "" {
		res := response.Error("Empty parameter.")
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	data := pc.Repo.GetPsgRidingConfirmed(day)

	if data == nil {
		ctx.JSON(http.StatusNoContent, response.Empty())
		return
	}

	res := response.Custom(data, http.StatusOK, "")
	ctx.JSON(http.StatusOK, res)
}