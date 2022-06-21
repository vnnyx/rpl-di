function myFunction() {
  var x = document.getElementById("password")
  var y = document.getElementById("hide1")
  var z = document.getElementById("hide2")
  if (x.type == "password") {
    x.type = "text";
    y.style.display = "inline"
    z.style.display = "none"
  } else {
    x.type = "password";
    y.style.display = "none"
    z.style.display = "inline"
  }
}

$('#btn-login').on('click', function (e) {
  e.preventDefault();
  $.ajax({
    url: 'https://sixmath.vnnyx.my.id/api/auth/login',
    type: 'post',
    dataType: 'json',
    contentType: "application/json",
    data: JSON.stringify({
      'username': $('#username').val(),
      'password': $('#password').val(),
    }),
    success: function (result) {
      localStorage.setItem(
        'access_token', result.data.access_token,
      )
      localStorage.setItem(
        'path', result.data.avatar,
      )
      if (result.code == "200") {
        let datas = result.data.role;
        if (datas == "student") {
          window.location = '../../html/student/statistik.html'
        } else if (datas == "teacher") {
          window.location = '../../html/teacher/halaman-create-soal.html'
        } else {
          window.location = '../../html/ortu/dashboard.html'
        }
      }
    },
    error: function () {
      console.log('ok')
    }
  });
});