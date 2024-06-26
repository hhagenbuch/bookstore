package database

import (
	"fmt"

	"github.com/hhagenbuch/bookstore/internal/models"
)

var (
	Book1ID uint
	Book2ID uint
	UserID  uint
)

// SeedDB seeds the database with initial data.
func SeedDB() error {
	// Clear existing data
	DB.Exec("DELETE FROM users")
	DB.Exec("DELETE FROM books")
	DB.Exec("DELETE FROM orders")
	DB.Exec("DELETE FROM order_books")

	books := []models.Book{
		{Title: "Harry Potter and the Sorcerer's Stone", Author: "J.K. Rowling"},
		{Title: "The Lord of the Rings", Author: "J.R.R. Tolkien"},
		{Title: "Dune", Author: "Frank Herbert"},
		{Title: "Foundation", Author: "Isaac Asimov"},
		{Title: "Snow Crash", Author: "Neal Stephenson"},
		{Title: "The Left Hand of Darkness", Author: "Ursula K. Le Guin"},
		{Title: "Ender's Game", Author: "Orson Scott Card"},
		{Title: "Hyperion", Author: "Dan Simmons"},
		{Title: "The Martian", Author: "Andy Weir"},
	}

	for i, book := range books {
		if err := DB.Create(&book).Error; err != nil {
			return fmt.Errorf("error seeding book: %w", err)
		}
		if i == 0 {
			Book1ID = book.ID
		} else if i == 1 {
			Book2ID = book.ID
		}
	}

	user := models.User{Email: "test@example.com", Password: "password"}
	if err := DB.Create(&user).Error; err != nil {
		return fmt.Errorf("error seeding user: %w", err)
	}
	UserID = user.ID

	return nil
}
