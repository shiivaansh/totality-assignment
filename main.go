package main

import (

	// "errors"
	// "go/parser"

	"log"
	"os"
	"totality-assignment/mod/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// errors "github.com/haproxytech/config-parser/v4/errors"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	controllers.ConnectToDB()
	router := gin.Default()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	router.GET("/users/:name", controllers.GetUsers)
	router.POST("/users", controllers.CreatePerson)
	router.Run("localhost:8080")
}

// func ConnectToDB() *mongo.Client {
// 	c, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	err = c.Connect(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err2 := c.Ping(ctx, nil)
// 	if err2 != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Connected to MongoDB database to fetch the users")
// 	return c
// }

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGOURI")
}

// func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
// 	collection := client.Database("Totality Corp").Collection(collectionName)
// 	return collection
// }
