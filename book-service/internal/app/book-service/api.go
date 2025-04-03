package book_service

import (
	book_service "github.com/daniil4142/book-market/book-service/internal/service/book"
	desc "github.com/daniil4142/book-market/book-service/pkg/book-service"
)

type Implementation struct {
	desc.UnimplementedProductServiceServer
	bookService *book_service.Service
}

func NewProductService(bookService *book_service.Service) desc.BookServiceServer {
	return &Implementation{
		bookService: bookService,
	}
}
