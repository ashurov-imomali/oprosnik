function Registration(){
    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value;
    var fullname = document.getElementById("fullname").value;
    
    var userData = {
        "username":username,
        "password":password,
        "fullname":fullname
    }

    fetch("http://127.0.0.1:8080/registration", {
    body: JSON.stringify(userData),
    method: "POST"
    }).then(response => {
        if (!response.ok) {
            throw new Error('Ошибка: ' + response.statusText);
        }
        return response.json();
    }).then(data => {
        console.log(data.data); // Вывод данных, без вызова .json()
        localStorage.setItem("user_id", data.data.Id)
        window.location.href = "../html/userpage.html";
    }).catch(error => {
        console.error('Произошла ошибка:', error.message);
        alert("ERROR: " + error.message);
        document.getElementById("username").value = "";
    });

}
