package main

import (
	"database/sql"
	"fmt"
	"github.com/joao-antonio-gomes/web-mvc/products"
	_ "github.com/lib/pq"
	template "html/template"
	"log"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := connectPostgres()
	defer db.Close()
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := connectPostgres()

	productsDb, err := db.Query("SELECT * FROM products")

	if err != nil {
		fmt.Println(err.Error())
	}

	var productSlice []products.Product
	for productsDb.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productsDb.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			fmt.Println(err.Error())
		}

		productSlice = append(productSlice, products.Product{Name: name, Description: description, Price: price, Quantity: quantity})
	}

	temp.ExecuteTemplate(w, "Index", productSlice)
}

func connectPostgres() *sql.DB {
	connStr := "user=postgres dbname=go-loja password=postgres host=localhost sslmode=disable port=5433"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
