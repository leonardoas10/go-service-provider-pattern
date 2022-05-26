package http

import (
	"encoding/json"
	"fmt"
	jsonplaceholders "github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/provider"
	service "github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/service"

	"github.com/labstack/echo/v4"
)

func StartHttp(port string)  {
	e := echo.New()

	jsonplaceholdersProvider := jsonplaceholders.NewProvider()
	jsonplaceholdersService := service.NewService(jsonplaceholdersProvider)

	e.GET("/json-placeholders", func(c echo.Context) error {
		response, status, err := jsonplaceholdersService.WhoAreThey()

		if err != nil {
			return c.JSON(status, err)
		}

		return c.JSON(status, response)
	})

	e.POST("/json-placeholder", func(c echo.Context) error {
		json_map := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&json_map)
		// id := json_map["id"]

		println(json_map)

		if err != nil {
			return c.JSON(500, err)
		}

		// response, status, err := jsonplaceholdersService.WhoIs(id)

		// if err != nil {
		// 	return c.JSON(status, err)
		// }

		// return c.JSON(status, response)
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", port)))
}