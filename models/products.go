package models

import (
	dbApp "github.com/joao-antonio-gomes/web-mvc/db"
	"log"
)

type Product struct {
	Name, Description string
	Price             float64
	Quantity          int
}

func FindAllProducts() []Product {
	log.Println("Buscando todos os produtos")
	db := dbApp.ConnectPostgres()

	statement := "SELECT * FROM products"
	log.Println("Executando query:", statement)
	productsDb, err := db.Query(statement)

	if err != nil {
		log.Panicln(err.Error())
	}

	var productSlice []Product
	for productsDb.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productsDb.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			log.Panicln(err.Error())
		}

		product := Product{Name: name, Description: description, Price: price, Quantity: quantity}
		productSlice = append(productSlice, product)
	}

	defer db.Close()
	return productSlice
}
