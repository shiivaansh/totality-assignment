package controllers

import (
	"context"
	"encoding/json"
	"entities"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var user entities.UserRespnse
	collection := client.Database("Totality Corp").Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	id := c.Params("id")
	err := collection.FindOne(ctx, user{ID: id}).Decode(&person)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(person)
}
