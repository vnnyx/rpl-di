const DB_QUIZ = [
  {
    question: `Alfi berenang setiap 6 hari sekali, Bisma setiap 8 hari sekali dan Charlie setiap
    12 hari sekali. Pada tanggal 19 Februari 2012 mereka berenang bersama-
    sama untuk pertama kali. Mereka akan berenang bersama untuk kedua kalinya
    pada tanggal.....`,
    answer: ['12 Maret 2012', '13 Maret 2012', '14 Maret 2012', '15 Maret 2012']
  },
  {
    question: `Alfi berenang setiap 7 hari sekali, Bisma setiap 8 hari sekali dan Charlie setiap
    12 hari sekali. Pada tanggal 19 Februari 2012 mereka berenang bersama-
    sama untuk pertama kali. Mereka akan berenang bersama untuk kedua kalinya
    pada tanggal.....`,
    answer: ['12 Maret 2012', '13 Maret 2012', '14 Maret 2012', '15 Maret 2012']
  },
  {
    question: `Alfi berenang setiap 8 hari sekali, Bisma setiap 8 hari sekali dan Charlie setiap
    12 hari sekali. Pada tanggal 19 Februari 2012 mereka berenang bersama-
    sama untuk pertama kali. Mereka akan berenang bersama untuk kedua kalinya
    pada tanggal.....`,
    answer: ['12 Maret 2012', '13 Maret 2012', '14 Maret 2012', '15 Maret 2012']
  },
  {
    question: `Alfi berenang setiap 9 hari sekali, Bisma setiap 8 hari sekali dan Charlie setiap
    12 hari sekali. Pada tanggal 19 Februari 2012 mereka berenang bersama-
    sama untuk pertama kali. Mereka akan berenang bersama untuk kedua kalinya
    pada tanggal.....`,
    answer: ['12 Maret 2012', '13 Maret 2012', '14 Maret 2012', '15 Maret 2012']
  },
  {
    question: `Alfi berenang setiap 9 hari sekali, Bisma setiap 8 hari sekali dan Charlie setiap
    12 hari sekali. Pada tanggal 19 Februari 2012 mereka berenang bersama-
    sama untuk pertama kali. Mereka akan berenang bersama untuk kedua kalinya
    pada tanggal.....`,
    answer: ['12 Maret 2012', '13 Maret 2012', '14 Maret 2012', '15 Maret 2012']
  },
  {
    question: `Alfi berenang setiap 9 hari sekali, Bisma setiap 8 hari sekali dan Charlie setiap
    12 hari sekali. Pada tanggal 19 Februari 2012 mereka berenang bersama-
    sama untuk pertama kali. Mereka akan berenang bersama untuk kedua kalinya
    pada tanggal.....`,
    answer: ['12 Maret 2012', '13 Maret 2012', '14 Maret 2012', '15 Maret 2012']
  },
  {
    question: `Alfi berenang setiap 9 hari sekali, Bisma setiap 8 hari sekali dan Charlie setiap
    12 hari sekali. Pada tanggal 19 Februari 2012 mereka berenang bersama-
    sama untuk pertama kali. Mereka akan berenang bersama untuk kedua kalinya
    pada tanggal.....`,
    answer: ['12 Maret 2012', '13 Maret 2012', '14 Maret 2012', '15 Maret 2012']
  },
  {
    question: `Alfi berenang setiap 9 hari sekali, Bisma setiap 8 hari sekali dan Charlie setiap
    12 hari sekali. Pada tanggal 19 Februari 2012 mereka berenang bersama-
    sama untuk pertama kali. Mereka akan berenang bersama untuk kedua kalinya
    pada tanggal.....`,
    answer: ['12 Maret 2012', '13 Maret 2012', '14 Maret 2012', '15 Maret 2012']
  },
  {
    question: `Alfi berenang setiap 9 hari sekali, Bisma setiap 8 hari sekali dan Charlie setiap
    12 hari sekali. Pada tanggal 19 Februari 2012 mereka berenang bersama-
    sama untuk pertama kali. Mereka akan berenang bersama untuk kedua kalinya
    pada tanggal.....`,
    answer: ['12 Maret 2012', '13 Maret 2012', '14 Maret 2012', '15 Maret 2012']
  },
  {
    question: `Alfi berenang setiap 9 hari sekali, Bisma setiap 8 hari sekali dan Charlie setiap
    12 hari sekali. Pada tanggal 19 Februari 2012 mereka berenang bersama-
    sama untuk pertama kali. Mereka akan berenang bersama untuk kedua kalinya
    pada tanggal.....`,
    answer: ['12 Maret 2012', '13 Maret 2012', '14 Maret 2012', '15 Maret 2012']
  },
  {
    question: `Alfi berenang setiap 9 hari sekali, Bisma setiap 8 hari sekali dan Charlie setiap
    12 hari sekali. Pada tanggal 19 Februari 2012 mereka berenang bersama-
    sama untuk pertama kali. Mereka akan berenang bersama untuk kedua kalinya
    pada tanggal.....`,
    answer: ['12 Maret 2012', '13 Maret 2012', '14 Maret 2012', '15 Maret 2012']
  },
  {
    question: `Alfi berenang setiap 9 hari sekali, Bisma setiap 8 hari sekali dan Charlie setiap
    12 hari sekali. Pada tanggal 19 Februari 2012 mereka berenang bersama-
    sama untuk pertama kali. Mereka akan berenang bersama untuk kedua kalinya
    pada tanggal.....`,
    answer: ['12 Maret 2012', '13 Maret 2012', '14 Maret 2012', '15 Maret 2012']
  },
]

