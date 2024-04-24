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
        console.log("error fsdhfkjsdf",error)
    })
}

function displayQuestions(questions) {
    console.log("blablabla", questions)
    const questionsList = document.getElementById("questionsList");

    questions.forEach(question => {
        const questionElement = document.createElement("div");
        questionElement.className='question'
        questionElement.innerHTML = `
            <h2>${question.Name}</h2>
            <p>${question.Description}</p>
            <ul>
                ${question.Variants.map(variant => `<li>${variant}</li>`).join('')}
            </ul>
        `;
        questionsList.appendChild(questionElement);
    });
}