package book_service

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func setup(t *testing.T) (*Service, *MockIRepository, *MockICategoryClient) {
	ctrl := gomock.NewController(t)
	repo := NewMockIRepository(ctrl)
	client := NewMockICategoryClient(ctrl)

	service := &Service{
		repo:   repo,
		client: client,
	}

	return service, repo, client
}

func TestCreateBook_Success_ReturnName(t *testing.T) {
	service, repo, client := setup(t)

	repo.EXPECT().
		SaveBook(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, book *Book) error {
			book.ID = 124
			return nil
		})

	client.EXPECT().
		IsCategoryExists(gomock.Any(), int64(4)).
		Return(true, nil)

	book, err := service.CreateBook(
		context.Background(),
		"new book",
		4,
	)

	require.Nil(t, err)
	require.Equal(t, "new book", book.Name)
}

func TestCreateBook_Success_ReturnCategoryID(t *testing.T) {
	service, repo, client := setup(t)

	repo.EXPECT().
		SaveBook(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, book *Book) error {
			book.ID = 124
			return nil
		})

	client.EXPECT().
		IsCategoryExists(gomock.Any(), int64(25)).
		Return(true, nil)

	book, err := service.CreateBook(
		context.Background(),
		"new book",
		25,
	)

	require.Nil(t, err)
	require.Equal(t, int64(25), book.CategoryID)
}

func TestCreateBook_Success_ReturnID(t *testing.T) {
	service, repo, client := setup(t)

	repo.EXPECT().
		SaveBook(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, book *Book) error {
			book.ID = 124
			return nil
		})

	client.EXPECT().
		IsCategoryExists(gomock.Any(), int64(3)).
		Return(true, nil)

	book, err := service.CreateBook(
		context.Background(),
		"another book",
		3,
	)

	require.Nil(t, err)
	require.Equal(t, int64(124), book.ID)
}

func TestCreateBook_CategoryDoesNotExist(t *testing.T) {
	service, _, client := setup(t)

	client.EXPECT().
		IsCategoryExists(gomock.Any(), int64(3)).
		Return(false, nil)

	_, err := service.CreateBook(
		context.Background(),
		"another book",
		3,
	)

	require.NotNil(t, err)
}
