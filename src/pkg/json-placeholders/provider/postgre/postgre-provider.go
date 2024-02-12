package postgresProvider

import (
	"database/sql"
	"fmt"
	"github/leonardoas10/go-provider-pattern/src/pkg/common/env"
	"github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/models"
	"log"
	"strconv"
	"sync"
	"sync/atomic"

	_ "github.com/lib/pq"
)

type provider struct {
    db *sql.DB
}

// User represents the structure of the users table.
type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
    Country  string `json:"country"`
    Age      int    `json:"age"`
    Role     string `json:"role"`
    Hobby    string `json:"hobby"`
}

func NewProvider() *provider {
	postgresUser := env.GetEnvVariable("POSTGRES_USER")
	postgresPassword := env.GetEnvVariable("POSTGRES_PASSWORD")
	postgresHost := env.GetEnvVariable("POSTGRES_HOSTNAME")
	postgresPort := env.GetEnvVariable("POSTGRES_PORT")
	postgresDb := env.GetEnvVariable("POSTGRES_DB")
    // Construct connection string

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
	postgresUser, postgresPassword, postgresHost, postgresPort, postgresDb)

    // Open a connection to the database
    db, err := sql.Open("postgres", connectionString)
    if err != nil {
        log.Fatal(err)
    }

    // Check if the connection is successful
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to PostgreSQL database!")

    return &provider{db: db}
}

func (p *provider) Close() {
    err := p.db.Close()
    if err != nil {
        log.Fatal(err)
    }
}

func (p *provider) GetJsonPlaceHolders() ([]models.JsonPlaceHolder, int, error) {
    rows, err := p.db.Query("SELECT id, username, role FROM users")
    if err != nil {
        fmt.Println("Error querying database:", err)
        return nil, 500, err
    }
    defer rows.Close()

    var users []models.JsonPlaceHolder
    for rows.Next() {
        var user User
        if err := rows.Scan(&user.ID, &user.Username, &user.Role); err != nil {
            fmt.Println("Error scanning row:", err)
            return nil, 500, err
        }

        users = append(users, models.JsonPlaceHolder{UserId: strconv.Itoa(user.ID), Id: strconv.Itoa(user.ID), Title: user.Role, Completed: false})
    }

    fmt.Printf("Number of users: %d\n", len(users))

    return users, 200, nil
}
func (p *provider) GetJsonPlaceHolder(id string) (models.JsonPlaceHolder, int, error) { 
	var user User

    // Execute the SQL statement
    row := p.db.QueryRow("SELECT id, role FROM users WHERE id = $1", id)

    // Scan the result into the user struct
    if err := row.Scan(&user.ID, &user.Role); err != nil {
        if err == sql.ErrNoRows {
            fmt.Println("User not found")
            return models.JsonPlaceHolder{}, 404, err
        }
        fmt.Println("Error querying database:", err)
        return models.JsonPlaceHolder{}, 500, err
    }

    // Return the user data
    return models.JsonPlaceHolder{
        UserId:    strconv.Itoa(user.ID),
        Id:        strconv.Itoa(user.ID),
        Title:     user.Role,
        Completed: false,
    }, 200, nil
}

func (p *provider) UpdateJsonPlaceHolder(jsonPlacerHolder models.UpdateJsonPlaceHolder) (models.JsonPlaceHolder, int, error) {
	   // Convert ID to integer
    id, err := strconv.Atoi(jsonPlacerHolder.Id)
    if err != nil {
        fmt.Println("Error converting string to integer:", err)
        return models.JsonPlaceHolder{}, 400, err
    }

    // Execute the SQL statement
    _, err = p.db.Exec("UPDATE users SET role = $1 WHERE id = $2", jsonPlacerHolder.Title, id)
    if err != nil {
        fmt.Println("Error updating document:", err)
        return models.JsonPlaceHolder{}, 500, err
    }

    // Retrieve the updated document
    updatedUser, status, err := p.GetJsonPlaceHolder(jsonPlacerHolder.Id)
    if err != nil {
        fmt.Println("Error retrieving updated document:", err)
        return models.JsonPlaceHolder{}, status, err
    }

    return updatedUser, 200, nil
}

func (p *provider) ConcurrentChangeTitles() ([]models.JsonPlaceHolder, int, error) {
	// Fetch users from the database
    users, status, err := p.GetJsonPlaceHolders()
    if err != nil {
        return nil, status, err
    }

    // Create a WaitGroup to wait for all goroutines to finish
    var wg sync.WaitGroup

    // Use a mutex to safely modify the users slice concurrently
    // var mu sync.Mutex

    var counter int64

    // Launch 5 goroutines to concurrently modify titles
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

                newTitle := fmt.Sprintf("Titulo PostgreSQL %d", userIndex+1)

                // No need for mutex when updating a specific index in the slice
                users[userIndex].Title = newTitle
            }

            // Alternatively, use approach with Mutex Locks
            // for userIndex := range users {
            // 	newTitle := fmt.Sprintf("Titulo %d", userIndex+1)
            // 	mu.Lock()
            // 	users[userIndex].Title = newTitle
            // 	mu.Unlock()
            // }
        }()
    }

    // Wait for all goroutines to finish
    wg.Wait()

    return users, status, nil
}
