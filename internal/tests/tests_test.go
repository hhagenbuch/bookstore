package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/hhagenbuch/bookstore/internal/database"
	"github.com/hhagenbuch/bookstore/internal/handlers"
	"github.com/hhagenbuch/bookstore/internal/models"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *mux.Router {
	database.InitDB()
	database.SeedDB()
	r := mux.NewRouter()
	handlers.RegisterHandlers(r)
	return r
}

func TestCreateUser(t *testing.T) {
	r := setupRouter()

	userJSON := `{"email": "test@example.com", "password": "password"}`
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer([]byte(userJSON)))
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected response code to be 200")

	var user models.User
	err := json.NewDecoder(rr.Body).Decode(&user)
	assert.Nil(t, err)
	assert.Equal(t, "test@example.com", user.Email)
}

func TestGetBooks(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/books", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected response code to be 200")

	var books []models.Book
	err := json.NewDecoder(rr.Body).Decode(&books)
	assert.Nil(t, err)
	assert.NotEmpty(t, books)
}

func TestCreateOrder(t *testing.T) {
	r := setupRouter()

	orderJSON := `{"user_id": 1, "books": [{"ID": 1}, {"ID": 2}]}`
	req, _ := http.NewRequest("POST", "/orders", bytes.NewBuffer([]byte(orderJSON)))
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected response code to be 200")

	var order models.Order
	err := json.NewDecoder(rr.Body).Decode(&order)
	assert.Nil(t, err)
	assert.Equal(t, uint(1), order.UserID)
	assert.Len(t, order.Books, 2)
}

func TestGetOrders(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/orders/1", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected response code to be 200")

	var orders []models.Order
	err := json.NewDecoder(rr.Body).Decode(&orders)
	assert.Nil(t, err)
	assert.NotEmpty(t, orders)
}
