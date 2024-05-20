package product

import "hexagonal-api/internal/category"

type Product struct {
	Id          string
	Name        string
	Price       float32
	Description string
}

type ProductCategory struct {
	Id          string
	Name        string
	Price       float32
	Description string
	Category    []category.Category
}
