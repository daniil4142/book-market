package book_service

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	DB *sqlx.DB
}

func newRepo(db *sqlx.DB) IRepository {
	return repository{
		DB: db,
	}
}

func (r repository) SaveBook(ctx context.Context, book *Book) (int64, error) {

	query := sq.Insert("books").PlaceholderFormat(sq.Dollar).Columns(
		"name", "category_id").Values(book.Name, book.CategoryID).Suffix("RETURNING id").RunWith(r.DB)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return 0, err
	}

	var id int64
	if rows.Next() {
		err = rows.Scan(&id)

		if err != nil {
			return 0, err
		}

		return id, nil
	} else {
		return 0, sql.ErrNoRows
	}
}

func (r repository) DeleteBooks(ctx context.Context, ids []int64) error {

	query := sq.Delete("books").PlaceholderFormat(sq.Dollar).Where(sq.Eq{"id": ids})

	s, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = r.DB.ExecContext(ctx, s, args...)
	return err
}

func (r repository) GetBooks(ctx context.Context, ids []int64) ([]Book, error) {

	query := sq.Select("*").PlaceholderFormat(sq.Dollar).From("books").Where(sq.Eq{"id": ids})

	s, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	var res []Book
	err = r.DB.SelectContext(ctx, &res, s, args...)

	return res, err
}
