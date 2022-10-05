package customer

import (
	"github.com/globalsign/mgo"
)

const customerCollection = "customer"

type RepositoryMongo struct {
	DBName       string
	DBConnection *mgo.Session
}

func (r RepositoryMongo) CreatedCustomer(customer CustomerInfo) error {
	return r.DBConnection.DB(r.DBName).C(customerCollection).Insert(&customer)
}