const CORRECT_ANSWER = [0, 1, 2, 3, 0, 0, 0, 0, 0, 0, 0, 0];

// -1 thats mean user didn't answer yet
let user_answer = [-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1];

let current_question = 0;

document.addEventListener("DOMContentLoaded", function(event){
  checkTotalQuestion();
  setupQuestion();
});

function setupQuestion(){

  isAnswered();

  document.getElementById("question").innerHTML = (current_question+1+". ")+DB_QUIZ[current_question]['question'];
  document.getElementById("choiceText0").innerText = DB_QUIZ[current_question]['answer'][0];
  document.getElementById("choiceText1").innerText = DB_QUIZ[current_question]['answer'][1];
  document.getElementById("choiceText2").innerText = DB_QUIZ[current_question]['answer'][2];
  document.getElementById("choiceText3").innerText = DB_QUIZ[current_question]['answer'][3];

  if(current_question == DB_QUIZ.length - 1){
    document.getElementById("btn_next").disabled = true;
  }
  else{
    document.getElementById("btn_next").disabled = false;
  }
}

function nextQuestion(){
  console.log(current_question);
  console.log(DB_QUIZ.length-1)
  
  if(current_question < DB_QUIZ.length-1){
    current_question++;
    var id_btn = `selectQuestion${current_question}`;
    document.getElementById(id_btn).checked = true;
    setupQuestion();
  }
}

function goToQuestion(n) {
  current_question = n;
  setupQuestion();
}

function checkTotalQuestion() {
  for (let i = 0; i < DB_QUIZ.length; i++) {
    var id_btn = `selectQuestion${i}`;
    document.getElementById(id_btn).disabled = false;
  }
  document.getElementById('selectQuestion0').checked = true;
}

function isAnswered(){
  document.querySelectorAll("input[name='choices']").forEach(a => a.checked = false);
  let ans = user_answer[current_question] 
  if(ans != -1){
    var id_choice = `choice${ans}`;
    document.getElementById(id_choice).checked = true;
  }
}

function saveAnswer(){
  const answer = document.querySelector("input[name='choices']:checked");
  if(answer != null) {
    user_answer[current_question] = parseInt(answer.getAttribute("data-id"))
  }
  console.log(user_answer)
}

function getScoreAndAnswer(){
  let result = []
  let score = 0;
  let correct = 0;
  let wrong = 0;
  let plus = 100 / DB_QUIZ.length;
  for(let i = 0; i < CORRECT_ANSWER.length; i++){
    var btn = document.getElementById(`selectQuestionBtn${i}`);
    if(user_answer[i] == CORRECT_ANSWER[i]){
      score += plus;
      correct += 1;
      btn.classList.remove('btn-outline-warning');
      btn.classList.add('btn-success');
    }
    else{
      btn.classList.remove('btn-outline-warning');
      btn.classList.add('btn-danger');
      wrong += 1;
    }
  }
  result.push(parseInt(score.toFixed()))
  result.push(correct);
  result.push(wrong);

  return result;
}

function getPredicate(score){
  let predicate = "";
  score = parseInt(score);
  if(score < 70){
    predicate = "Study Harder";
  }
  else if(score > 70 && score < 85){
    predicate = "Good but still have to learn more"
  }
  else if(score > 85 && score < 99){
    predicate = "Almost perfect";
  }
  else if(score >= 100){
    predicate = "Perfect"
  }
  console.log(predicate)
  return predicate;
}


function stopQuiz(){
  stopTimer();
  let finBtn = document.getElementById('finish-btn');
  finBtn.innerText = "See Result";
  finBtn.removeAttribute('data-bs-toggle');
  finBtn.removeAttribute('data-bs-target');
  finBtn.setAttribute('onclick', 'stopQuiz()');

  document.querySelectorAll("input[name='choices']").forEach(a => a.disabled = true);
  let data = getScoreAndAnswer();
  let score = data[0];
  let correct = data[1];
  let wrong = data[2];
  let predicate = getPredicate(score);
  document.getElementById('predicate').innerText = `${predicate}`;
  document.getElementById('score').innerText = `${score}`;
  document.getElementById('correct').innerText = `${correct}`;
  document.getElementById('wrong').innerText = `${wrong}`;
  $('#staticBackdrop').modal('show');
  return
}

function Ask(){
  var node = document.getElementById('my-node');

  domtoimage.toPng(node)
    .then(function (dataUrl) {
        var img = new Image();
        img.src = dataUrl;
        console.log(img.src);
        localStorage.setItem('photo', dataUrl)
        document.body.appendChild(img);
        location.href = "tanya.html";
    })
    .catch(function (error) {
        console.error('oops, something went wrong!', error);
  });
  window.location
}