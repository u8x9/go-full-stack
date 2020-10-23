package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/u8x9/go-full-stack/day14/11ex/db"
	"github.com/u8x9/go-full-stack/day14/11ex/entry/form"
	"net/http"
	"strconv"
)

func Index(c *gin.Context) {
	books, err := db.AllBook()
	if err != nil {
		c.HTML(http.StatusBadRequest, "get book list failed: %v", err)
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{"books": books})
}

func Add(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		c.HTML(http.StatusOK, "add.html", nil)
		return
	}
	var book form.BookForm
	if err := c.Bind(&book); err != nil {
		c.HTML(http.StatusBadRequest, "invalid input: %v", err)
		return
	}
	if _, err := db.AddBook(&book); err != nil {
		c.HTML(http.StatusBadRequest, "add book failed: %v", err)
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}

func Del(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		c.HTML(http.StatusBadRequest, "invalid input: %v", err)
		return
	}
	if _, err := db.DelBook(id); err != nil {
		c.HTML(http.StatusBadRequest, "delete book failed: %v", err)
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/")
}
