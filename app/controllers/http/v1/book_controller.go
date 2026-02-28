package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kidboy-man/8-level-desent/app/models"
	"github.com/kidboy-man/8-level-desent/app/repositories"
	"github.com/kidboy-man/8-level-desent/app/services"
)

type BookController struct {
	bookService *services.BookService
}

func NewBookController(bookService *services.BookService) *BookController {
	return &BookController{bookService: bookService}
}

func (ctrl *BookController) Create(c *gin.Context) {
	var req models.CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	book, err := ctrl.bookService.CreateBook(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func (ctrl *BookController) GetAll(c *gin.Context) {
	filter := repositories.BookFilter{
		Author: c.Query("author"),
	}

	if pageStr := c.Query("page"); pageStr != "" {
		if page, err := strconv.Atoi(pageStr); err == nil {
			filter.Page = page
		}
	}

	if limitStr := c.Query("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil {
			filter.Limit = limit
		}
	}

	books, total, err := ctrl.bookService.GetAllBooks(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	response := gin.H{
		"data":  books,
		"total": total,
	}

	if filter.Page > 0 && filter.Limit > 0 {
		response["page"] = filter.Page
		response["limit"] = filter.Limit
	}

	c.JSON(http.StatusOK, response)
}

func (ctrl *BookController) GetByID(c *gin.Context) {
	id := c.Param("id")

	book, err := ctrl.bookService.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (ctrl *BookController) Update(c *gin.Context) {
	id := c.Param("id")

	var req models.UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	book, err := ctrl.bookService.UpdateBook(id, &req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (ctrl *BookController) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := ctrl.bookService.DeleteBook(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
