package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hhagenbuch/bookstore/internal/models"
	"github.com/hhagenbuch/bookstore/internal/services"
)

// Response represents the structure of API responses.
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// RegisterHandlers registers the API endpoints with the router.
func RegisterHandlers(r *mux.Router) {
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/orders", createOrder).Methods("POST")
	r.HandleFunc("/orders/{userID}", getOrders).Methods("GET")
}

// createUser handles the creation of a new user.
func createUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if user.Email == "" || user.Password == "" {
		sendErrorResponse(w, http.StatusBadRequest, "Email and password are required")
		return
	}
	if err := services.CreateUser(&user); err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendSuccessResponse(w, "User created successfully", user)
}

// getBooks handles retrieving the list of books.
func getBooks(w http.ResponseWriter, r *http.Request) {
	books, err := services.GetBooks()
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendSuccessResponse(w, "Books retrieved successfully", books)
}

// createOrder handles the creation of a new order.
func createOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := services.CreateOrder(&order); err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccessResponse(w, "Order created successfully", order)
}

// getOrders handles retrieving orders for a specific user.
func getOrders(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["userID"])
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	orders, err := services.GetOrders(uint(userID))
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendSuccessResponse(w, "Orders retrieved successfully", orders)
}

// sendErrorResponse sends a JSON-encoded error response.
func sendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	response := Response{Status: "error", Message: message}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

// sendSuccessResponse sends a JSON-encoded success response.
func sendSuccessResponse(w http.ResponseWriter, message string, data interface{}) {
	response := Response{Status: "success", Message: message, Data: data}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}
