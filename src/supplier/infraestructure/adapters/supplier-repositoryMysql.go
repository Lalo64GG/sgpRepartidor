package adapters

import (
	"database/sql"
	"fmt"

	"github.com/lalo64/sgp/src/database"
	"github.com/lalo64/sgp/src/supplier/domain/entities"
)


type SupplierRepositoryMysql struct {
	DB *sql.DB
}

func NewSupplierRepositoryMysql() (*SupplierRepositoryMysql, error) {
	db, err := database.Connect();

	if err != nil {
		return nil, err
	}

	return &SupplierRepositoryMysql{DB: db}, nil
}


func (r *SupplierRepositoryMysql) Create(supplier entities.Supplier) (entities.Supplier, error) {
	query := `INSERT INTO Supplier (Name, Email, Password, Address, ContactInfo) VALUES (?, ?, ?, ?, ?)`
	stmt, err := r.DB.Prepare(query)

	if err != nil {
		return entities.Supplier{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(supplier.Name, supplier.Email, supplier.Password, supplier.Address, supplier.ContactInfo)

	if err != nil {
		return entities.Supplier{}, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return entities.Supplier{}, err
	}

	supplier.ID = int(id)
	supplier.Password = ""

	return supplier, nil
}

func (r *SupplierRepositoryMysql) CheckEmail(email string) (bool, error ){
	query := `SELECT EXISTS (SELECT 1 FROM Supplier WHERE email = ?) AS existe`
	stmt, err := r.DB.Prepare(query)

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	var existe bool
	err = stmt.QueryRow(email).Scan(&existe)
	if err != nil {
        return false, err
    }
	return existe, nil
}

func (r *SupplierRepositoryMysql) GetAll(limit, page int64, orderBy, orderDir string ) ([]entities.Supplier, error) {
	offset := limit * (page -1)

	query := fmt.Sprintf("SELECT * FROM Supplier ORDER BY %s %s LIMIT ? OFFSET ?", orderBy, orderDir)

	stmt, err := r.DB.Prepare(query)
	if err != nil {
        return nil, err
    }

	defer stmt.Close()

	rows, err := stmt.Query(limit, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var suppliers []entities.Supplier
	for rows.Next() {
		var supplier entities.Supplier
		err = rows.Scan(&supplier.ID, &supplier.Name, &supplier.Email, &supplier.Password, &supplier.Address, &supplier.ContactInfo)
		if err != nil {
			return nil, err
		}
		suppliers = append(suppliers, supplier)
	}

	if err := rows.Err(); err != nil{
		return nil, err
	}

	return suppliers, nil
}


func (r *SupplierRepositoryMysql) GetByEmail(email string) (entities.Supplier, error) {
	query := `SELECT Id, Name, Email, Password, Address FROM Supplier WHERE email = ?`
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return entities.Supplier{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(email)

	var client entities.Supplier

	err = row.Scan(&client.ID, &client.Name, &client.Email, &client.Password, &client.Address)

	if err != nil {
		return entities.Supplier{}, err
	}

	return client, nil

}

func (r *SupplierRepositoryMysql) GetById(id int64) (entities.Supplier, error) {
	query := `SELECT id, email, password, username FROM users WHERE id =?`
    stmt, err := r.DB.Prepare(query)
    if err != nil {
        return entities.Supplier{}, err
    }
    defer stmt.Close()

    row := stmt.QueryRow(id)

    var client entities.Supplier

    err = row.Scan(&client.ID, &client.Name, &client.Email, &client.Password, &client.Address)

    if err != nil {
        return entities.Supplier{}, err
    }

    return client, nil
}