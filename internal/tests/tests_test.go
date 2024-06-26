package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"github.com/hhagenbuch/bookstore/internal/database"
	"github.com/hhagenbuch/bookstore/internal/handlers"
	"github.com/hhagenbuch/bookstore/internal/middleware"
	"github.com/stretchr/testify/assert"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func setupRouter() *mux.Router {
	database.InitDB()
	database.SeedDB()
	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.RecoverMiddleware)
	handlers.RegisterHandlers(r)
	return r
}

func TestCreateUser(t *testing.T) {
	r := setupRouter()

	userJSON := `{"email": "test2@example.com", "password": "password"}`
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer([]byte(userJSON)))
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected response code to be 200")

	var response Response
	err := json.NewDecoder(rr.Body).Decode(&response)
	assert.Nil(t, err)
	assert.Equal(t, "test2@example.com", response.Data.(map[string]interface{})["email"])
}

func TestGetBooks(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/books", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected response code to be 200")

	var response Response
	err := json.NewDecoder(rr.Body).Decode(&response)
	assert.Nil(t, err)
	assert.NotEmpty(t, response.Data)
}

func TestCreateOrder(t *testing.T) {
	r := setupRouter()

	// Create a new user and get the user ID
	userJSON := `{"email": "order_user@example.com", "password": "password"}`
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer([]byte(userJSON)))
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var userResponse Response
	err := json.NewDecoder(rr.Body).Decode(&userResponse)
	assert.Nil(t, err)
	userID := userResponse.Data.(map[string]interface{})["ID"].(float64)

	orderJSON := fmt.Sprintf(`{"user_id": %d, "books": [{"ID": %d}, {"ID": %d}]}`, int(userID), database.Book1ID, database.Book2ID)
	req, _ = http.NewRequest("POST", "/orders", bytes.NewBuffer([]byte(orderJSON)))
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected response code to be 200")

	var response Response
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.Nil(t, err)
	log.Printf("Response Data: %+v", response.Data)
	assert.NotNil(t, response.Data.(map[string]interface{})["books"], "Expected books to be not nil")
	assert.Len(t, response.Data.(map[string]interface{})["books"], 2)
}

func TestGetOrders(t *testing.T) {
	r := setupRouter()

	// Create a new user and get the user ID
	userJSON := `{"email": "get_orders_user@example.com", "password": "password"}`
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer([]byte(userJSON)))
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	var userResponse Response
	err := json.NewDecoder(rr.Body).Decode(&userResponse)
	assert.Nil(t, err)
	userID := userResponse.Data.(map[string]interface{})["ID"].(float64)

	// Create an order for the user
	orderJSON := fmt.Sprintf(`{"user_id": %d, "books": [{"ID": %d}, {"ID": %d}]}`, int(userID), database.Book1ID, database.Book2ID)
	req, _ = http.NewRequest("POST", "/orders", bytes.NewBuffer([]byte(orderJSON)))
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Get orders for the user
	req, _ = http.NewRequest("GET", "/orders/"+strconv.Itoa(int(userID)), nil)
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected response code to be 200")

	var response Response
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.Nil(t, err)
	log.Printf("Orders Response Data: %+v", response.Data)
	assert.NotEmpty(t, response.Data)
}

func TestCreateUser_InvalidPayload(t *testing.T) {
	r := setupRouter()

	invalidJSON := `{"email": "test@example.com"}`
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer([]byte(invalidJSON)))
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code, "Expected response code to be 400")
}

func TestGetOrders_InvalidUserID(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/orders/invalid", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code, "Expected response code to be 400")
}
