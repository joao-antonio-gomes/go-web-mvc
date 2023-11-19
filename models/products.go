package models

import (
	"fmt"
	dbApp "github.com/joao-antonio-gomes/web-mvc/db"
)

type Product struct {
	Name, Description string
	Price             float64
	Quantity          int
}

func FindAllProducts() []Product {
	db := dbApp.ConnectPostgres()

	productsDb, err := db.Query("SELECT * FROM products")

	if err != nil {
		fmt.Println(err.Error())
	}

	var productSlice []Product
	for productsDb.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productsDb.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			fmt.Println(err.Error())
		}

		product := Product{Name: name, Description: description, Price: price, Quantity: quantity}
		productSlice = append(productSlice, product)
	}

	defer db.Close()
	return productSlice
}
