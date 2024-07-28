const express = require('express');
const app = express();
app.use(express.json());

const moviesDb = [
    { id: 1, title: "Inception", director: "Christopher Nolan" },
    { id: 2, title: "The Matrix", director: "Lana Wachowski, Lilly Wachowski" },
    { id: 3, title: "Interstellar", director: "Christopher Nolan" },
];

// GET /movies - Returns a list of movies
app.get('/movies', (req, res) => {
    //Complete the function
});

// POST /movies - Adds a new movie to the list
app.post('/movies', (req, res) => {
    //Complete the function
});

// GET /movies/:id - Returns the details of a movie by ID
app.get('/movies/:id', (req, res) => {
    //Complete the function
});

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
    console.log(`Server is running on port ${PORT}`);
});