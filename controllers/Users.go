package controllers

import (
	"context"
	"encoding/json"

	"gopkg.in/mgo.v2/bson"

	"log"
	"net/http"
	"os"
	"totality-assignment/mod/db/entities"

	// entities "gitlab.com/shiiivaaansh/totality-corp/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectToDB() (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://shivansh:shivansh@cluster0.crdbf.mongodb.net/users?retryWrites=true&w=majority"))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return client, nil

}

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGOURI")
}
func GetCollection(DbName string, CollectionName string) (*mongo.Collection, error) {

	client, _ := ConnectToDB()

	collection := client.Database(DbName).Collection(CollectionName)

	return collection, nil
}
func GetUsers(c *gin.Context) {
	collection, _ := GetCollection("totality", "users")
	userId := c.Param("name")
	var results entities.UserRespnse
	err := collection.FindOne(context.Background(), bson.M{"id": userId}).Decode(results)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error")
		return
	}
	log.Println(results)
	c.JSON(http.StatusOK, results)
}

func CreatePerson(c *gin.Context) {
	collection, err := GetCollection("totality", "users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error")
		return
	}
	var user = new(entities.UserRespnse)

	if err := c.BindJSON(user); err != nil {

		c.JSON(http.StatusOK, "ERRORRR from JSON!!!!!!")
		return
	}
	res, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error")
		return
	}

	response, _ := json.Marshal(res)
	c.JSON(http.StatusOK, response)
}
