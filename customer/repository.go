package customer

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

const customerCollection = "customer"

type RepositoryMongo struct {
	DBName       string
	DBConnection *mgo.Session
}

func (r RepositoryMongo) CreatedCustomer(customer CustomerInfo) error {
	customer.ID = bson.NewObjectId()
	return r.DBConnection.DB(r.DBName).C(customerCollection).Insert(&customer)
}
