package book_service

import "context"

var nextID int64 = 1

type repository struct{}

func newRepo() IRepository {
	return repository{}
}

func (r repository) SaveBook(ctx context.Context, book *Book) error {
	book.ID = nextID

	nextID++

	return nil
}
