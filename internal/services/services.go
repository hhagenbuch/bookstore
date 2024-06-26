package services

import (
	"log"

	"github.com/hhagenbuch/bookstore/internal/database"
	"github.com/hhagenbuch/bookstore/internal/models"
)

// CreateUser creates a new user in the database.
func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

// GetBooks retrieves all books from the database.
func GetBooks() ([]models.Book, error) {
	var books []models.Book
	result := database.DB.Find(&books)
	return books, result.Error
}

// CreateOrder creates a new order in the database.
// It preloads the books to ensure they exist before creating the order.
func CreateOrder(order *models.Order) error {
	var books []models.Book
	for _, book := range order.Books {
		var b models.Book
		if err := database.DB.First(&b, book.ID).Error; err != nil {
			return err
		}
		books = append(books, b)
	}
	order.Books = books

	err := database.DB.Create(order).Error
	if err != nil {
		log.Printf("Error creating order: %v", err)
	} else {
		log.Printf("Created Order: %+v", order)
	}
	return err
}

// GetOrders retrieves all orders for a given user ID from the database.
func GetOrders(userID uint) ([]models.Order, error) {
	var orders []models.Order
	result := database.DB.Preload("Books").Where("user_id = ?", userID).Find(&orders)
	return orders, result.Error
}
