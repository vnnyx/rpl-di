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


// $('#btn-daftar').on('click', function (e) {
//     e.preventDefault();
//     var data = new FormData();

//     var avatar = document.getElementById("avatar-guru");
//     data.append('email', $('#email').val())
//     data.append('username', $('#username').val())
//     data.append('handphone', $('#handphone').val())
//     data.append('password', $('#myInput').val())
//     data.append('avatar', avatar.files[0])

//     console.log($('#myInput').val())
//     $.ajax({
//         url: 'https://sixmath.vnnyx.my.id/api/user/teacher',
//         type: 'post',
//         enctype: 'multipart/form-data',
//         processData: false,
//         contentType: false,
//         data: data,
//         success: function (result) {
//             window.location = "login.html"
            
//         },
//         error: function () {
//             console.log("Response bla bla bla")
//         }
//     });
// });