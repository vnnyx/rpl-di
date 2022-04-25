function myFunction() {
    var inp = document.getElementById("myInput")
    var hide1 = document.getElementById("hide1")
    var hide2 = document.getElementById("hide2")
    if (inp.type == "password") {
        inp.type = "text";
        hide1.style.display = "inline"
        hide2.style.display = "none"
    } else {
        inp.type = "password";
        hide1.style.display = "none"
        hide2.style.display = "inline"
    }
}
function myFunction2() {
    var inp = document.getElementById("myInput2")
    var hide3 = document.getElementById("hide3")
    var hide4 = document.getElementById("hide4")
    if (inp.type == "password") {
        inp.type = "text";
        hide3.style.display = "inline"
        hide4.style.display = "none"
    } else {
        inp.type = "password";
        hide3.style.display = "none"
        hide4.style.display = "inline"
    }
}
$(function () {
    $(".btn").click(function () {
        var password = $("#myInput").val();
        var confirmPassword = $("#myInput2").val();
        if (password != confirmPassword) {
            $('#err').html('*Password tidak sama!');
            return false;
        }
        return true;
    });
});

$('#btn-daftar').on('click', function(){
    $.ajax({
        url : 'https://1307-125-164-234-150.ap.ngrok.io/api/student',
        type : 'post',
        dataType : 'json',
        data : {
            'email' : $('#email').val(),
            'username' : $('#username').val(),
            'handphone' : $('#handphone').val(),
            'password' : $('#myInput2').val(),

        },
        success : function(result){
            console.log(result)
            
        },
        failure : function(){
            console.log('ok')
        }
    });
});