document.addEventListener("DOMContentLoaded", function() {
    var loginForm = document.getElementById("loginForm");
    if(loginForm){
    loginForm.addEventListener("submit", function(event) {
        event.preventDefault(); // Предотвратить отправку формы
        
        var username = document.getElementById("username").value;
        var password = document.getElementById("password").value;
        
        var formData = new FormData();
        formData.append("username", username)
        formData.append("password", password)
        
        var userData = {
            "username":username,
            "password":password
        }

        fetch("http://127.0.0.1:8080/login",{
            body: JSON.stringify(userData),
            method: "POST"
        }).then((response) => {
         response.json().then((data) => {
            console.log(data.role);
            if (data.role == 'admin'){
                localStorage.setItem("user_id", data.user_id)
                window.location.href = 'userpage.html';
            }else{
                
            }
            }).catch((err) => {
            console.log(err);
            })
       .catch(error => {
            console.log("error fsdhfkjsdf",error)
        })})
        // Здесь можно добавить логику для проверки введенных данных или отправки запроса на сервер для аутентификации

        console.log("Логин с именем пользователя: " + username + " и паролем: " + password);
    });
    }
});
