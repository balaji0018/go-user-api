# Go User API

A RESTful API built with Go, Fiber, and PostgreSQL to manage users and calculate their age dynamically.

## 📌 Tech Stack
* **Language:** Go (Golang)
* **Framework:** Fiber v2
* **Database:** PostgreSQL
* **Data Access:** SQLC (Type-safe SQL generation)
* **Logging:** Uber Zap

## 🚀 Setup & Run

### 1. Prerequisites
* Go 1.23+
* PostgreSQL

### 2. Database Setup
Create a database named `userdb`. Then, run the following SQL query to create the table:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL
);