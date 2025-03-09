package adapters

import (
	"database/sql"
	"fmt"

	"github.com/lalo64/sgp/src/database"
	"github.com/lalo64/sgp/src/deliverydriver/domain/entities"
)

type DriverRepositoryMysql struct{
	DB *sql.DB
} 

func NewDriverRepositoryMysql() (*DriverRepositoryMysql, error){
	db, err := database.Connect();

    if err != nil {
        return nil, err
    }

    return &DriverRepositoryMysql{DB: db}, nil
}


func (r *DriverRepositoryMysql) Create(driver entities.Driver) (entities.Driver, error) {
	query := `INSERT INTO DeliveryDriver (Name, Email, Password, fcm_token) VALUES (?, ?, ?, ? )`
	stmt, err := r.DB.Prepare(query)

	if err != nil {
		return entities.Driver{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(driver.Name, driver.Email, driver.Password, driver.FCM_TOKEN)

	if err != nil {
		return entities.Driver{}, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return entities.Driver{}, err
	}

	driver.ID = int(id)
	driver.Password = ""

	return driver, nil
}

func (r *DriverRepositoryMysql) GetByEmail(email string) (entities.Driver, error) {
	fmt.Println("GetByEmail", email)
	query := `SELECT ID, Name, Email, Password FROM deliverydriver WHERE Email =?`
    row := r.DB.QueryRow(query, email)

    var driver entities.Driver

    err := row.Scan(&driver.ID, &driver.Name, &driver.Email, &driver.Password)

    if err == sql.ErrNoRows {
        return entities.Driver{}, nil
    } else if err != nil {
        return entities.Driver{}, err
    }

    return driver, nil
}

func (r *DriverRepositoryMysql) GetById(id int64) (entities.Driver, error) {
	fmt.Println("GetDriverByIdUseCase", id)
    query := `SELECT ID, Name, Email, Password, fcm_token FROM deliverydriver WHERE ID =?`
    row := r.DB.QueryRow(query, id)

    var driver entities.Driver

    err := row.Scan(&driver.ID, &driver.Name, &driver.Email, &driver.Password, &driver.FCM_TOKEN)

    if err == sql.ErrNoRows {
        return entities.Driver{}, nil
    } else if err != nil {
        return entities.Driver{}, err
    }

    return driver, nil
}