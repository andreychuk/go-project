# Golang REST API with GORM, PostgreSQL, and Clean Architecture

This project is a simple REST API built with Golang using the Gin framework, GORM for ORM, and PostgreSQL as the database. It follows Clean Architecture principles to ensure a maintainable and scalable codebase.

## Features

- CRUD operations for Users, Groups, and Roles.
- Many-to-many relationship between Users and Groups.
- One-to-many relationship between Users and Roles.
- Automatic database migrations using GORM.
- Clean Architecture structure for clear separation of concerns.

## Project Structure

```bash
├── cmd/
│ └── main.go # Entry point of the application
├── internal/
│ ├── domain/
│ │ └── models/ # Domain models (User, Group, Role, UserGroup)
│ ├── infrastructure/
│ │ ├── db/ # Database connection setup
│ │ └── repository/ # Repositories for accessing the database
│ ├── interfaces/
│ │ └── http/ # HTTP handlers (controllers)
│ └── usecase/ # Business logic (use cases)
├── db/
├── go.mod # Go module file
├── go.sum # Go dependencies
└── README.md # Project README
```

## Prerequisites

- Go 1.22 or later
- PostgreSQL
- [Gin](https://github.com/gin-gonic/gin) for the HTTP framework
- [GORM](https://gorm.io/) for ORM

## Run the Application

Start the application with:

```bash
go run cmd/main.go
```

The server should start on http://localhost:8080.

## Postman Environment Setup

To test the API using Postman, you need to set up an environment with the appropriate base URL. Follow these steps to create and configure a new environment in Postman:

### 1. Create the Environment

1. **Open Postman**:

   - Launch the Postman application on your computer.

2. **Navigate to the Environments Tab**:

   - Click on the "Environments" tab in the top-right corner of the Postman window.

3. **Create a New Environment**:

   - Click on the "Add" button to create a new environment.

4. **Configure the Environment**:
   - **Name the Environment**:
     - Enter a name for the environment, such as `Local`.
   - **Add a New Variable**:
     - Click on "Add" to create a new variable.
     - **Variable**: `base_url`
     - **Initial Value**: `http://localhost:8080`
     - **Current Value**: `http://localhost:8080`
   - **Save the Environment**:
     - Click the "Save" button to save your new environment.
