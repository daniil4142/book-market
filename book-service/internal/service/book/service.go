package book_service

import (
	"context"
	"errors"

	category_service "github.com/daniil4142/book-market/category-service/pkg/category-service"
	"github.com/jmoiron/sqlx"
)

var ErrWrongCategory = errors.New("category does not exist")

//go:generate mockgen -package=book_service -destination=service_mocks_test.go -self_package=github.com/daniil4142/book-market/book-service/internal/service/book . IRepository,ICategoryClient

type IRepository interface {
	SaveBook(ctx context.Context, book *Book) (int64, error)
	GetBooks(ctx context.Context, ids []int64) ([]Book, error)
	DeleteBooks(ctx context.Context, ids []int64) error
}

type ICategoryClient interface {
	IsCategoryExists(ctx context.Context, categoryID int64) (ok bool, err error)
}

type Service struct {
	repo   IRepository
	client ICategoryClient
}

func NewService(grpcClient category_service.CategoryServiceClient, db *sqlx.DB) *Service {
	return &Service{
		repo:   newRepo(db),
		client: newClient(grpcClient),
	}
}

func (s *Service) CreateBook(
	ctx context.Context,
	name string,
	categoryID int64,
) (*Book, error) {
	exists, err := s.client.IsCategoryExists(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, ErrWrongCategory
	}

	book := &Book{
		Name:       name,
		CategoryID: categoryID,
	}

	if id, err := s.repo.SaveBook(ctx, book); err != nil {
		return nil, err
	} else {
		book.ID = id
	}

	return book, nil
}

func (s *Service) DeleteBooks(ctx context.Context, ids []int64) error {
	return s.repo.DeleteBooks(ctx, ids)
}

func (s *Service) GetBooks(ctx context.Context, ids []int64) ([]Book, error) {
	return s.repo.GetBooks(ctx, ids)
}
