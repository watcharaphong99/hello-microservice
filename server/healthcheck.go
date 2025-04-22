package server

import (
	"net/http"

	"github.com/bunkai888/hello-microservice/pkg/response"
	"github.com/labstack/echo/v4"
)

type helthCheck struct {
	App    string `json:"app"`
	Status string `json:"staus"`
}

func (s *server) healthCheckService(c echo.Context) error {
	return response.SuccessResposse(c, http.StatusOK, &helthCheck{
		App:    s.cfg.App.Name,
		Status: "OK",
	})
}

// go run main.go ./env/dev/.env.inventory = คำสั่ง Run
//
