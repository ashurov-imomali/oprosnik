document.addEventListener("DOMContentLoaded", function() {
    fetchQuestions();
});


function fetchQuestions(){
    fetch("http://localhost:8080/questions",{
        method: "GET"
    }).then((response) => response.json())
    .then(
        data=>{
            displayQuestions(data.data);
        }
    )
   .catch(error => {
        console.log("error",error)
    })
}

function displayQuestions(questions) {
    const questionsList = document.getElementById("questionsList");
    questionsList.innerHTML = ""; // Очищаем список вопросов перед добавлением новых
  
    questions.forEach(question => {
      const questionElement = document.createElement("div");
      questionElement.className = 'question';
      questionElement.innerHTML = `
          <h2>${question.Name}</h2>
          <p>${question.Description}</p>
          <div class="variants">
              ${question.Variants === null || question.Variants === undefined ? 
                  `<input type="text" placeholder="Введите ответ" onchange="changeColor(this)" required>` :
                  `${question.Variants.map(variant => `
                      <input type="radio" id="${variant}" name="${question.Name}" value="${variant}" onchange="changeColor(this)">
                      <label for="${variant}">${variant}</label>
                  `).join('')}`
              }
          </div>
      `;
      questionsList.appendChild(questionElement);
    });
  
    const buttons = document.createElement("div");
    buttons.className = 'buttons';
    buttons.innerHTML =`
        <button type="submit" id="submit-button" onclick="submitForm()">Отправить</button>
        <button type="button" onclick="clearAnswers()">Очистить</button>
    `;
    questionsList.appendChild(buttons);
  }
  
  function changeColor(element) {
    const labels = document.querySelectorAll('label');
    labels.forEach(label => {
      if (label.htmlFor === element.id) {
        label.style.backgroundColor = "#afeeee";
      } else {
        label.style.backgroundColor = "transparent"; // Сброс цвета для других меток
      }
    });
  }
  
  function clearAnswers() {
    const answers = document.querySelectorAll('input[type="text"], input[type="radio"]');
    answers.forEach(answer => {
      if (answer.type === 'text') {
        answer.value = ""; // Очистка поля ввода
      } else if (answer.type === 'radio') {
        answer.checked = false; // Сброс выбранного радио варианта
        const label = document.querySelector(`label[for="${answer.id}"]`);
        label.style.backgroundColor = "transparent"; // Сброс цвета фона у метки
      }
    });
  }
  
  
  function submitForm() {
    const formData = {};
    const inputs = document.querySelectorAll('input[type="text"], input[type="radio"]:checked');
    inputs.forEach(input => {
      formData[input.name] = input.value;
    });

    console.log(formData);
  }