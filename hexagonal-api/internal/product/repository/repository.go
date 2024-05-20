package productRepository

import "hexagonal-api/internal/product"

type Product struct {
	Product product.Product
}

func CreateProduct(data Product) (product Product, err error) {
	return
}

func FindProduct(id string) (product Product, err error) {
	return
}

func ListProducts() (product []Product, err error) {
	return
}

func UpdateProduct(id string, data Product) (product Product, err error) {
	return
}

func DeleteProduct(id string) (product Product, err error) {
	return
}
