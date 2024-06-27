package handlers

import (
	"encoding/json"
	"github.com/hhagenbuch/bookstore/internal/errors"
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
		sendErrorResponse(w, errors.NewBadRequestError("Invalid request payload"))
		return
	}
	if user.Email == "" || user.Password == "" {
		sendErrorResponse(w, errors.NewBadRequestError("Email and password are required"))
		return
	}
	if err := services.CreateUser(&user); err != nil {
		sendErrorResponse(w, err.(*errors.AppError))
		return
	}
	sendSuccessResponse(w, "User created successfully", user)
}

// getBooks handles retrieving the list of books.
func getBooks(w http.ResponseWriter, r *http.Request) {
	books, err := services.GetBooks()
	if err != nil {
		sendErrorResponse(w, err.(*errors.AppError))
		return
	}
	sendSuccessResponse(w, "Books retrieved successfully", books)
}

// createOrder handles the creation of a new order.
func createOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		sendErrorResponse(w, errors.NewBadRequestError("Invalid request payload"))
		return
	}
	if err := services.CreateOrder(&order); err != nil {
		sendErrorResponse(w, err.(*errors.AppError))
		return
	}
	sendSuccessResponse(w, "Order created successfully", order)
}

// getOrders handles retrieving orders for a specific user.
func getOrders(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["userID"])
	if err != nil {
		sendErrorResponse(w, errors.NewBadRequestError("Invalid user ID"))
		return
	}
	orders, err := services.GetOrders(uint(userID))
	if err != nil {
		sendErrorResponse(w, err.(*errors.AppError))
		return
	}
	sendSuccessResponse(w, "Orders retrieved successfully", orders)
}

// sendErrorResponse sends an error response to the client.
func sendErrorResponse(w http.ResponseWriter, appErr *errors.AppError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(appErr.Code)
	err := json.NewEncoder(w).Encode(Response{Status: "error", Message: appErr.Message})
	if err != nil {
		return
	}
}

// sendSuccessResponse sends a success response to the client.
func sendSuccessResponse(w http.ResponseWriter, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(Response{Status: "success", Message: message, Data: data})
	if err != nil {
		return
	}
}
