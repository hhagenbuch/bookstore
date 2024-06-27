package services

import (
	"github.com/hhagenbuch/bookstore/internal/database"
	"github.com/hhagenbuch/bookstore/internal/errors"
	"github.com/hhagenbuch/bookstore/internal/models"
)

// CreateUser creates a new user in the database.
func CreateUser(user *models.User) error {
	if err := database.DB.Create(user).Error; err != nil {
		return errors.NewInternalServerError(err)
	}
	return nil
}

// GetBooks retrieves all books from the database.
func GetBooks() ([]models.Book, error) {
	var books []models.Book
	result := database.DB.Find(&books)
	if result.Error != nil {
		return nil, errors.NewInternalServerError(result.Error)
	}
	return books, nil
}

// CreateOrder creates a new order in the database.
// It preloads the books to ensure they exist before creating the order.
func CreateOrder(order *models.Order) error {
	var books []models.Book
	for _, book := range order.Books {
		var b models.Book
		if err := database.DB.First(&b, book.ID).Error; err != nil {
			return errors.NewNotFoundError("Book not found")
		}
		books = append(books, b)
	}
	order.Books = books

	if err := database.DB.Create(order).Error; err != nil {
		return errors.NewInternalServerError(err)
	}
	return nil
}

// GetOrders retrieves all orders for a given user ID from the database.
func GetOrders(userID uint) ([]models.Order, error) {
	var orders []models.Order
	result := database.DB.Preload("Books").Where("user_id = ?", userID).Find(&orders)
	if result.Error != nil {
		return nil, errors.NewInternalServerError(result.Error)
	}
	return orders, nil
}
