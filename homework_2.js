
  // homework //

  // try catch

  const increment = num => {
    if(typeof(num) != 'number') {
      throw new TypeError('paramiter must be a number');    
    } 
    if(num < 0 || num > 10) {
      throw new RangeError('number should be bettwen 0 and 10')
    }

    return ++num;
  }

  try {
    console.log(increment(50))
  } catch(e) {
    console.log(e.name + '-' + e.message)
  }

  // primise 

  const incrementPromis = num => {
    return new Promise((resolve,reject) => {
      if(typeof(num) != 'number') {
        reject(new TypeError('paramiter must be number'));
      }

      if (num < 0 || num > 10) {
        reject(new RangeError("number should be bettwen 0 and 10"));
      }

      resolve(++num)
    })
  }

  incrementPromis(60).then(result => {
    console.log(result);
  }).catch(e => {
    console.log(e.name + "-" + e.message);
  })

  // sending request 

  fetch("https://jsonplaceholder.typicode.com/posts", {
    method: "POST",
    body: JSON.stringify({
      title: "Post About Js",
      body: "body)",
      Id: 1,
    }),
    headers: {
      "Content-type": "application/json; charset=UTF-8",
    },
  })
    .then((response) => response.json())
    .then((json) => console.log('id of the created post is : ' + json.Id));


  // copying objects 

  const Georgiy = {
    name: "Georgiy",
    passport :  {
      number : "3524452",
      date : new Date(),
    }
  };

  function deepObjCopy(obj) {
      return JSON.parse(JSON.stringify(obj));
  }

  const GeorgiyCopy = deepObjCopy(Georgiy);

  Georgiy.passport.number = "new number 42431431"

  console.log(Georgiy);
  console.log(GeorgiyCopy);
  

function GetSuppliers() {
  fetch("http://localhost:8080/suppliers")
    .then((response) => response.json())
    .then((data) => {
      console.log(data); // Process the data received from the server
    })
    .catch((error) => {
      console.error(error); // Handle any errors that occurred during the request
    });
}





