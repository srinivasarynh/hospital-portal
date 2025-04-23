# Hospital Portal API Documentation

A hospital management system API that handles user authentication (doctors, receptionists), patient management, and more.

---

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

---

## Project Structure and Other Details

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

---

Setup Instructions

Pre-requisites:
- Docker installed
- Go installed
- PostgreSQL running (through Docker)

Steps to Run:

1. Clone the Repository
git clone https://github.com/srinivasarynh/hospital-portal.git
cd hospital-portal


2. Create `.env` file in the root directory with the following content:
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=your_password
DB_NAME=hospital_db
JWT_SECRET=your_jwt_secret

3. Start Docker containers (if using Docker for PostgreSQL and app):
docker-compose up --build

4. Run the application locally (if not using Docker):
go run main.go

5. Seeding Data (if using Docker for PostgreSQL and app):
docker exec -it hospital_backend go run cmd/cli/main.go --seed

6. Seeding Data (if not using Docker):
go run cmd/cli/main.go --seed

---

API Endpoints

🔐 Authentication

`POST /auth/login`
Authenticate a user and return a JWT token.

Request:
{
  "username": "receptionist1", // username for doctor "doctor1"
  "password": "receptionistpassword" // password for doctor "doctorpassword"
}

Response:
{
  "token": "your_jwt_token"
}


👩‍⚕️ Receptionist Endpoints

`POST /reception/patients`
Creates a new patient record.

Request:
{
  "full_name": "John Doe",
  "age": 30,
  "gender": "Male",
  "symptoms": "Fever",
  "notes": "Needs immediate attention"
}

Response:
Status 201 - Created

`GET /reception/patients`
Retrieves a list of all patients.

Response:
[
  {
    "id": 1,
    "full_name": "John Doe",
    "age": 45,
    "gender": "Male",
    "symptoms": "Fever, Cough",
    "notes": "Patient needs urgent attention"
  },
  ...
]

`GET /reception/patients/:id`
Retrieves a specific patient record by ID.

Response:
{
  "id": 1,
  "full_name": "John Doe",
  "age": 45,
  "gender": "Male",
  "symptoms": "Fever, Cough",
  "notes": "Patient needs urgent attention"
}


`PUT /reception/patients/:id`
Description: Updates patient details by ID.

Request:
{
  "full_name": "John Doe",
  "age": 46,
  "gender": "Male",
  "symptoms": "Fever, Cough, Headache",
  "notes": "Patient requires additional tests"
}


`DELETE /reception/patients/:id`
Description: Deletes a patient record by ID.

Response:
{
  "message": "Patient record deleted successfully"
}


👩‍⚕️ Doctor Endpoints

`GET /doctor/patients`
Description: Retrieves a list of all patients.

Response:
[
  {
    "id": 1,
    "full_name": "John Doe",
    "age": 45,
    "gender": "Male",
    "symptoms": "Fever, Cough",
    "notes": "Patient needs urgent attention"
  },
  ...
]

`GET /doctor/patients/:id`
Description: Retrieves a specific patient record by ID.

Response:
{
  "id": 1,
  "full_name": "John Doe",
  "age": 45,
  "gender": "Male",
  "symptoms": "Fever, Cough",
  "notes": "Patient needs urgent attention"
}

`PUT /doctor/patients/:id`
Description: Updates patient details by ID.

Request:
{
  "full_name": "John Doe",
  "age": 46,
  "gender": "Male",
  "symptoms": "Fever, Cough, Headache",
  "notes": "Patient requires additional tests"
}


`PATCH /doctor/patients/:id/notes`
Description: Updates a patient's notes.

Request:
{
  "notes": "Patient's condition is worsening"
}

---

Database Design

User Model:
- `Username` (unique)
- `Password` (hashed)
- `Role` (doctor/receptionist)

Patient Model:
- `Full Name`
- `Age`
- `Gender`
- `Symptoms`
- `Notes`
- `Registered At` (auto-generated)

Relationships:
- One-to-many relationship between users and patients (e.g., doctors managing patients).

---

Environment Configuration

.env Configuration:
- DB_HOST: Host of your PostgreSQL database.
- DB_PORT: Port for PostgreSQL.
- DB_USER: Database username.
- DB_PASS: Database password.
- DB_NAME: Database name.
- JWT_SECRET: Secret for JWT token signing.

---

Middlewares

Auth Middleware
The Auth Middleware is used to ensure that the user is authenticated and authorized to perform certain actions. It checks the JWT token passed in the request header and ensures the role is correct (e.g., "doctor" or "receptionist").

---

Conclusion
This Hospital Portal API provides essential functionalities for managing patients, doctor-receptionist roles, and authentication. The API is built with Go and PostgreSQL and is secured with JWT for authentication. You can expand the system further by adding features like patient appointments, notifications, and more.
