package book_service

import (
	"context"

	desc "github.com/daniil4142/book-market/book-service/pkg/book-service"
)

func (i *Implementation) DeleteBook(ctx context.Context, req *desc.DeleteBookRequest) (*desc.DeleteBookResponse, error) {
	err := i.bookService.DeleteBooks(ctx, req.GetBookIds)
	if err != nil {
		return nil, err
	}

	return &desc.DeleteBookResponse{}, nil
}
