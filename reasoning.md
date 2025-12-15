My Approach & Design Decisions

This document explains how I built the API, why I chose certain tools, and the challenges I faced while working on it.

1. Choosing the Tech Stack

I’m relatively new to Go, so I wanted tools that are simple but also powerful.

For the web framework, I chose Fiber because it is fast and easy to use. It made defining routes and handling requests straightforward compared to writing everything using the Go standard library.

For the database, I used SQLC. It lets me write SQL queries directly and generates type-safe Go code automatically. This was helpful because I could see exactly what the queries were doing and avoid runtime errors.

2. Handling Age Calculation

One key requirement was returning a user’s age. I chose not to store age in the database because it changes every year. Storing it would require extra work to keep it updated.

Instead, I store the Date of Birth (DOB), which never changes. In the Service layer, I calculate the age dynamically using Go’s time package by comparing the DOB to the current date. This ensures the age is always correct.

3. Project Structure

I organized the project into layers for clarity:

Handler – Handles HTTP requests and responses

Service – Contains business logic, like calculating age

Repository – Handles database queries

This makes debugging easier. For example, if the JSON response is wrong, I check the Handler. If the age calculation is wrong, I check the Service.

4. Challenges & Learnings

Working with dates was tricky. PostgreSQL has its own date format, and Go uses time.Time. Mapping them correctly took some trial and error.

Overall, this project taught me a lot about structuring a backend, working with strict types in Go, and building maintainable code.