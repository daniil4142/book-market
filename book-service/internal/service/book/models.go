package book_service

type Book struct {
	ID         int64  `db:"id"`
	Name       string `db:"name"`
	CategoryID int64  ` db:"category_id"`
}
