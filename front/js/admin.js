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


// Добавление кнопки для добавления нового вопроса
function addNewQuestionButton() {
    const addButton = document.createElement("button");
    addButton.textContent = "Добавить новый вопрос";
    addButton.addEventListener("click", function() {
    const questionElement = document.createElement("div");
    questionElement.className = 'editQuestions';
    questionElement.innerHTML = `
        <div class="someName">
            <div class="delete-gifka">
                <input type="text" class="input-name">
                <button onclick="addSubVariant()">+</button>
                <img src="../images/delete.png" class="images">
            </div>
            <input type="text" class="input-description">
            <div id="subVariant${333}"></div>
        </div>`;
        document.getElementById("editQuestionList").appendChild(questionElement)
    });
    const saveButton = document.createElement("button");
    saveButton.textContent = "Сохранить";
    document.getElementById("editQuestionList").appendChild(addButton);
    document.getElementById("editQuestionList").appendChild(saveButton);

}

function addSubVariant(variantId){
        const questionElement = document.createElement("div");
        questionElement.className = 'editQuestions';
        questionElement.innerHTML = `
                <input type="text" class="check-variant">
          `;
            document.getElementById("subVariant").appendChild(questionElement)
}

// Функция для отображения списка вопросов
function displayQuestions(questions) {
    const questionsList = document.getElementById("editQuestionList");
    questionsList.innerHTML = ""; // Очищаем список вопросов перед добавлением новых

    questions.forEach((question, index) => {
        const questionElement = createQuestionElement(question, index);
        questionsList.appendChild(questionElement);
    });
    
    const deleteButtons = document.querySelectorAll('.images');
    deleteButtons.forEach(button => {
        button.addEventListener('click', function() {
            const index = parseInt(this.getAttribute('data-index'));
            const questionElement = this.closest('.editQuestions');
            if (questionElement) {
                questionElement.classList.add('fade-out');
                questionElement.addEventListener('animationend', function() {
                    questions.splice(index, 1); // Удаляем элемент из массива
                    displayQuestions(questions); // Обновляем отображение списка вопросов
                });
            }
        });
    });

    addNewQuestionButton(); // Добавляем кнопку для добавления нового вопроса
}

// Функция для создания элемента вопроса
function createQuestionElement(question, index) {
    const questionElement = document.createElement("div");
    questionElement.className = 'editQuestions';
    questionElement.innerHTML = `
        <div class="someName">
            <div class="delete-gifka">
                <input type="text" class="input-name" value="${question.Name}">
                <img src="../images/delete.png" class="images" data-index="${index}">
            </div>
            <input type="text" class="input-description" value="${question.Description}">
            <div class="editVariants">
                ${!(question.Variants === null || question.Variants === undefined) ?
                    `
                    <h3> Варианты </h3>
                    ${question.Variants.map(variant => `
                        <input type="text" class="input-variant" value="${variant}">
                    `).join('')}` : ''}
            </div>
        </div>`;
    return questionElement;
}

// Пример массива вопросов


// Инициализация отображения вопросов при загрузке страницы
//displayQuestions(questions);


  


