console.log("JS Loaded")
const url = "http://localhost:8080/theme.html"

function sendLetter(letter) {
    fetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'text/plain'
      },
      body: letter
    }).then(function(response) {
      return response.text();
    }).then(function(text) {
      console.log(text);
    }).catch(function(error) {
      console.error(error);
    });
  }