fetch('http://localhost:8080/users')
    .then(response => response.json())
    .then(users => {
        const usersContainer = document.querySelector('.users');

        users.forEach(user => {
            const userElement = document.createElement('div');
            userElement.innerHTML = `<p>ID: ${user.id}</p>
                               <p>Username: ${user.username}</p>
                               <p>Age: ${user.age}</p>
                               <p>Password: ${user.password}</p>`;
            usersContainer.appendChild(userElement);
        });
    })
    .catch(error => {
        console.error('Error:', error);
    });