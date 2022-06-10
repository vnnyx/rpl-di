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
  console.log($('#username').val())
  console.log($('#password').val())
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
      console.log("tes")
      console.log(result)
      localStorage.setItem(
        'access_token', result.data.access_token
      )
      window.location = 'video.html'
    },
    failure: function () {
      console.log('ok')
    }
  });
});