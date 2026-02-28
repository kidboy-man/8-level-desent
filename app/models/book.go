package models

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Year   int    `json:"year" binding:"required,gt=0"`
}

type CreateBookRequest struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Year   int    `json:"year" binding:"required,gt=0"`
}

type UpdateBookRequest struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Year   int    `json:"year" binding:"required,gt=0"`
}
