# Go Smartcerti – Backend API

**Go Smartcerti** is a backend API service for the Smartcerti platform, designed to manage certifications, trainings, users, and vendors. Built with [Golang](https://golang.org/) using the [Fiber](https://gofiber.io/) framework and [MySQL](https://www.mysql.com/) for the database.

---

## 🚀 Features

- 🔐 JWT-based Authentication
- 👤 User Management (CRUD)
- 🏢 Vendor Management
- 📚 Subject & Area of Interest Management
- 🏆 Training & Certification Management
- ⚙️ Middleware-protected routes
- ✅ RESTful API structure

---

## 📦 Tech Stack

- Go (Golang)
- Fiber Web Framework
- MySQL
- JWT for Authentication

---

## 🛠️ Setup Instructions

### Prerequisites

Ensure you have the following installed:

- Go 1.20+
- MySQL Server
- Git

### Installation

1. **Clone the repository**

```bash
git clone https://github.com/wisamahmad123/go-smartcerti.git
cd go-smartcerti
```


2. **Create a .env file**

SECRET_KEY=YOUR_SECRET_KEY
PORT=:8080
DB=DBUSERNAME:DBPASSWORD@tcp(HOST:DBPORT)/YOUR_DB_NAME?charset=utf8mb4&parseTime=True&loc=Local

3. **Install dependencies**

   go mod tidy

4. **Run the server**

   go run main.go



📌 API Endpoints
Public Routes
POST /login – Authenticate and receive a JWT token

GET /validate – Validate token (JWT required)

Protected Routes (JWT Required)
Users
GET /users

GET /users/:id

POST /users

PUT /users/:id

DELETE /users/:id

Vendors
GET /vendors

GET /vendors/:id

POST /vendors

PUT /vendors/:id

DELETE /vendors/:id

Areas of Interest (/bidangMinats)
Subjects (/mataKuliahs)
Trainings (/pelatihans)
Certifications (/sertifikasis)
All have standard CRUD endpoints: GET, GET by ID, POST, PUT, and DELETE.
