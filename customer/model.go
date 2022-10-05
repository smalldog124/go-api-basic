package customer

import "github.com/globalsign/mgo/bson"

type CustomerInfo struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	LastName string        `json:"last_name" bson:"last_name"`
}
