package provider

import (
	"encoding/json"
	"fmt"
	env "github/leonardoas10/go-provider-pattern/src/pkg/common/env"
	models "github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/models"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/gommon/log"
)

type provider struct {}

func NewProvider() *provider  {
	return &provider{}
}

func (p *provider) GetJsonPlaceHolders() ([]models.JsonPlaceHolder, int, error) {
	res, err := http.Get(env.GetEnvVariable("URL"))
	
	if err != nil {
		return []models.JsonPlaceHolder{}, 500, err
	}

	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
        log.Fatal(err)
    }
    bodyString := string(bodyBytes)

	var users []models.JsonPlaceHolder
	if err := json.Unmarshal([]byte(bodyString), &users); err != nil {
        panic(err)
    }

	return users, res.StatusCode, nil
}

func (p *provider) GetJsonPlaceHolder(id int) (models.JsonPlaceHolder, int, error) {
	res, err := http.Get(env.GetEnvVariable("URL") + "/" +strconv.Itoa(id))
	
	if err != nil {
		return models.JsonPlaceHolder{}, 500, err
	}
	defer res.Body.Close()

	jsonPlaceHolder := new(models.JsonPlaceHolder)
	errors := json.NewDecoder(res.Body).Decode(jsonPlaceHolder)

	if errors != nil {
		return models.JsonPlaceHolder{}, 500, err
	}

	// Print the content of the jsonPlaceHolder variable
	fmt.Printf("jsonPlaceHolder Object: %+v\n", jsonPlaceHolder)

	return models.JsonPlaceHolder{
		UserId: jsonPlaceHolder.UserId,
		Id: jsonPlaceHolder.Id,
		Title: jsonPlaceHolder.Title,
		Completed: jsonPlaceHolder.Completed,
	}, res.StatusCode, nil
}

func (p *provider) UpdateJsonPlaceHolder(jsonPlacerHolder models.UpdateJsonPlaceHolder) (models.JsonPlaceHolder, int, error)  {
	retrieveJsonPlaceHolder, status, err := p.GetJsonPlaceHolder(jsonPlacerHolder.Id)

	if err != nil {
		return models.JsonPlaceHolder{}, 500, err
    }

	retrieveJsonPlaceHolder.Title = jsonPlacerHolder.Title

	return retrieveJsonPlaceHolder, status, nil

}
