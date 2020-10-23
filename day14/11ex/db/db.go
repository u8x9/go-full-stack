package db

import (
	"database/sql"
	"github.com/u8x9/go-full-stack/day14/11ex/entry"
	"github.com/u8x9/go-full-stack/day14/11ex/entry/form"
	"github.com/u8x9/go-full-stack/day14/11ex/entry/model"
)
import _ "github.com/go-sql-driver/mysql"

var db *sql.DB

func Init() error {
	dsn := "root:root@tcp(127.0.0.1:3306)/gofullstack"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return nil
}

func AddBook(form *form.BookForm) (*model.BookModel, error) {
	sqlStr := "INSERT INTO book(name, price) VALUES (?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	result, err := stmt.Exec(form.Name, form.Price)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &model.BookModel{
		ID: id,
		Book: entry.Book{
			Name:  form.Name,
			Price: form.Price,
		},
	}, nil
}
func AllBook() ([]*model.BookModel, error) {
	sqlStr := "SELECT id,name,price FROM book ORDER BY id DESC LIMIT 30"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	books := make([]*model.BookModel, 0, 30)
	for rows.Next() {
		var b model.BookModel
		if err := rows.Scan(&b.ID, &b.Name, &b.Price); err != nil {
			return nil, err
		}
		books = append(books, &b)
	}
	return books, nil
}
func DelBook(id int64) (int64, error) {
	sqlstr := "DELETE FROM book WHERE id=?"
	stmt, err := db.Prepare(sqlstr)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
