package adapters

import (
	"database/sql"
	"log"

	"github.com/lalo64/sgp/src/database"
	"github.com/lalo64/sgp/src/products/domain/entities"
)

type ProductRepositoryMysql struct {
	DB *sql.DB
}

func NewProductRepositoryMysql() (*ProductRepositoryMysql, error) {
	db, err := database.Connect()

	if err != nil {
		return nil, err
	}

	return &ProductRepositoryMysql{DB: db}, nil
}

func (r *ProductRepositoryMysql) Create(product entities.Products) (entities.Products, error) {
	query := `INSERT INTO Product (name, price, supplier_id) VALUES (?, ?, ?)`
	stmt, err := r.DB.Prepare(query)

	if err != nil {
		return entities.Products{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(product.Name, product.Price, product.Supplier_Id)

	if err != nil {
		return entities.Products{}, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return entities.Products{}, err
	}

	product.ID = int(id)

	return product, nil
}

func (r *ProductRepositoryMysql) GetAllByIdSupplier(id int64) ([]entities.Products, error) {
	query := `SELECT id, name, price, supplier_id  FROM Product WHERE supplier_id = ?`
	stmt, err := r.DB.Prepare(query)

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var products []entities.Products
	for rows.Next() {
		var product entities.Products
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Supplier_Id)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil	
}