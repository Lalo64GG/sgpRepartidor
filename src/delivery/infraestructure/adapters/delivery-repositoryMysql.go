package adapters

import (
	"database/sql"
	"github.com/lalo64/sgp/src/database"
	"github.com/lalo64/sgp/src/delivery/domain/entities"
)

type DeliveryRepositoryMysql struct {
	DB *sql.DB
}

func NewDeliveryRepositoryMysql() (*DeliveryRepositoryMysql, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	return &DeliveryRepositoryMysql{DB: db}, nil
}

func (r *DeliveryRepositoryMysql) Create(delivery entities.Delivery) (entities.Delivery, error) {
	query := `INSERT INTO Delivery (ClientID, DeliveryDate, ProductID, Status, SupplierID) 
	          VALUES (?, ?, ?, ?, ?)`

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return entities.Delivery{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(delivery.ClientID, delivery.DeliveryDate, delivery.ProductID, "Pending", delivery.SupplierID)

	if err != nil {
		return entities.Delivery{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entities.Delivery{}, err
	}


	delivery.DeliveryID = int(id) 

	return delivery, nil
}

func (r *DeliveryRepositoryMysql) AssignDriver(deliveryID int, driverID int) (entities.Delivery, error) {

	query := `UPDATE Delivery SET DriverID = ?, Status = "In Transit" WHERE DeliveryID = ? AND Status = "Pending"`

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return entities.Delivery{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(driverID, deliveryID)
	if err != nil {
		return entities.Delivery{}, err
	}
	
	return r.GetById(deliveryID)
}

// MarkAsDelivered cambia el estado del pedido a "Delivered"
func (r *DeliveryRepositoryMysql) MarkAsDelivered(deliveryID int) (entities.Delivery, error) {
	// Actualizar el estado de la entrega a "Delivered"
	query := `UPDATE Delivery SET Status = "Delivered" WHERE DeliveryID = ?`

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return entities.Delivery{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(deliveryID)
	if err != nil {
		return entities.Delivery{}, err
	}

	// Recuperar los datos actualizados del pedido
	return r.GetById(deliveryID)
}

// MarkAsCancelled cambia el estado del pedido a "Cancelled"
func (r *DeliveryRepositoryMysql) MarkAsCancelled(deliveryID int) (entities.Delivery, error) {
	// Actualizar el estado de la entrega a "Cancelled"
	query := `UPDATE Delivery SET Status = "Cancelled" WHERE DeliveryID = ?`

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return entities.Delivery{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(deliveryID)
	if err != nil {
		return entities.Delivery{}, err
	}

	// Recuperar los datos actualizados del pedido
	return r.GetById(deliveryID)
}

// GetByID obtiene una entrega por su ID
func (r *DeliveryRepositoryMysql) GetById(deliveryID int) (entities.Delivery, error) {
	query := `SELECT DeliveryID, ClientID, DeliveryDate, ProductID, Status, 
          IFNULL(DriverID, 0), SupplierID FROM Delivery WHERE DeliveryID = ?`


	row := r.DB.QueryRow(query, deliveryID)

	var delivery entities.Delivery
	err := row.Scan(&delivery.DeliveryID, &delivery.ClientID, &delivery.DeliveryDate, &delivery.ProductID, &delivery.Status, &delivery.DriverID, &delivery.SupplierID)
	if err != nil {
		if err == sql.ErrNoRows {
			return entities.Delivery{}, nil
		}
		return entities.Delivery{}, err
	}

	return delivery, nil
}


func (r *DeliveryRepositoryMysql) UpdateStatus(status string) (bool, error) {
	
	return false, nil
}

func (r *DeliveryRepositoryMysql) GetAllSupplierID(id int64)([]entities.Delivery, error){
	query := `SELECT DeliveryID, ClientID, DeliveryDate, Status, ProductID, IFNULL(DriverID, 0), SupplierID FROM Delivery WHERE SupplierID = ?`

	rows, err := r.DB.Query(query, id)
	if err != nil {
		return []entities.Delivery{}, err
	}
	defer rows.Close()

	var deliverys []entities.Delivery
	for rows.Next() {
		var delivery entities.Delivery
		var status []byte // Capturamos el status como []byte para evitar el error

		err := rows.Scan(
			&delivery.DeliveryID,
			&delivery.ClientID,
			&delivery.DeliveryDate,
			&status, // Recibimos el dato como []byte
			&delivery.ProductID,
			&delivery.DriverID,
			&delivery.SupplierID,
		)
		if err != nil {
			return []entities.Delivery{}, err
		}

		delivery.Status = string(status) // Convertimos []byte a string antes de asignarlo
		deliverys = append(deliverys, delivery)
	}

	return deliverys, nil
}

func (r* DeliveryRepositoryMysql) GetAllDriverID(id int64)([]entities.Delivery,error){
	query := `SELECT DeliveryID, ClientID, DeliveryDate, Status,ProductID, DriverID, SupplierID FROM Delivery WHERE DriverID = ?`

	rows, err := r.DB.Query(query, id)
	if err != nil {
		return []entities.Delivery{}, err
	}
	
	defer rows.Close()

	var deliverys []entities.Delivery
	for rows.Next() {
		var delivery entities.Delivery
		var status []byte // Capturamos el status como []byte para evitar el error

		err := rows.Scan(
			&delivery.DeliveryID,
			&delivery.ClientID,
			&delivery.DeliveryDate,
			&status, // Recibimos el dato como []byte
			&delivery.ProductID,
			&delivery.DriverID,
			&delivery.SupplierID,
		)
		if err != nil {
			return []entities.Delivery{}, err
		}

		delivery.Status = string(status) // Convertimos []byte a string antes de asignarlo
		deliverys = append(deliverys, delivery)
	}

	return deliverys, nil
}