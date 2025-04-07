package book_service

import (
	"context"

	book_service "github.com/daniil4142/book-market/book-service/internal/service/book"
	desc "github.com/daniil4142/book-market/book-service/pkg/book-service"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateBook(ctx context.Context, req *desc.CreateBookRequest) (*desc.CreateBookResponse, error) {
	log.Info().
		Msgf("start createbook")
	res, err := i.bookService.CreateBook(ctx, req.GetName(), req.GetCategoryId())
	if err != nil {
		if err == book_service.ErrWrongCategory {
			// details := &errdetails.BadRequest{
			// 	FieldViolations: []*errdetails.BadRequest_FieldViolation{
			// 		{
			// 			Field:       "categoryId",
			// 			Description: "wrong category",
			// 		},
			// 	},
			// }

			// st := status.New(codes.InvalidArgument, "wrong category")

			// withDetails, stErr := st.WithDetails(details)
			// if stErr != nil {
			// 	return nil, st.Err()
			// }

			return nil, status.Error(codes.InvalidArgument, "wrong category")
		}

		return nil, err
	}

	return &desc.CreateBookResponse{
		Result: convertBookToPb(res),
	}, nil
}

func convertBookToPb(res *book_service.Book) *desc.Book {
	return &desc.Book{
		Id:         res.ID,
		Name:       res.Name,
		CategoryId: res.CategoryID,
	}
}
