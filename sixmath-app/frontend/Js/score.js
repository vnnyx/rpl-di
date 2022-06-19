function createTimer(n) {
    const Score = n;
    var time = 100;
    var counter = document.getElementById('Timer').getContext('2d');
    var no = time;
    var pointToFill = 4.72;
    var cw = counter.canvas.width;
    var ch = counter.canvas.height;
    var diff;
      
    function fillCounter() {
        diff = ((Score) * Math.PI / 50 + pointToFill);
        console.log((diff))
        counter.clearRect(0, 0, cw, ch);
        counter.lineWidth = 30;
        counter.fillStyle = '#fcc630';
        counter.strokeStyle = '#fcc630';
        counter.textAlign = 'center';
        counter.font = "bold 19px Poppins";

        counter.shadowBlur = 5;
        counter.shadowOffsetX = 1;
        counter.shadowOffsetY = 1;
        counter.fillText(`${Score}`, 98, 117);
        counter.fillStyle = '#828282';
        counter.font = "bold 15px Poppins";
        counter.fillText(`/100`, 128, 120);
        counter.beginPath();
        counter.arc(110, 110, 70, pointToFill, diff);
        counter.stroke();

    }  
    fillCounter();
}
createTimer(localStorage.getItem("score"));