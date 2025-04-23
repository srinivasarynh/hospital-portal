# Hospital Portal API Documentation

## Project Overview

**Project Name**: Hospital Portal API  
**Description**: A hospital management system API built with Go. It provides an interface for doctors and receptionists to manage patient data, authenticate users, and interact with the hospital database securely.  
**Tech Stack**:
- **Go (Golang)**: Backend API development
- **PostgreSQL**: Database for storing user and patient information
- **Gin Framework**: Web framework for Go to build APIs
- **JWT (JSON Web Tokens)**: Authentication mechanism for user login and access control
- **GORM**: ORM (Object-Relational Mapping) library for Go to interact with PostgreSQL
- **Docker**: Containerization for development, testing, and deployment

## Project Structure

```plaintext
hospital-portal/
│
├── cmd/                # CLI commands for the application (e.g., seeding users)
│   └── cli/            # Command-line interface logic for seeding and other tasks
│       └── main.go     # Main entry point for CLI commands
│
├── controllers/        # Contains all the handlers for different API endpoints
│   ├── auth_controller.go
│   ├── patient_controller.go
│   ├── doctor_controller.go
│   └── reception_controller.go
│
├── middlewares/        # Middlewares for API (authentication, etc.)
│   └── auth_middleware.go
│
├── models/             # GORM models for interacting with the database
│   ├── user.go
│   └── patient.go
│
├── routes/             # Define the routes and group them
│   └── routes.go
│
├── utils/              # Helper functions for common tasks (e.g., password hashing)
│   ├── jwt.go
│   └── response.go
│
├── main.go             # Main entry point to Application
├── Dockerfile          # Docker configuration file for building the app image
├── docker-compose.yml  # Docker Compose configuration for multi-container setup
└── go.mod              # Go module file with dependencies
# hospital-portal
