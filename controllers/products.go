package controllers

import (
	"github.com/joao-antonio-gomes/web-mvc/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.FindAllProducts()
	err := temp.ExecuteTemplate(w, "Index", products)
	if err != nil {
		log.Panic("Erro ao abrir a página Index:", err)
	}
}

func New(w http.ResponseWriter, r *http.Request) {
	err := temp.ExecuteTemplate(w, "New", nil)
	if err != nil {
		log.Panic("Erro ao abrir a página New:", err)
	}
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

	name := r.FormValue("name")
	description := r.FormValue("description")
	price := r.FormValue("price")
	quantity := r.FormValue("quantity")

	priceFloat64, err := strconv.ParseFloat(price, 64)
	if err != nil {
		log.Panic("Erro na conversão do preço:", err)
	}

	quantityInt, err := strconv.Atoi(quantity)
	if err != nil {
		log.Panic("Erro na conversão da quantidade:", err)
	}

	product := models.Product{Name: name, Description: description, Price: priceFloat64, Quantity: quantityInt}

	models.InsertProduct(&product)
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	if productId == "" {
		log.Panic("Chamou DELETE PRODUTO sem passar ID")
		http.Redirect(w, r, "/", 301)
		return
	}
	productIdInt, err := strconv.Atoi(productId)
	if err != nil {
		log.Panic("Erro na conversão do ID:", err)
	}

	product := models.FindProductById(productIdInt)
	if product.Id == 0 {
		log.Panic("Produto não encontrado")
		http.Redirect(w, r, "/", 301)
		return
	}

	models.DeleteProduct(&product)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	if productId == "" {
		log.Panic("Chamou EDIT PRODUTO sem passar ID")
		http.Redirect(w, r, "/", 301)
		return
	}
	productIdInt, err := strconv.Atoi(productId)
	if err != nil {
		log.Panic("Erro na conversão do ID:", err)
	}

	product := models.FindProductById(productIdInt)
	if product.Id == 0 {
		log.Panic("Produto não encontrado")
		http.Redirect(w, r, "/", 301)
		return
	}

	err = temp.ExecuteTemplate(w, "Edit", product)
	if err != nil {
		log.Panic("Erro ao abrir a página Edit:", err)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

	name := r.FormValue("name")
	description := r.FormValue("description")
	price := r.FormValue("price")
	quantity := r.FormValue("quantity")
	productId := r.FormValue("id")
	productIdInt, err := strconv.Atoi(productId)
	if err != nil {
		log.Panic("Erro na conversão do ID:", err)
	}
	priceFloat64, err := strconv.ParseFloat(price, 64)
	if err != nil {
		log.Panic("Erro na conversão do preço:", err)
	}
	quantityInt, err := strconv.Atoi(quantity)
	if err != nil {
		log.Panic("Erro na conversão da quantidade:", err)
	}
	product := models.Product{Name: name, Description: description, Price: priceFloat64, Quantity: quantityInt, Id: productIdInt}

	productDb := models.FindProductById(productIdInt)
	if productDb.Id == 0 {
		log.Panic("Produto não encontrado")
		http.Redirect(w, r, "/", 301)
		return
	}

	models.UpdateProduct(product)

	err = temp.ExecuteTemplate(w, "Edit", product)
	if err != nil {
		log.Panic("Erro ao abrir a página Edit:", err)
	}
}
