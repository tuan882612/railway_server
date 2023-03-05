package test

import (
	"encoding/json"
	"main/internal/api"
	"main/test"
	"main/pkg/response"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func GetServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	server := api.InitServer("localhost", 2000)
	server.LoadConfig()
	return server.Engine
}

func Test_Health(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, test.BaseUrl+"/health", nil)
	writer := httptest.NewRecorder()
	GetServer().ServeHTTP(writer, req)

	var res map[string]string
	json.Unmarshal(writer.Body.Bytes(), &res)

	expected := map[string]string{"health":"good"}

	assert.Equal(t, expected, res)
}

func Test_BaseHandler(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, test.BaseUrl, nil)
	writer := httptest.NewRecorder()
	GetServer().ServeHTTP(writer, req)

	var res response.Base
	json.Unmarshal(writer.Body.Bytes(), &res)

	body := map[string]interface{}{}
	expected := response.Custom(body, http.StatusOK, "Base endpoint")

	assert.Equal(t, expected, res)
}

func Test_NoHandlers(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "invalid", nil)
	writer := httptest.NewRecorder()
	GetServer().ServeHTTP(writer, req)

	assert.Equal(t, http.StatusNotFound, writer.Code)
}
