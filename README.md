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

```bash
SECRET_KEY=YOUR_SECRET_KEY
PORT=:8080
DB=DBUSERNAME:DBPASSWORD@tcp(HOST:DBPORT)/YOUR_DB_NAME?charset=utf8mb4&parseTime=True&loc=Local
```

3. **Install dependencies**

```bash
go mod tidy
```

4. **Run the server**
   
```bash
go run main.go
```



## 📌 API Endpoints

### 🔓 Public Routes

| Method | Endpoint      | Description                   |
|--------|---------------|-------------------------------|
| POST   | `/login`      | Authenticate user and return JWT token |
| GET    | `/validate`   | Validate current JWT token     |

---

### 🔐 Protected Routes (Require JWT)

#### 👤 Users

| Method | Endpoint        | Description              |
|--------|-----------------|--------------------------|
| GET    | `/users`        | Get all users            |
| GET    | `/users/:id`    | Get user by ID           |
| POST   | `/users`        | Create a new user        |
| PUT    | `/users/:id`    | Update user by ID        |
| DELETE | `/users/:id`    | Delete user by ID        |

#### 🏢 Vendors

| Method | Endpoint          | Description              |
|--------|-------------------|--------------------------|
| GET    | `/vendors`        | Get all vendors          |
| GET    | `/vendors/:id`    | Get vendor by ID         |
| POST   | `/vendors`        | Create a new vendor      |
| PUT    | `/vendors/:id`    | Update vendor by ID      |
| DELETE | `/vendors/:id`    | Delete vendor by ID      |

#### 🧠 Areas of Interest (`bidangMinats`)

| Method | Endpoint               | Description                   |
|--------|------------------------|-------------------------------|
| GET    | `/bidangMinats`        | Get all areas of interest     |
| GET    | `/bidangMinats/:id`    | Get area of interest by ID    |
| POST   | `/bidangMinats`        | Create new area of interest   |
| PUT    | `/bidangMinats/:id`    | Update area of interest by ID |
| DELETE | `/bidangMinats/:id`    | Delete area of interest by ID |

#### 📚 Subjects (`mataKuliahs`)

| Method | Endpoint              | Description                |
|--------|-----------------------|----------------------------|
| GET    | `/mataKuliahs`        | Get all subjects           |
| GET    | `/mataKuliahs/:id`    | Get subject by ID          |
| POST   | `/mataKuliahs`        | Create a new subject       |
| PUT    | `/mataKuliahs/:id`    | Update subject by ID       |
| DELETE | `/mataKuliahs/:id`    | Delete subject by ID       |

#### 🏆 Trainings (`pelatihans`)

| Method | Endpoint            | Description              |
|--------|---------------------|--------------------------|
| GET    | `/pelatihans`       | Get all trainings        |
| GET    | `/pelatihans/:id`   | Get training by ID       |
| POST   | `/pelatihans`       | Create a new training    |
| PUT    | `/pelatihans/:id`   | Update training by ID    |
| DELETE | `/pelatihans/:id`   | Delete training by ID    |

#### 📜 Certifications (`sertifikasis`)

| Method | Endpoint              | Description                  |
|--------|-----------------------|------------------------------|
| GET    | `/sertifikasis`       | Get all certifications       |
| GET    | `/sertifikasis/:id`   | Get certification by ID      |
| POST   | `/sertifikasis`       | Create a new certification   |
| PUT    | `/sertifikasis/:id`   | Update certification by ID   |
| DELETE | `/sertifikasis/:id`   | Delete certification by ID   |

