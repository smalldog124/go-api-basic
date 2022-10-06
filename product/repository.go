package product

import "github.com/jmoiron/sqlx"

type RepositoryPg struct {
	DBConnection *sqlx.DB
}

type Product struct {
	ID    int     `json:"id" db:"id"`
	Name  string  `json:"name" db:"name"`
	Price float64 `json:"price" db:"price"`
}

func (r RepositoryPg) CreatedProduct(product Product) error {
	_, err := r.DBConnection.NamedExec("INSERT INTO product (name,price) VALUES (:name,:price)", product)
	return err
}
