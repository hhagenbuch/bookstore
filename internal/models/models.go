package models

import "gorm.io/gorm"

// User represents a user in the bookstore application.
type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

// Book represents a book in the bookstore application.
type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Order represents an order in the bookstore application.
type Order struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	Books  []Book `json:"books" gorm:"many2many:order_books;"`
}
