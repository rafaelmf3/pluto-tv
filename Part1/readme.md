### Part 1: Node.js Knowledge

**Objective:** Demonstrate proficiency in Node.js concepts, including asynchronous programming and Express.js API development.

#### Tests:

1. **Asynchronous Code:**
    Rewrite the function below from promises to use async/await style. Use a fake API endpoint `https://jsonplaceholder.typicode.com/todos/1` to fetch data:

```javascript
    const axios = require('axios');

    function getData() {
        axios.get('https://jsonplaceholder.typicode.com/todos/1')
            .then(response => {
                console.log(response.data);
            })
            .catch(error => {
                console.error(error);
            });
    }

    getData();
```

2. **Express API:**
    - Create a simple Express.js server that has the following endpoints:
        - `GET /movies` - Returns a list of movies from the provided array.
        - `POST /movies` - Adds a new movie to the list.
        - `GET /movies/:id` - Returns the details of a movie by ID.

```javascript
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
```

3. **Concurrent Requests:**
    Change this code so instead of doing 10 requests one after the other, execute them all at once and log the response:

```javascript
    const axios = require('axios');

    async function makeRequests() {
        const urls = [
            'https://api.example.com/data1',
            'https://api.example.com/data2',
            'https://api.example.com/data3',
            'https://api.example.com/data4',
            'https://api.example.com/data5',
            'https://api.example.com/data6',
            'https://api.example.com/data7',
            'https://api.example.com/data8',
            'https://api.example.com/data9',
            'https://api.example.com/data10'
        ];

        for (const url of urls) {
            console.log((await axios.get(url)).data)
        }
    }
```