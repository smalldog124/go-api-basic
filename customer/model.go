package customer

import "github.com/globalsign/mgo/bson"

type CustomerInfo struct {
	ID       bson.ObjectId
	Name     string
	LastName string
}
