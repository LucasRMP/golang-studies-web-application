package models

import (
	"github.com/lucasrmp/web-application-studies/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func FindAllProducts() []Product {
	db := db.Connect()
	defer db.Close()

	queryResult, err := db.Query("SELECT * FROM products order by id asc")
	if err != nil {
		panic(err.Error())
	}

	products := []Product{}
	for queryResult.Next() {
		var product Product
		err = queryResult.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Quantity)
		if err != nil {
			panic(err.Error())
		}
		products = append(products, product)
	}

	return products
}

func CreateProduct(product Product) {
	db := db.Connect()
	defer db.Close()

	insert, err := db.Prepare("INSERT INTO products(name, description, price, quantity) VALUES($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insert.Exec(product.Name, product.Description, product.Price, product.Quantity)
}

func DeleteProduct(id int) {
	db := db.Connect()
	defer db.Close()

	delete, err := db.Prepare("DELETE FROM products WHERE id = $1")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)
}

func FindProduct(id int) Product {
	db := db.Connect()
	defer db.Close()

	queryResult, err := db.Query("SELECT * FROM products WHERE id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	var product Product
	for queryResult.Next() {
		err = queryResult.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Quantity)
		if err != nil {
			panic(err.Error())
		}
	}

	return product
}

func UpdateProduct(product Product) {
	db := db.Connect()
	defer db.Close()

	update, err := db.Prepare("UPDATE products SET name = $1, description = $2, price = $3, quantity = $4 WHERE id = $5")
	if err != nil {
		panic(err.Error())
	}

	update.Exec(product.Name, product.Description, product.Price, product.Quantity, product.Id)
}
