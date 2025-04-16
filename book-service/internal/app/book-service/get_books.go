package book_service

import (
	"context"

	desc "github.com/daniil4142/book-market/book-service/pkg/book-service"
)

func (i *Implementation) GetBook(ctx context.Context, req *desc.GetBookRequest) (*desc.GetBookResponse, error) {
	books, err := i.bookService.GetBooks(ctx, req.GetBookIds)
	if err != nil {
		return nil, err
	}

	br := make([]*desc.Book, len(books))
	for idx := range books {
		br[idx] = convertBookToPb(&books[idx])
	}

	return &desc.GetBookResponse{
		Books: br,
	}, nil
}
