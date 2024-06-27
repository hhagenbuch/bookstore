# Bookstore API

## Overview

The Bookstore API is a robust and scalable RESTful API designed for managing users, books, and orders in an online bookstore. It is built using Go, leveraging Gorilla Mux for routing and GORM for database management.

## Features

- Create a user account
- Retrieve a list of books
- Create an order with multiple books
- Retrieve the order history for a user

## Technologies

- **Programming Language:** Go
- **Routing:** Gorilla Mux
- **ORM:** GORM
- **Database:** SQLite

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
            },
            {
                "ID": 2,
                "title": "The Lord of the Rings",
                "author": "J.R.R. Tolkien"
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
                },
                {
                    "ID": 2,
                    "title": "The Lord of the Rings",
                    "author": "J.R.R. Tolkien"
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
                    },
                    {
                        "ID": 2,
                        "title": "The Lord of the Rings",
                        "author": "J.R.R. Tolkien"
                    }
                ],
                "createdAt": "...",
                "updatedAt": "..."
            }
        ]
    }
    ```

## Running Tests

To run the automated tests, use the following command:

```sh
go test ./internal/tests
```

## Future Improvements 

- **Authentication and Authorization:** Adding JWT-based authentication and role-based access control.
- **Pagination and Filtering:** Implementing pagination and filtering for the list of books and orders.
- **Dockerization:** Containerizing the application using Docker for easier deployment and scalability.


## Contact

For any inquiries, please contact [heyward360@gmail.com](mailto:heyward360@gmail.com)