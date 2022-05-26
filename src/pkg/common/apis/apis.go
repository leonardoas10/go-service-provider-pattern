package api

import (
	handlersJsonPlaceHolders "github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/handlers"

	"github.com/labstack/echo/v4"
)

func JsonplaceholdersGroup(e *echo.Group)  {
	e.GET("", handlersJsonPlaceHolders.GetJsonPlaceHolders)
	e.POST("", handlersJsonPlaceHolders.GetJsonPlaceHolder)
}