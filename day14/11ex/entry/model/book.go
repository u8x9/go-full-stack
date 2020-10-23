package model

import "github.com/u8x9/go-full-stack/day14/11ex/entry"

type BookModel struct {
	ID int64 `json:"id" form:"id"`
	entry.Book
}
