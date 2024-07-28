const axios = require('axios');

async function makeRequests() {
    const urls = [
        'https://jsonplaceholder.typicode.com/todos/1',
        'https://jsonplaceholder.typicode.com/todos/2',
        'https://jsonplaceholder.typicode.com/todos/3',
        'https://jsonplaceholder.typicode.com/todos/4',
        'https://jsonplaceholder.typicode.com/todos/5',
        'https://jsonplaceholder.typicode.com/todos/6',
        'https://jsonplaceholder.typicode.com/todos/7',
        'https://jsonplaceholder.typicode.com/todos/8',
        'https://jsonplaceholder.typicode.com/todos/9',
        'https://jsonplaceholder.typicode.com/todos/10'
    ];

    for (const url of urls) {
        console.log((await axios.get(url)).data)
    }
}

makeRequests();