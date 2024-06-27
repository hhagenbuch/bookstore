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
- [Docker](https://www.docker.com/products/docker-desktop) (optional)

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

## Running the Application with Docker

1. **Build the Docker image:**

    ```sh
    docker build -t bookstore .
    ```

2. **Run the Docker container:**

    ```sh
    docker run -p 8000:8000 bookstore
    ```

   This command maps the `bookstore.db` file from your current directory to the `/app/bookstore.db` path inside the Docker container, ensuring data persistence.

## Usage
3. **Testing the API:**


- curl -X GET http://localhost:8000/books
- curl -X POST http://localhost:8000/users -d '{"email": "example@test.com", "password": "password"}'
- curl -X POST http://localhost:8000/orders -d '{"user_id": 1, "books": [{"ID": 1}, {"ID": 2}]}'
- curl -X GET http://localhost:8000/orders/1



## Future Improvements

### Authentication and Authorization:

- **Implementation:** Use JWT (JSON Web Tokens) to authenticate users. Implement role-based access control (RBAC) to restrict access to specific endpoints based on user roles.
- **Steps:**
    1. Install the JWT library (e.g., `github.com/dgrijalva/jwt-go`).
    2. Create middleware to verify the JWT token.
    3. Modify your handlers to check user roles and permissions.

### Pagination and Filtering:

- **Implementation:** Allow users to paginate through books and orders, and apply filters to narrow down results based on certain criteria (e.g., author, title).
- **Steps:**
    1. Add query parameters (e.g., `page`, `limit`, `author`, `title`) to your endpoints.
    2. Modify your database queries to handle pagination and filtering.
    3. Update the response format to include pagination metadata (e.g., total items, total pages).

### Validation:

- **Implementation:** Add input validation for all API endpoints to ensure that the data being processed meets certain criteria.
- **Steps:**
    1. Install a validation library (e.g., `github.com/go-playground/validator`).
    2. Create validation rules for your models.
    3. Validate incoming requests in your handlers before processing them.


## Contact

For any inquiries, please contact [heyward360@gmail.com](mailto:heyward360@gmail.com)