// Get the modal
var modal = document.getElementById("myModal");

// Get the button that opens the modal
var btn = document.getElementById("myBtn");

// Get the <span> element that closes the modal
var span = document.getElementsByClassName("close")[0];

// When the user clicks the button, open the modal 
btn.onclick = function() {
  modal.style.display = "block";
}

// When the user clicks on <span> (x), close the modal
span.onclick = function() {
  modal.style.display = "none";
}

// When the user clicks anywhere outside of the modal, close it
window.onclick = function(event) {
  if (event.target == modal) {
    modal.style.display = "none";
  }
}

function easyClicked() {
    // Récupère l'état du bouton "Easy"
    var easyCheckbox = document.querySelector('input[id="easyCheckbox"]');
    
    if (easyCheckbox.checked) {
      // Si le bouton "Easy" est coché, désactive les boutons "Medium" et "Hard"
      document.getElementById("mediumCheckbox").disabled = true;
      document.getElementById("hardCheckbox").disabled = true;
    } else {
      // Si le bouton "Easy" est décoché, active les boutons "Medium" et "Hard"
      document.getElementById("mediumCheckbox").disabled = false;
      document.getElementById("hardCheckbox").disabled = false;
  
      // Ajouter ce code pour activer le bouton "Easy" automatiquement
      document.getElementById("easyCheckbox").disabled = false;
    }
  }
  
  function mediumClicked() {
    // Récupère l'état du bouton "Medium"
    var mediumCheckbox = document.querySelector('input[id="mediumCheckbox"]');
    
    if (mediumCheckbox.checked) {
      // Si le bouton "Medium" est coché, désactive les boutons "Easy" et "Hard"
      document.getElementById("easyCheckbox").disabled = true;
      document.getElementById("hardCheckbox").disabled = true;
    } else {
      // Si le bouton "Medium" est décoché, active les boutons "Easy" et "Hard"
      document.getElementById("easyCheckbox").disabled = false;
      document.getElementById("hardCheckbox").disabled = false;
  
      // Ajouter ce code pour activer le bouton "Medium" automatiquement
      document.getElementById("mediumCheckbox").disabled = false;
    }
  }
  
  function hardClicked() {
    // Récupère l'état du bouton "Hard"
    var hardCheckbox = document.querySelector('input[id="hardCheckbox"]');
    
    if (hardCheckbox.checked) {
      // Si le bouton "Hard" est coché, désactive les boutons "Easy" et "Medium"
      document.getElementById("easyCheckbox").disabled = true;
      document.getElementById("mediumCheckbox").disabled = true;
    } else {
      // Si le bouton "Hard" est décoché, active les boutons "Easy" et "Medium"
      document.getElementById("easyCheckbox").disabled = false;
      document.getElementById("mediumCheckbox").disabled = false;
  
      // Ajouter ce code pour activer le bouton "Hard" automatiquement
      document.getElementById("hardCheckbox").disabled = false;
    }
  }