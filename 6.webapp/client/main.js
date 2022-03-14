document.onload = getUsers();

function getUsers() {
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var users = JSON.parse(this.responseText);
            document.getElementById("users-count").innerHTML = users.length;
        }
    }
    xhttp.open("GET", "/users/", true);
    xhttp.send();
}

function addUser(e) {
    e.preventDefault();
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        getUsers();
    }
    xhttp.open("POST", "/users/", true);
    xhttp.setRequestHeader("Content-type", "application/json");
    var data = JSON.stringify({
        "username": document.getElementById("username").value,
        "password": document.getElementById("password").value
    })
    xhttp.send(data);
}