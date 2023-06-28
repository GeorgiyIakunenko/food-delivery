const axios = require('axios');

axios.get('/users')
    .then(response => {
        const users = response.data;
        const usersContainer = document.querySelector('.users');

        users.forEach(user => {
            const userElement = document.createElement('div');
            userElement.innerHTML = `<p>ID: ${user.id}</p>
                               <p>Username: ${user.username}</p>
                               <p>Age: ${user.age}</p>
                               <p>Password: ${user.password}</p>`
            usersContainer.appendChild(userElement);
        });
    })
    .catch(error => {
        console.error('Error:', error);
    });
