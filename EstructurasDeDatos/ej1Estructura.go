package main

import "fmt"

type Product struct {
	name        string
	description string
	category    string
	price       float32
}

var products []Product

func (this Product) Save() {
	products = append(products, this)
}

func (this Product) GetAll() {
	for _, singleProdduct := range products {
		fmt.Printf("El producto %s, %s, de la categoria %s, tiene un precio de %0.2f\n", singleProdduct.name, singleProdduct.description, singleProdduct.category, singleProdduct.price)
	}
}

func (this Product) GetById(id int) (productoID Product) {
	productoID = products[id]
	return
}
