const DB_QUIZ = [
  {
    question: `Faktor prima dari bilangan 693 adalah.....`,
    answer: ['3, 5, 7', '3, 5, 9', '3, 5, 11', '3, 7, 11']
  },
  {
    question: `Faktorisasi prima dari 630 adalah.....`,
    answer: ['2 x 3 x 3 x 5 x 7', '2 x 2 x 3 x 3 x 5 x 7', '2 x 3 x 3 x 5 x 5 x 7', '2 x 2 x 2 x 5 x 7']
  },
  {
    question: `KPK dari bilangan 48, 72, dan 96 adalah.....`,
    answer: ['488', '388', '288', '188']
  },
  {
    question: `Bu Umi mempunyai 54 kue bronis dan 63 kue donat. Kue tersebut akan dimasukkan ke dalam kardus dengan jumlah yang sama banyak. 
    Maka kardus yang digunakan ibu sebanyak.....`,
    answer: ['9', '7', '5', '3']
  },
  {
    question: `KPK dari 36, 48 dan 72 adalah.....`,
    answer: ['134', '144', '154', '164']
  },
  {
    question: `FPB dari 48 dan 90 adalah.....`,
    answer: ['6', '8', '10', '12']
  },
  {
    question: `Ningrum memperoleh oleh oleh dari ibunya berupa 60 coklat dan 72 permen. Coklat dan permen tersebut akan dibungkus dan dibagikan kepada temannya dengan jumlah yang sama banyak. Maka banyak plastik yang dibutuhkan untuk membungkus adalah.....`,
    answer: ['12', '14', '16', '18']
  },
  {
    question: `Faktor prima dari bilangan 528 adalah.....`,
    answer: ['2, 3, 11', '2, 5, 11', '2, 7, 11', '2, 9, 11']
  },
  {
    question: `KPK dari bilangan 36 dan 60 adalah.....`,
    answer: ['120', '160', '180', '210']
  },
  {
    question: `FPB dari bilangan 40, 48 dan 56 adalah.....`,
    answer: ['2', '4', '6', '8']
  },
  {
    question: `KPK dari bilangan 144 dan 192 adalah.....`,
    answer: ['566', '576', '586', '596']
  },
  {
    question: `FPB dari bilangan 98 dan 112 adalah.....`,
    answer: ['10', '11', '12', '14']
  },
];

const CORRECT_ANSWER = [3, 0, 2, 0, 1, 0, 0, 0, 2, 3, 1, 3];

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
    // document.getElementById(id_btn).disabled = false;
  }
  // document.getElementById('selectQuestion0').checked = true;
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
  return predicate;
}


function stopQuiz(){
  stopTimer();
  var time = getTime();
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
  localStorage.setItem("score",score);
  localStorage.setItem("soal",DB_QUIZ.length);
  localStorage.setItem("minutes", time[0]);
  localStorage.setItem("seconds", time[1]);
  localStorage.setItem("correct", correct);
  localStorage.setItem("wrong", wrong);
  localStorage.setItem("userAnswer", user_answer);
  localStorage.setItem("rightAnswer", CORRECT_ANSWER);
  location.href = "result.html";
  return
}

