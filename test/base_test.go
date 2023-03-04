package test

import (
	"main/internal/api"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	BaseUrl = "/api/v1"
)

func Test_Health(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, BaseUrl+"/health", nil)
	writer := httptest.NewRecorder()

	server := api.InitServer()
	server.Start("localhost",1000)
	server.Engine.ServeHTTP(writer, req)
}

func Test_BaseHandler(t *testing.T) {

}

func Test_NoHandlers(t *testing.T) {

}