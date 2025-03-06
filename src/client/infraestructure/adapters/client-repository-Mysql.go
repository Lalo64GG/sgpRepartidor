package adapters

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lalo64/sgp/src/client/domain/entities"
	"github.com/lalo64/sgp/src/database"
)

type ClientRepositoryMysql struct {
	DB *sql.DB
}

func NewClientRepositoryMysql() (*ClientRepositoryMysql, error) {
	db, err := database.Connect();

	if err != nil {
		return nil, err
	}

	return &ClientRepositoryMysql{DB: db}, nil
}


func (r *ClientRepositoryMysql) Create(client entities.Client) (entities.Client, error) {
	query := `INSERT INTO Client (Name, Email, Password, Address) VALUES (?, ?, ?, ?)`
	stmt, err := r.DB.Prepare(query)

	if err != nil {
		log.Fatal(err, 1)
	}
	defer stmt.Close()

	result, err := stmt.Exec(client.Name, client.Email, client.Password, client.Address)

	if err != nil {
		return entities.Client{}, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return entities.Client{}, err
	}

	client.ID = int(id)
	client.Password = ""

	return client, nil
}

func (r *ClientRepositoryMysql) CheckEmail(email string) (bool, error ){
	query := `SELECT EXISTS (SELECT 1 FROM Client WHERE Email = ?) AS existe`

	stmt, err := r.DB.Prepare(query)

	if err != nil {
		return false, err
	}

	defer stmt.Close()


	var existe bool
	log.Print(existe, email)
	err = stmt.QueryRow(email).Scan(&existe)

	if err != nil {
		log.Print(err, 2)
		return false, err
	}

	return existe, nil
}

func (r *ClientRepositoryMysql) GetAll(limit, page int64, orderBy, orderDir string ) ([]entities.Client, error) {
	offset := limit * (page -1)

	query := fmt.Sprintf("SELECT Id, Name, Email, Address FROM Client ORDER BY %s %s LIMIT ? OFFSET ?", orderBy, orderDir)
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(limit, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var clients []entities.Client
	for rows.Next() {
		var client entities.Client
		err := rows.Scan(&client.ID, &client.Name, &client.Email, &client.Address)
		if err != nil {
			return nil, err
		}

		clients = append(clients, client)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Rows error: %s", err)
		return nil, err
	}


	return clients, nil
} 

func (r *ClientRepositoryMysql) GetByEmail(email string) (entities.Client, error) {
	query := `SELECT Id, Name, Email, password, Address FROM Client WHERE Email = ?`
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return entities.Client{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(email)

	var client entities.Client

	err = row.Scan(&client.ID, &client.Name, &client.Email, &client.Password, &client.Address)

	if err != nil {
		return entities.Client{}, err
	}

	return client, nil

}


func (r *ClientRepositoryMysql) GetById(id int) (entities.Client, error) {
	query := `SELECT Id, Email, Password, Name FROM users WHERE Id = ?`
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return entities.Client{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	var client entities.Client

	err = row.Scan(&client.ID, &client.Name, &client.Email, &client.Password, &client.Address)

	if err != nil {
		return entities.Client{}, err
	}

	return client, nil

}


func (r *ClientRepositoryMysql) UpdateUsername(client entities.Client){
	query := `UPDATE Client SET Name =? WHERE id =?`

    stmt, err := r.DB.Prepare(query)

    if err != nil {
        log.Fatal(err, 1)
    }
    defer stmt.Close()

    _, err = stmt.Exec(client.Name, client.ID)

    if err != nil {
        log.Fatal(err, 2)
    }
}

func (r *ClientRepositoryMysql) Delete(id int) (bool, error) {
	query := `DELETE FROM Client WHERE Id = ?`

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return false, err
	}

	return true, nil
}
