package cat_repository

import (
	"context"

	"github.com/daniil4142/book-market/category-service/internal/service/category"
)

var categories = category.Categories{
	{
		ID:   1,
		Name: "Fantasy",
	},
	{
		ID:   2,
		Name: "Sci-fi",
	},
	{
		ID:   3,
		Name: "Romance novel",
	},
}

type Repository struct{}

func New( /* db */ ) *Repository {
	return &Repository{}
}

func (Repository) FindAllCategories(_ context.Context) (category.Categories, error) {
	return categories, nil
}
