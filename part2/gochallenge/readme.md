### Part 2: Go Challenge

**Objective:** Demonstrate proficiency in Go by building a RESTful API that exposes information about movies.

#### Instructions:

You are required to build an API project using Go that exposes information about movies. The API should follow the RESTful approach and should include endpoints to handle movies.

Use https://www.omdbapi.com or https://developer.themoviedb.org/reference/intro/getting-started to retrieve movie data

#### API Requirements:

1. **Endpoints:**
-  **GET /movies** - Returns a list of movies.
-  **GET /movies/{id}** - Returns details of a movie by ID.
-  **POST /movies** - Add a movie to your in memory cache
2. **Cache**: After resources are retrieved, cache it in memory and make sure it's concurrent safe
3. **Unit Tests:** Implement unit tests for what you think it's essential for this API