# Go User API 🧪

## 📌 Overview

This project is a **RESTful API** built with Go to manage users with their **name** and **date of birth (DOB)**. The API dynamically calculates a user’s age whenever user data is fetched.

It demonstrates:

- **Clean backend architecture:** Handler → Service → Repository  
- **Fiber web framework** for fast and simple HTTP handling  
- **SQLC** for type-safe database queries  
- **Dynamic age calculation** using Go’s `time` package  

---

## 📂 Project Structure


📂 Project Structure

- **cmd/server/main.go**
- **config/**
- **db/**
  - migrations/
  - sqlc/
- **internal/**
  - handler/       → HTTP request handlers
  - repository/    → Database access
  - service/       → Business logic (age calculation, etc.)
  - routes/        → Route definitions
  - middleware/
  - models/
  - logger/
       

---

## 🔧 Tech Stack

- **Language:** Go (Golang)  
- **Web Framework:** Fiber  
- **Database:** PostgreSQL + SQLC  
- **Logging:** Uber Zap  
- **Validation:** go-playground/validator  

---

## 🗄️ Database Schema

**`users` table:**

| Field | Type   | Constraints    |
|-------|--------|----------------|
| id    | SERIAL | PRIMARY KEY    |
| name  | TEXT   | NOT NULL       |
| dob   | DATE   | NOT NULL       |

---
## 🚀 Running the Project

1. **Clone the repository**
```bash
git clone 
cd go-user-api

2.Set up PostgreSQL database

CREATE DATABASE userdb;
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL
);


3.Run the Go server

go run ./cmd/server/main.go


Server runs on: http://localhost:3000

🔄 API Endpoints
-> Create User
curl -X POST -H "Content-Type: application/json" \
-d "{\"name\":\"Alice\",\"dob\":\"1990-05-10\"}" \
http://localhost:3000/users


Response:
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10",
  "age": 35
}

-> Get User by ID
curl http://localhost:3000/users/1

-> List All Users
curl http://localhost:3000/users

-> Update User
curl -X PUT -H "Content-Type: application/json" \
-d "{\"name\":\"Alice Updated\",\"dob\":\"1991-03-15\"}" \
http://localhost:3000/users/1

-> Delete User
curl -X DELETE http://localhost:3000/users/1

✅ Key Features

Dynamic Age Calculation – No need to store age in DB.

Clean Architecture – Separation of concerns (Handler / Service / Repository).

Type-Safe SQL Queries using SQLC.

Logging & Validation – Robust API with proper error handling.

💡 Notes

API runs on HTTP, not HTTPS. Use http://localhost:3000 in requests.

