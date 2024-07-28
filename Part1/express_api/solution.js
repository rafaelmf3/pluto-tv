const express = require('express');
const app = express();
app.use(express.json());

const moviesDb = [
    { id: 1, title: "Inception", director: "Christopher Nolan" },
    { id: 2, title: "The Matrix", director: "Lana Wachowski, Lilly Wachowski" },
    { id: 3, title: "Interstellar", director: "Christopher Nolan" },
];

// GET /movies - Retorna uma lista de filmes
app.get('/movies', (_, res) => {
    res.json(moviesDb);
});

// POST /movies - Adiciona um novo filme à lista
app.post('/movies', (req, res) => {
    const newMovie = {};
    const {title, director} = req.body;
    newMovie.title = title ? title : "";
    newMovie.director = director ? director : "";
    if (newMovie.title == "") return res.status(401).json({ message: "missing title" });
    if (newMovie.director == "") return res.status(401).json({ message: "missing director" }); 
    newMovie.id = moviesDb.length ? moviesDb[moviesDb.length - 1].id + 1 : 1;
    moviesDb.push(newMovie);
    res.status(201).json(newMovie);
});

// GET /movies/:id - Retorna os detalhes de um filme por ID
app.get('/movies/:id', (req, res) => {
    const movieId = parseInt(req.params.id, 10);
    const movie = moviesDb.find(m => m.id === movieId);
    if (movie) {
        res.json(movie);
    } else {
        res.status(404).json({ message: "Movie not found" });
    }
});

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
    console.log(`Server is running on port ${PORT}`);
});