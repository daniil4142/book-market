package book_service

import (
	book_service "github.com/daniil4142/book-market/book-service/internal/service/book"
	desc "github.com/daniil4142/book-market/book-service/pkg/book-service"
)

type Implementation struct {
	desc.UnimplementedBookServiceServer
	bookService *book_service.Service
}

func NewBookService(bookService *book_service.Service) desc.BookServiceServer {
	return &Implementation{
		bookService: bookService,
	}
}
