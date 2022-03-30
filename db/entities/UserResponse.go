package entities

import (
	"gopkg.in/mgo.v2/bson"
)

type UserRespnse struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Id      int           `json:"id" bson:"id"`
	FName   string        `json:"fname" bson:"fname"`
	City    string        `json:"city" bson:"city"`
	Phone   int64         `json:"phone" bson:"phone"`
	Height  int           `json:"height"`
	Married bool          `json:"Married"`
}
