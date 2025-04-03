package book_service

import (
	"context"

	book_service "github.com/daniil4142/book-market/book-service/internal/service/book"
	desc "github.com/daniil4142/book-market/book-service/pkg/book-service"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateProduct(ctx context.Context, req *desc.CreateProductRequest) (*desc.CreateProductResponse, error) {
	res, err := i.bookService.CreateProduct(ctx, req.GetName(), req.GetCategoryId())
	if err != nil {
		if err == book_service.ErrWrongCategory {
			details := &errdetails.BadRequest{
				FieldViolations: []*errdetails.BadRequest_FieldViolation{
					{
						Field:       "categoryId",
						Description: "wrong category",
					},
				},
			}

			st := status.New(codes.InvalidArgument, "wrong category")

			withDetails, stErr := st.WithDetails(details)
			if stErr != nil {
				return nil, st.Err()
			}

			return nil, withDetails.Err()
		}

		return nil, err
	}

	return &desc.CreateProductResponse{
		Result: convertProductToPb(res),
	}, nil
}

func convertProductToPb(res *book_service.Book) *desc.Book {
	return &desc.Book{
		Id:         res.ID,
		Name:       res.Name,
		CategoryId: res.CategoryID,
	}
}
