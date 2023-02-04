package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/lucasrmp/web-application-studies/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func parseProductFromFormValues(request *http.Request) models.Product {
	name := request.FormValue("name")
	priceString := request.FormValue("price")
	description := request.FormValue("description")
	quantityString := request.FormValue("quantity")

	price, err := strconv.ParseFloat(priceString, 64)
	if err != nil {
		panic(err)
	}

	quantity, err := strconv.Atoi(quantityString)
	if err != nil {
		panic(err)
	}

	return models.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity,
	}
}

func Index(response http.ResponseWriter, request *http.Request) {
	products := models.FindAllProducts()
	templates.ExecuteTemplate(response, "Index", products)
}

func New(response http.ResponseWriter, request *http.Request) {
	templates.ExecuteTemplate(response, "New", nil)
}

func Insert(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		product := parseProductFromFormValues(request)
		models.CreateProduct(product)
	}

	http.Redirect(response, request, "/", http.StatusMovedPermanently)
}

func Delete(response http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.URL.Query().Get("id"))
	if err != nil {
		panic(err)
	}

	models.DeleteProduct(id)
	http.Redirect(response, request, "/", http.StatusMovedPermanently)
}

func Edit(response http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.URL.Query().Get("id"))
	if err != nil {
		panic(err)
	}

	product := models.FindProduct(id)
	templates.ExecuteTemplate(response, "Edit", product)
}

func Update(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		id, err := strconv.Atoi(request.FormValue("id"))
		if err != nil {
			panic(err)
		}

		product := parseProductFromFormValues(request)
		product.Id = id
		models.UpdateProduct(product)
	}

	http.Redirect(response, request, "/", http.StatusMovedPermanently)
}
