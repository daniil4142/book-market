package cat_repository

import (
	"context"

	"github.com/daniil4142/book-market/category-service/internal/pkg/db"
	"github.com/daniil4142/book-market/category-service/internal/service/category"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return Repository{db: db}
}

func (r Repository) FindAllCategories(ctx context.Context) (category.Categories, error) {
	sb := db.StatementBuilder.
		Select("id", "name").
		From("category")

	query, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}

	var cats category.Categories
	err = r.db.SelectContext(ctx, &cats, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "db.SelectContext()")
	}

	return cats, nil
}
