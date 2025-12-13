# My Approach & Design Decisions

Hi! This document explains how I built the API, the tools I choose, and the challenges I faced while working on this task.

## 1. Choosing the Tech Stack
Since I am relatively new to Go, I wanted tools that were powerful but didn't hide too much of what was happening under the hood.

For the web framework, I chose **Fiber** because it is known for being fast and has a syntax very similar to Express.js. This made it easier for me to define routes and handle requests compared to using the standard library.

For the database interaction, I decided to use **SQLC**. I initially looked at ORMs like GORM, but they felt like too much "magic." I prefer writing actual SQL queries so I know exactly what is happening. SQLC was the perfect middle ground—it let me write raw SQL, but then generated the type-safe Go structs for me. This saved me from many potential spelling mistakes and runtime errors.

## 2. Solving the "Age" Problem
One of the key requirements was to return the user's age. I made a specific design decision **not** to store `age` as a column in the database.

If I stored a static number like "25," that data would become incorrect the moment the user had a birthday. I would need complex background jobs to keep it updated. Instead, I stored the **Date of Birth (DOB)** because that is a constant fact.

In my Service layer, I wrote a logic function that compares the stored DOB with the current date (`time.Now`) to calculate the age dynamically. This ensures the data is 100% accurate every time it is requested, without any extra maintenance.

## 3. Project Structure
I organized the code into layers to keep it clean, rather than putting everything in one main file.

I separated the project into **Handlers** (for HTTP logic), **Services** (for business logic like the age calculation), and **Repositories** (for database access). This separation made debugging much easier. If the JSON response was wrong, I checked the Handler. If the math was wrong, I checked the Service.

## 4. Challenges & Learning
The biggest challenge I faced was handling Date types. PostgreSQL has its own date format, and Go uses `time.Time`. Figuring out how to map these correctly using the driver took some trial and error.

Overall, this project taught me a lot about strict typing in Go and how to structure a professional backend application.