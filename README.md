# Bookstore API

## Overview

The Bookstore API is a simple RESTful API for managing users, books, and orders in an online bookstore. It is built using Go and utilizes Gorilla Mux for routing and GORM for database management.

## Features

- Create a user account
- Retrieve a list of books
- Create an order with multiple books
- Retrieve the order history for a user

## Technologies

- Go
- Gorilla Mux
- GORM
- SQLite

## Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16 or later)
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

## Setup

1. **Clone the repository:**

    ```sh
    git clone https://github.com/hhagenbuch/bookstore.git
    cd bookstore
    ```

2. **Install dependencies:**

    ```sh
    go mod tidy
    ```

3. **Run the application:**

    ```sh
    go run cmd/main.go
    ```

## API Endpoints

### Create User

- **URL:** `/users`
- **Method:** `POST`
- **Payload:**
    ```json
    {
        "email": "test@example.com",
        "password": "password"
    }
    ```
- **Response:**
    ```json
    {
        "status": "success",
        "message": "User created successfully",
        "data": {
            "ID": 1,
            "email": "test@example.com",
            "password": "password",
            "createdAt": "...",
            "updatedAt": "..."
        }
    }
    ```

### Get Books

- **URL:** `/books`
- **Method:** `GET`
- **Response:**
    ```json
    {
        "status": "success",
        "message": "Books retrieved successfully",
        "data": [
            {
                "ID": 1,
                "title": "Harry Potter and the Sorcerer's Stone",
                "author": "J.K. Rowling"
            }
        ]
    }
    ```

### Create Order

- **URL:** `/orders`
- **Method:** `POST`
- **Payload:**
    ```json
    {
        "user_id": 1,
        "books": [
            {"ID": 1},
            {"ID": 2}
        ]
    }
    ```
- **Response:**
    ```json
    {
        "status": "success",
        "message": "Order created successfully",
        "data": {
            "ID": 1,
            "user_id": 1,
            "books": [
                {
                    "ID": 1,
                    "title": "Harry Potter and the Sorcerer's Stone",
                    "author": "J.K. Rowling"
                }
            ],
            "createdAt": "...",
            "updatedAt": "..."
        }
    }
    ```

### Get Orders

- **URL:** `/orders/{userID}`
- **Method:** `GET`
- **Response:**
    ```json
    {
        "status": "success",
        "message": "Orders retrieved successfully",
        "data": [
            {
                "ID": 1,
                "user_id": 1,
                "books": [
                    {
                        "ID": 1,
                        "title": "Harry Potter and the Sorcerer's Stone",
                        "author": "J.K. Rowling"
                    }
                ],
                "createdAt": "...",
                "updatedAt": "..."
            }        ]
    }
    ```

## Running Tests

To run the automated tests, use the following command:

```sh
go test ./internal/tests