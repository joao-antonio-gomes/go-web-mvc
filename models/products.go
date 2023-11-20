package models

import (
	dbApp "github.com/joao-antonio-gomes/web-mvc/db"
	"log"
)

type Product struct {
	Name, Description string
	Price             float64
	Id, Quantity      int
}

func FindAllProducts() []Product {
	log.Println("Buscando todos os produtos")
	db := dbApp.ConnectPostgres()

	statement := "SELECT * FROM products ORDER BY id"
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

		product := Product{Id: id, Name: name, Description: description, Price: price, Quantity: quantity}
		productSlice = append(productSlice, product)
	}

	defer db.Close()
	return productSlice
}

func InsertProduct(p *Product) {
	log.Println("Inserindo novo produto")

	db := dbApp.ConnectPostgres()
	statement := "INSERT INTO products(name, description, price, quantity) values ($1, $2, $3, $4)"
	log.Println("Executando query:", statement)
	insereDadosNoBanco, err := db.Prepare(statement)
	if err != nil {
		log.Panic("Houve um erro ao construir a query de inserção de produtos:", err)
	}

	_, err = insereDadosNoBanco.Exec(p.Name, p.Description, p.Price, p.Quantity)
	if err != nil {
		log.Panic("Houve um erro ao inserir o produto:", p, err)
	}

	defer db.Close()
}

func FindProductById(id int) Product {
	log.Println("Buscando produto pelo id: ", id)

	db := dbApp.ConnectPostgres()
	defer db.Close()

	statement := "SELECT * FROM products WHERE id = $1"
	log.Println("Executando query:", statement)
	insereDadosNoBanco, err := db.Prepare(statement)
	if err != nil {
		log.Panic("Houve um erro ao construir a query de busca de produto por id:", id, err)
	}

	rows, err := insereDadosNoBanco.Query(id)
	if err != nil {
		log.Panic("Houve um erro ao buscar o produto pelo id:", id, err)
	}

	var product Product
	for rows.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = rows.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			log.Panicln(err.Error())
			return Product{}
		}

		product = Product{Id: id, Name: name, Description: description, Price: price, Quantity: quantity}
	}
	return product
}

func DeleteProduct(p *Product) {
	log.Println("Deletando produto de id: ", p.Id)

	db := dbApp.ConnectPostgres()
	defer db.Close()

	statement := "DELETE FROM products WHERE id = $1"
	log.Println("Executando query:", statement)
	deleteDadosNoBanco, err := db.Prepare(statement)
	if err != nil {
		log.Panic("Houve um erro ao construir a query de delete de produto por id:", p.Id, err)
	}

	_, err = deleteDadosNoBanco.Exec(p.Id)
	if err != nil {
		log.Panic("Houve um erro ao buscar o produto pelo id:", p.Id, err)
		return
	}

	log.Println("Produto deletado de id:", p.Id)
}

func UpdateProduct(p Product) {
	log.Println("Atualizando produto de id: ", p.Id)

	db := dbApp.ConnectPostgres()
	defer db.Close()

	statement := "UPDATE products SET name=$1, description=$2, price=$3, quantity=$4 WHERE id=$5"
	log.Println("Executando query:", statement)
	atualizaDadosNoBanco, err := db.Prepare(statement)
	if err != nil {
		log.Panic("Houve um erro ao construir a query de update de produto por id:", p.Id, err)
	}

	_, err = atualizaDadosNoBanco.Exec(p.Name, p.Description, p.Price, p.Quantity, p.Id)
	if err != nil {
		log.Panic("Houve um erro ao atualizar o produto pelo id:", p.Id, err)
	}
}
