package mongoProvider

import (
	"context"
	"fmt"
	"github/leonardoas10/go-provider-pattern/src/pkg/common/env"
	"github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/models"
	"log"
	"strconv"
	"sync"
	"sync/atomic"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type provider struct {
	client *mongo.Client
	database *mongo.Database
}


func NewProvider() (*provider)  {
	mongoUsername := env.GetEnvVariable("MONGO_USERNAME")
	mongoPassword := env.GetEnvVariable("MONGO_PASSWORD")
	mongoHostname := env.GetEnvVariable("MONGO_HOSTNAME")
	mongoPort := env.GetEnvVariable("MONGO_PORT")
	mongoDB := env.GetEnvVariable("MONGO_DB")

	// Construct the MongoDB connection string
	url := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin",
		mongoUsername, mongoPassword, mongoHostname, mongoPort, mongoDB)
	
	fmt.Println("Url => ", url)

	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	database := client.Database(mongoDB)

	return &provider{client: client, database: database}
}

type JsonPlaceHolderFromMongo struct {
	ID   primitive.ObjectID    `bson:"_id"`
	Role string `bson:"role"`
}


func (p *provider) GetJsonPlaceHolders() ([]models.JsonPlaceHolder, int, error) {
    collection := p.database.Collection(env.GetEnvVariable("MONGO_USER_COLLECTION"))
    cursor, err := collection.Find(context.Background(), bson.D{{}})

    if err != nil {
        fmt.Println("Error finding documents:", err)
        return nil, 500, err
    }
    defer cursor.Close(context.Background())

    var users []models.JsonPlaceHolder
    for cursor.Next(context.Background()) {
        var user JsonPlaceHolderFromMongo
        if err := cursor.Decode(&user); err != nil {
            fmt.Println("Error decoding document:", err)
            return nil, 500, err
        }
		fmt.Println("user => ", user)

		if err != nil {
            fmt.Println("Error converting ObjectID to int:", err)
            return nil, 500, err
        }

        users = append(users, models.JsonPlaceHolder{UserId:  user.ID.Hex(), Id: user.ID.Hex(), Title: user.Role, Completed: false})
    }

    fmt.Printf("Number of users: %d\n", len(users))
	fmt.Println("Users => ", users)

    return users, 200, nil
}

func (p *provider) GetJsonPlaceHolder(id string) (models.JsonPlaceHolder, int, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
		fmt.Println("Error converting string to ObjectID:", err)
        return models.JsonPlaceHolder{}, 400, err
    }
	collection := p.database.Collection(env.GetEnvVariable("MONGO_USER_COLLECTION"))

    filter := bson.D{{Key: "_id", Value: objectID}}
    result := collection.FindOne(context.Background(), filter)

    if result.Err() != nil {
        fmt.Println("Error finding document:", result.Err())
        return models.JsonPlaceHolder{}, 404, result.Err()
    }

    var user JsonPlaceHolderFromMongo
    if err := result.Decode(&user); err != nil {
        fmt.Println("Error decoding document:", err)
        return models.JsonPlaceHolder{}, 500, err
    }

    return models.JsonPlaceHolder{
        UserId:    user.ID.Hex(),
        Id:        user.ID.Hex(),
        Title:     user.Role,
        Completed: false,
    }, 200, nil
}

func (p *provider) UpdateJsonPlaceHolder(jsonPlacerHolder models.UpdateJsonPlaceHolder) (models.JsonPlaceHolder, int, error) {
	objectID, err := primitive.ObjectIDFromHex(jsonPlacerHolder.Id)

    if err != nil {
		fmt.Println("Error converting string to ObjectID:", err)
        return models.JsonPlaceHolder{}, 400, err
    }
	collection := p.database.Collection(env.GetEnvVariable("MONGO_USER_COLLECTION"))

    filter := bson.D{{Key: "_id", Value: objectID}}
    update := bson.D{{Key: "$set", Value: bson.D{{Key: "role", Value: jsonPlacerHolder.Title}}}}

    result, err := collection.UpdateOne(context.Background(), filter, update)

    if err != nil {
        fmt.Println("Error updating document:", err)
        return models.JsonPlaceHolder{}, 500, err
    }

    if result.ModifiedCount == 0 {
		fmt.Println("No documents matched the filter, indicating that the document with the given ID was not found")
        return models.JsonPlaceHolder{}, 404, nil
    }

    // Retrieve the updated document
    updatedResult := collection.FindOne(context.Background(), filter)

    if updatedResult.Err() != nil {
        fmt.Println("Error finding updated document:", updatedResult.Err())
        return models.JsonPlaceHolder{}, 500, updatedResult.Err()
    }

    var updatedUser JsonPlaceHolderFromMongo
    if err := updatedResult.Decode(&updatedUser); err != nil {
        fmt.Println("Error decoding updated document:", err)
        return models.JsonPlaceHolder{}, 500, err
    }

    return models.JsonPlaceHolder{
        UserId:    updatedUser.ID.Hex(),
        Id:        updatedUser.ID.Hex(),
        Title:     updatedUser.Role,
        Completed: false, // Assuming a default value
    }, 200, nil
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

