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