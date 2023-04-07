package main

import (
	"github.com/stretchr/testify/assert"
	"go-fiber-demo/configs"
	"go-fiber-demo/controllers"
	"go-fiber-demo/routes"
	"go-fiber-demo/services"
	"net/http"
	"testing"
)

func TestIndexRoute(t *testing.T) {
	tests := []struct {
		description   string
		route         string
		method        string
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "index route",
			route:         "/demo",
			method:        "GET",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "OK",
		},
		{
			description:   "/demo/i18n route",
			route:         "/demo/i18n",
			method:        "GET",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "OK",
		},
	}
	configs.LoadEnvVariables()
	app := StartApp()
	i18n := services.NewI18n()
	baseController := controllers.NewBaseController(i18n)

	routes.NewApiRoutes(app, baseController).Load()

	for _, test := range tests {
		req, _ := http.NewRequest(
			test.method,
			test.route,
			nil,
		)

		res, err := app.Test(req, -1)
		assert.Equalf(t, test.expectedError, err != nil, test.description)
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)
	}
}
