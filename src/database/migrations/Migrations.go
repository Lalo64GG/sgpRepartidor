package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lalo64/sgp/src/database"
)

func main() {

	db, err := database.Connect()

	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}

	if err := Migrate(db); err != nil {
		log.Fatalf("Error ejecutando las migraciones: %v", err)
	}

}

func Migrate(db *sql.DB) error {


	sqlFiles := []string{
		"src/database/migrations/sql/000_create_client_table.sql",
		"src/database/migrations/sql/001_create_supplier_table.sql",
		"src/database/migrations/sql/002_create_deliverydriver_table.sql",
		"src/database/migrations/sql/003_create_delivery_table.sql",
		"src/database/migrations/sql/004_create_products_table.sql",
	}

	for _,file := range sqlFiles {

		err := executeSqlFiles(db, file)
		if err != nil {
			log.Fatalf("Error al ejecutar archivo %s: %v", file, err)
		}

	}

	log.Println("Migraciones aplicadas con éxito")
	return nil
}

func executeSqlFiles(db *sql.DB, filepath string) error {

	content, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("no se pudo leer el archivo %w", err)
	}

	_, err = db.Exec(string(content))
	if err != nil {
		return fmt.Errorf("error al ejecutar el archivo SQL: %w", err)
	}

	return nil

}
