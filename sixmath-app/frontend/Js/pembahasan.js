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
    question: `Bu Umi mempunyai 54 kue bronis dan 63 kue donat. Kue tersebut akan dimasukkan ke dalam kardus dengan jumlah yang sama banyak. Maka kardus yang digunakan ibu sebanyak.....`,
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

const DB_PEMBAHASAN = [
  {
    content: `Faktorisasi prima dari 693 adalah 3 x 3 x 7 x 11, maka faktor primanya adalah 3, 7 dan 11.`
  },
  {
    content: `kita misalkan 630 = 63 x 10

    63 = 9 x 7 = 3 x 3 x 7
    
    10 = 2 x 5
    
    jika digabungkan menjadi 2 x 3 x 3 x 5 x 7`,
  },
  {
    content: `48 = 2 x 2 x 2 x 2 x 3 = 2 pangkat 4 x 3

    72 = 2 x 2 x 2 x 3 x 3 = 2 pangkat 3 x 3 pangkat 2
    
    96 = 2 x 2 x 2 x 2 x 2 x 3 = 2 pangkat lima x 3
    
    KPK dari 48, 72 dan 96 = 2 x 2 x 2 x 2 x 2 x 3 x 3 atau 2 pangkat 5 x 3 pangkat 2 = 32 x 9 = 288.`,
  },
  {
    content: `untuk menyelesaikan soal cerita KPK dan FPB di atas dengan memakai mencari FPB yakni FPB dari 54 dan 63.

    54 = 9 x 6 = (3 x 3) x (2 x 3) atau 2 x 3 pangkat 3
    
    63 = 9 x 7 = (3 x 3) x 7 atau 3 pangkat 2 x 7
    
    maka FPB dari 54 dan 63 adalah 9.`,
  },
  {
    content: `36 = 4 x 9 atau (2 x 2 ) x ( 3 x 3 ) = 2 pangkat 2 x 3 pangkat 2

    48 = 8 x 6 atau ( 2 x 2 x 2 ) x ( 2 x 3 ) = 2 pangat 4 x 3
    
    72 = 8 x 9 atau ( 2 x 2 x 2 ) x ( 3 x 3 ) = 2 pangkat 3 x 3 pangkat 2
    
    KPK dari 36, 48 dan 72 adalah 2 pangkat 4 x 3 pangkat 2 = 16 x 9 = 144`,
  },
  {
    content: `48 = 2 x 2 x 2 x 2 x 3

    90 = 2 x 3 x 3 x 5
    
    FPB dari 48 dan 90 = 2 x 3 = 6.`,
  },
  {
    content: `Untuk menyelesaikan soal cerita di atas dengan FPB yakni FPB dari 60 dan 72

    60 = 6 x 10 atau (2 x 3) x (2 x 5) = 2 x 2 x 3 x 5 atau 2 pangkat 2 x 3 x 5
    
    72 = 2 x 2 x 2 x 3 x 3
    
    maka FPB dari 60 dan 72 adalah 2 x 2 x 3 = 12.`,
  },
  {
    content: `Faktorisasi dari 528 = 2 x 2 x 2 x 2 x 3 x 11, maka faktor prima dari 528 adalah 2, 3 dan 11.`,
  },
  {
    content: `36 = 4 x 9 atau (2 x2) x (3 x 3) = 2 x 2 x 3 x 3

    60 = 6 x 10 atau (2 x 3) x (2 x 5) = 2 x 2 x 3 x 5
    
    maka KPK dari 36 dan 60 adalah 2 x 2 x 3 x 3 x 5 = 4 x 9 x 5 = 36 x 5 = 180.`,
  },
  {
    content: `40 = 4 x 10 atau (2 x 2) x (2 x 5) = 2 x 2 x 2 x 5

    48 = 6 x 8 atau (2 x 3) x (2 x 2 x 2) = 2 x 2 x 2 x 2 x 2 x 3
    
    56 = 7 x 8 atau 7 x (2 x 2 x 2) = 2 x 2 x 2 x 7
    
    maka FPB dari 40, 48 dan 56 = 2 x 2 x 2 = 8.`,
  },
  {
    content: `144 = 12 x 12 atau (2 x 2 x 3) x (2 x 2 x 3) = 2 x 2 x 2 x 2 x 3 x 3

    192 = 3 x 8 x 8 = 3 x (2 x 2 x 2) x (2 x 2 x 2) = 2 x 2 x 2 x 2 x 2 x 2 x 3
    
    Sehingga KPK dari 144 dan 192 adalah 2 x 2 x 2 x 2 x 2 x 2 x 3 x 3 = 64 x 9 = 576`,
  },
  {
    content: `98 = 2 x 49 =  2 x (7 x 7) = 2 x 7 x 7

    112 = 2 x 56 = 2 x (7 x 8) = 2 x 7 x 2 x 2 x 2
    
    maka FPB dari 98 dan 112 adalah 2 x 7 = 14.`,
  },
];

const CORRECT_ANSWER = [3, 0, 2, 0, 1, 0, 0, 0, 2, 3, 1, 3];

// -1 thats mean user didn't answer yet
let user_answer = [-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1];

let current_question = 0;

document.addEventListener("DOMContentLoaded", function (event) {
  checkTotalQuestion();
  setUserAnswer();
  setupQuestion();
  stopQuiz();
});

