package jsonPlaceHolderProvider

import (
	"encoding/json"
	"fmt"
	env "github/leonardoas10/go-provider-pattern/src/pkg/common/env"
	models "github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/models"
	"io"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/labstack/gommon/log"
)

type provider struct {}

func NewProvider() *provider  {
	return &provider{}
}

type jsonPlaceHolderIntermediate struct {
	UserId int `json:"userId"`
	Id int `json:"id"`
	Title string `json:"title"`
	Completed  bool `json:"completed"`
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

	// Decode into intermediate structure
	var intermediateUsers []jsonPlaceHolderIntermediate
	if err := json.Unmarshal([]byte(bodyString), &intermediateUsers); err != nil {
		log.Fatal(err)
	}

	// Convert intermediate structure to the desired type
	var users []models.JsonPlaceHolder
	for _, intermediateUser := range intermediateUsers {
		user := models.JsonPlaceHolder{
			Id: strconv.Itoa(intermediateUser.Id),
			UserId: strconv.Itoa(intermediateUser.UserId),
			Title: intermediateUser.Title,
			Completed: intermediateUser.Completed,
			
		}
		users = append(users, user)
	}	

	fmt.Printf("Number of users: %d\n", len(users))

	return users, res.StatusCode, nil
}

func (p *provider) GetJsonPlaceHolder(id string) (models.JsonPlaceHolder, int, error) {
	fmt.Println("id ==> ", id)
	res, err := http.Get(env.GetEnvVariable("URL") + "/" + id)
	
	if err != nil {
		return models.JsonPlaceHolder{}, 500, err
	}
	defer res.Body.Close()

	jsonPlaceHolder := new(jsonPlaceHolderIntermediate)
	errors := json.NewDecoder(res.Body).Decode(jsonPlaceHolder)

	if errors != nil {
		log.Fatal(errors)
		return models.JsonPlaceHolder{}, 500, err
	}

	// Print the content of the jsonPlaceHolder variable
	fmt.Printf("jsonPlaceHolder Object: %+v\n", jsonPlaceHolder)

	return models.JsonPlaceHolder{
		UserId: strconv.Itoa(jsonPlaceHolder.UserId),
		Id: strconv.Itoa(jsonPlaceHolder.Id),
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

func (p *provider) ConcurrentChangeTitles() ([]models.JsonPlaceHolder, int, error) {
	users, status, _ := p.GetJsonPlaceHolders()
	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Use a mutex to safely modify the users slice concurrently
	// var mu sync.Mutex

	var counter int64
	// 200

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Modify titles concurrently with atomic
			for {
				// Atomically increment the counter and get the current value
				userIndex := int(atomic.AddInt64(&counter, 1) - 1)

				if userIndex >= len(users) {
					break // Break when all users are processed
				}

				newTitle := fmt.Sprintf("Titulo %s", strconv.Itoa(userIndex + 1))

				// No need for mutex when updating a specific index in the slice
				users[userIndex].Title = newTitle

			}
			// or use approach with Mutex Locks
			// for user := range users {
			// 	newTitle := fmt.Sprintf("Titulo %s", strconv.Itoa(user))
			// 	mu.Lock()
			// 		users[user].Title = newTitle
			// 	mu.Unlock()
			// }
		}()
	}

	wg.Wait()

	return users, status, nil
}
