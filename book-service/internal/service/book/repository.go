package book_service

import (
	"context"
	"database/sql"

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
		"name", "category_id", "info").Values(book.Name, book.CategoryID).Suffix("RETURNING id").RunWith(r.DB)

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
