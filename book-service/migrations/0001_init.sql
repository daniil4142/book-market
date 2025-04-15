-- +goose Up
CREATE TABLE books (
                          id SERIAL PRIMARY KEY,
                          name TEXT NOT NULL,
                          category_id BIGINT NOT NULL
);

-- +goose Down
DROP TABLE books;