var fill;
var minute;
var seconds
function createTimer(n) {
    const startingMinutes = n;
    var time = startingMinutes * 60;
    var counter = document.getElementById('Timer').getContext('2d');
    var no = time;
    var pointToFill = 4.72;
    var cw = counter.canvas.width;
    var ch = counter.canvas.height;
    var diff;
      
    function fillCounter() {
        const minutes = Math.floor(time / 60);
        seconds = time % 60;
        diff = ((time / no) * Math.PI * 2 * 10);
        seconds = seconds < 10 ? '0' + seconds : seconds;
        minute = minutes < 10 ? '0' + minutes : minutes;
        counter.clearRect(0, 0, cw, ch);
        counter.lineWidth = 30;
        counter.fillStyle = '#1466ff';
        counter.strokeStyle = '#fcc630';
        counter.textAlign = 'center';
        counter.font = "bold 17px Poppins";
        counter.shadowColor = '#c28f00';
        counter.shadowBlur = 5;
        counter.shadowOffsetX = 1;
        counter.shadowOffsetY = 1;
        if (time == 0) {
            counter.fillText(`Time's Up`, 110, 117);
            clearTimeout(fill);
            stopQuiz();
        }
        else{
            counter.fillText(`${minute}:${seconds}`, 110, 117);
        }
        counter.beginPath();
        counter.arc(110, 110, 70, pointToFill, diff / 10 + pointToFill);
        counter.stroke();
        
        time--;
    }
    
    fill = setInterval(fillCounter, 1000);
}

function getTime(){
    var t = [minute,seconds];
    console.log(t);
    return t;
}

function stopTimer(){
    clearTimeout(fill);
}

createTimer(15);