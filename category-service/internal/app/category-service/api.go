package category_service

import (
	"github.com/daniil4142/book-market/category-service/internal/service/category"
	desc "github.com/daniil4142/book-market/category-service/pkg/category-service"
)

type Implementation struct {
	desc.UnimplementedCategoryServiceServer

	categoryService category.Service
}

func NewCategoryService(categoryService category.Service) desc.CategoryServiceServer {
	return &Implementation{
		categoryService: categoryService,
	}
}
