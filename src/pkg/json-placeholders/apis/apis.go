package api

import (
	handlersJsonPlaceHolders "github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/handlers"
	structs "github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/structs"
	customMiddlewares "github/leonardoas10/go-provider-pattern/src/pkg/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JsonplaceholdersGroup(e *echo.Group)  {
	e.GET("/:id", handlersJsonPlaceHolders.GetJsonPlaceHolder, customMiddlewares.ParamsValidatorMiddleware("id"), middleware.Logger())
	e.GET("", handlersJsonPlaceHolders.GetJsonPlaceHolders, middleware.Logger())
	e.PUT("", handlersJsonPlaceHolders.UpdateJsonPlaceHolder,  customMiddlewares.BodyValidatorMiddleware(&structs.RequestUpdateJsonPlaceHolder{}), middleware.Logger())
	e.GET("/concurrent-change-titles",  handlersJsonPlaceHolders.ConcurrentChangeTitles, middleware.Logger())
}