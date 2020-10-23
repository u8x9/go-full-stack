package entry

type Book struct {
	Name string `json:"name" form:"name" binding:"required"`
	Price int `json:"price" form:"price" binding:"required"`
}
