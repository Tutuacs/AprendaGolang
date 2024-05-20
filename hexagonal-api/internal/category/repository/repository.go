package categoryRepository

import "hexagonal-api/internal/category"

type Category struct {
	Category category.Category
}

func CreateCategory(data Category) (category Category, err error) {
	return
}

func FindCategory(id string) (category Category, err error) {

	return
}

func ListCategorys() (category []Category, err error) {

	return
}

func UpdateCategory(id string, data Category) (category Category, err error) {

	return
}

func DeleteCategory(id string) (category Category, err error) {

	return
}
