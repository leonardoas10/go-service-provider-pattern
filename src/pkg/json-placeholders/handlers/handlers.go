package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/models"
	service "github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/service"
	"io/ioutil"

	"github.com/labstack/echo/v4"
)

func GetJsonPlaceHolders(p service.JsonPlaceHoldersProvider, c echo.Context) error {
    service := service.NewService(p)

	response, status, err := service.WhoAreThey()

	if err != nil {
		return c.JSON(status, err)
	}

	return c.JSON(status, response)
}

func GetJsonPlaceHolder(p service.JsonPlaceHoldersProvider, c echo.Context) error {
    service := service.NewService(p)
      
    // Get the ID from the URL parameters
    id := c.Param("id")
    fmt.Println("ID from URL: ", id)

    response, status, err := service.WhoIs(id)

    if err != nil {
        return c.JSON(status, map[string]string{"error": err.Error()})
    }

    return c.JSON(status, response)
}

func ConcurrentChangeTitles(p service.JsonPlaceHoldersProvider, c echo.Context) error {
    service := service.NewService(p)

    response, status, err := service.ConcurrentChangeTitles()
    if err != nil {
        return c.JSON(status, map[string]string{"error": err.Error()})
    }

    return c.JSON(status, response)
}

func UpdateJsonPlaceHolder(p service.JsonPlaceHoldersProvider, c echo.Context) error  {
    service := service.NewService(p)

       // Read the request body
    body, err := ioutil.ReadAll(c.Request().Body)
    if err != nil {
        return c.JSON(500, err)
    }
    defer c.Request().Body.Close() // Close the body after reading
   // Print the request body as a string
    fmt.Println("Request Body: ", string(body))

    // Create an instance of PostJsonPlaceHolder to decode the JSON
    postJsonPlaceHolder := new(models.UpdateJsonPlaceHolder)

    // Reset the position of the original request body
    c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))

    // Unmarshal the JSON into the postJsonPlaceHolder struct
    err = json.NewDecoder(c.Request().Body).Decode(&postJsonPlaceHolder)
    if err != nil {
        return c.JSON(400, map[string]string{"error": "Invalid JSON format"})
    }

    response, status, err := service.UpdateJsonPlaceHolder(*postJsonPlaceHolder)
    if err != nil {
        return c.JSON(status, map[string]string{"error": err.Error()})
    }

    return c.JSON(status, response)
}



