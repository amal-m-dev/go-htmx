package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Product table
// - ID
// - Name
// - Price
// - Available
// - Date Created

type Product struct {
	Name      string
	Price     float64
	Available bool
}

func main() {
	fmt.Println("Hello, World!")
	connStr := "postgres://postgres:@maLrk0N@localhost:5432/gopgtest?sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	createProductTable(db)

	// product := Product{
	// 	"Book", 9.99, true}

	// pk := insertProduct(db, product)

	var name string
	var available bool
	var price float64

	pk := 10

	query := `SELECT name, price, available FROM product WHERE id = $1`

	err = db.QueryRow(query, pk).Scan(&name, &price, &available)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found for ID =", pk)
			return
		}

		log.Fatal(err)
	}

	fmt.Println("Name:", name)
	fmt.Println("Price:", price)
	fmt.Println("Available:", available)
}
func createProductTable(db *sql.DB) {

	query := `CREATE TABLE IF NOT EXISTS product (
	id SERIAL PRIMARY KEY, 
	name VARCHAR(100) NOT NULL, 
	price NUMERIC(6,2) NOT NULL, 
	available BOOLEAN, 
	date_created timestamp DEFAULT NOW()
	)`

	_, err := db.Exec(query)

	if err != nil {
		log.Fatal(err)
	}

}

func insertProduct(db *sql.DB, product Product) int {
	query := `INSERT INTO product (name, price, available)
	VALUES ($1, $2, $3) RETURNING id`

	var pk int
	err := db.QueryRow(query, product.Name, product.Price, product.Available).Scan(&pk)

	if err != nil {
		log.Fatal(err)
	}
	return pk
}