function setupQuestion() {
  // isAnswered();
  resetCorrect();
  document.getElementById("question").innerHTML =
    current_question + 1 + ". " + DB_QUIZ[current_question]["question"];
  for (i = 0; i < 4; i++) {
    var choice = document.getElementById(`choiceText${i}`);
    choice.innerText = DB_QUIZ[current_question]["answer"][i];
  }
  setCorrect();

  if (current_question == DB_QUIZ.length - 1) {
    document.getElementById("btn_next").disabled = true;
  } else {
    document.getElementById("btn_next").disabled = false;
  }
}

function resetCorrect() {
  for (i = 0; i < 4; i++) {
    var choice = document.getElementById(`choiceText${i}`);
    var icon = document.getElementById(`choice-icon${i}`);
    if (choice.classList.contains("btn-outline-success")) {
      choice.classList.replace("btn-outline-success", "btn-outline-primary");
    } else if (choice.classList.contains("btn-outline-danger")) {
      choice.classList.replace("btn-outline-danger", "btn-outline-primary");
    }
    if (icon != null) icon.remove();
  }
}

function setCorrect() {
  var choice = document.getElementById(
    `choiceText${CORRECT_ANSWER[current_question]}`
  );
  choice.classList.replace("btn-outline-primary", "btn-outline-success");
  choice.innerHTML += ` <i id="choice-icon${x}" class="text-success fa fa-check-circle" aria-hidden="true"><i>`;
  console.log(user_answer);
  if (user_answer[current_question] != CORRECT_ANSWER[current_question]) {
    var x = user_answer[current_question];
    console.log(x);
    var choice = document.getElementById(`choiceText${x}`);
    if (choice != null) {
      choice.classList.replace("btn-outline-primary", "btn-outline-danger");
      choice.innerHTML += ` <i id="choice-icon${x}" class="text-danger fa fa-times-circle" aria-hidden="true"><i>`;
    }
  }
  isAnswered();
}

function nextQuestion() {
  console.log(current_question);
  console.log(DB_QUIZ.length - 1);

  if (current_question < DB_QUIZ.length - 1) {
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
  document.getElementById("selectQuestion0").checked = true;
}

function isAnswered() {
  document
    .querySelectorAll("input[name='choices']")
    .forEach((a) => (a.checked = false));
  let ans = user_answer[current_question];
  if (ans != -1) {
    var id_choice = `choice${ans}`;
    document.getElementById(id_choice).checked = true;
    var choice = document.getElementById(`choiceText${ans}`);
    if(choice.classList.contains("btn-outline-success")){
        choice.classList.replace("btn-outline-success", "btn-outline-primary");
    }
    else{
        choice.classList.replace("btn-outline-danger", "btn-outline-primary");
    }
  }
}

// BACA DISINI BANG, KURANG KASI MODAL PEMBAHASAN AJA

function correctAnswer() {
  document
    .querySelectorAll("input[name='choices']")
    .forEach((a) => (a.checked = false));
  let ans = CORRECT_ANSWER[current_question];
  var id_choice = `choice${ans}`;
  document.getElementById(id_choice).checked = true;
}

function saveAnswer() {
  const answer = document.querySelector("input[name='choices']:checked");
  if (answer != null) {
    user_answer[current_question] = parseInt(answer.getAttribute("data-id"));
  }
  console.log(user_answer);
}

function stopQuiz() {

  document
    .querySelectorAll("input[name='choices']")
    .forEach((a) => (a.disabled = true));
  setWCButton();
  return;
}

function setWCButton() {
  for (let i = 0; i < CORRECT_ANSWER.length; i++) {
    var btn = document.getElementById(`selectQuestionBtn${i}`);
    if (user_answer[i] == CORRECT_ANSWER[i]) {
      btn.classList.remove("btn-outline-warning");
      btn.classList.add("btn-success");
    } else {
      btn.classList.remove("btn-outline-warning");
      btn.classList.add("btn-danger");
    }
  }
}

function createDownload() {
  var node = document.getElementById("node-for-download");
  node.innerHTML += 
  `
    <div class="card container bg-white col-5" >
        <div class="card-body p-4 bg-white" id="download-pembahasan">
            <div id="content-pembahasan" class="p-4 border border-warning rounded"></div>
        </div>
    </div>
  `;
  let content = document.getElementById("content-pembahasan");
  content.innerText = DB_PEMBAHASAN[current_question]['content'];
}

function destroyDownload() {
  var node = document.getElementById("node-for-download");
  node.innerHTML = ""
}

function downloadPembahasan() {
  
  var node = document.getElementById("download-pembahasan");

  domtoimage.toJpeg(node, { quality: 1 })
    .then(function (dataUrl) {
        var link = document.createElement('a');
        link.download = `PembahasanSoal_${current_question}.jpeg`;
        link.href = dataUrl;
        link.click();
    });
    setTimeout(() => {
      destroyDownload();
    }, 500)
    
  }

function setUserAnswer() {
  user_answer = localStorage.getItem("userAnswer").split(",");
  console.log(user_answer);
}

function showPembahasan() {
  let pembahasan = document.getElementById("pembahasan");
  pembahasan.innerText = DB_PEMBAHASAN[current_question]['content'];
  $("#modal-pembahasan").modal("show");
  createDownload();
}