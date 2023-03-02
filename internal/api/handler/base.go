package handler

import (
	"main/tools/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Default(ctx *gin.Context) {
	body := map[string]interface{}{}
	res := response.Custom(body, http.StatusOK, "Base endpoint")
	ctx.JSON(http.StatusOK, res)
}

func Health(ctx *gin.Context) {
	res := map[string]string{"heatlh":"good"}
	ctx.JSON(http.StatusOK, res)
}

func None(ctx *gin.Context) {
	res := response.Custom(
		map[string]string{},
		http.StatusNotFound,
		"That endpoint does not exist",
	)
	ctx.JSON(http.StatusNotFound, res)
}