package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/models"
	provider "github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/provider"
	service "github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/service"
	"io/ioutil"
	"strconv"

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
      
    // Get the ID from the URL parameters
    id := c.Param("id")
    fmt.Println("ID from URL: ", id)
  
    idInt, err := strconv.Atoi(id)
    if err != nil {
        return c.JSON(400, map[string]string{"error": "Invalid ID format"})
    }

    response, status, err := jsonplaceholdersService.WhoIs(idInt)

    if err != nil {
        return c.JSON(status, map[string]string{"error": err.Error()})
    }

    return c.JSON(status, response)
}

func ConcurrentChangeTitles(c echo.Context) error {
    jsonplaceholdersProvider := provider.NewProvider()
    jsonplaceholdersService := service.NewService(jsonplaceholdersProvider)

    response, status, err := jsonplaceholdersService.ConcurrentChangeTitles()
    if err != nil {
        return c.JSON(status, map[string]string{"error": err.Error()})
    }

    return c.JSON(status, response)
}

func UpdateJsonPlaceHolder(c echo.Context) error  {
    jsonplaceholdersProvider := provider.NewProvider()
    jsonplaceholdersService := service.NewService(jsonplaceholdersProvider)

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

    response, status, err := jsonplaceholdersService.UpdateJsonPlaceHolder(*postJsonPlaceHolder)
    if err != nil {
        return c.JSON(status, map[string]string{"error": err.Error()})
    }

    return c.JSON(status, response)
}



