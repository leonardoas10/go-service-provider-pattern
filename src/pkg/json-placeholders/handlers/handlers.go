package handlers

import (
	"encoding/json"
	models "github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/models"
	provider "github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/provider"
	service "github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/service"

	"github.com/labstack/echo/v4"
)

func GetJsonPlaceHolders(c echo.Context) error {
	jsonplaceholdersProvider := provider.NewProvider()
	jsonplaceholdersService := service.NewService(jsonplaceholdersProvider)

	response, status, err := jsonplaceholdersService.WhoAreThey()

	if err != nil {
		return c.JSON(status, err)
	}

	return c.JSON(status, response)
}

func GetJsonPlaceHolder(c echo.Context) error {
	jsonplaceholdersProvider := provider.NewProvider()
	jsonplaceholdersService := service.NewService(jsonplaceholdersProvider)

	postJsonPlaceHolder := new(models.PostJsonPlaceHolder)
	err := json.NewDecoder(c.Request().Body).Decode(&postJsonPlaceHolder)

	if err != nil {
		return c.JSON(500, err)
	}

	response, status, err := jsonplaceholdersService.WhoIs(postJsonPlaceHolder.Id)

	if err != nil {
		return c.JSON(status, err)
	}

	return c.JSON(status, response)
}