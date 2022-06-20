function myFunction() {
  var inp = document.getElementById("password");
  var hide1 = document.getElementById("hide1");
  var hide2 = document.getElementById("hide2");
  if (inp.type == "password") {
    inp.type = "text";
    hide1.style.display = "inline";
    hide2.style.display = "none";
  } else {
    inp.type = "password";
    hide1.style.display = "none";
    hide2.style.display = "inline";
  }
}
function myFunction2() {
  var inp = document.getElementById("passwordConfirm");
  var hide3 = document.getElementById("hide3");
  var hide4 = document.getElementById("hide4");
  if (inp.type == "password") {
    inp.type = "text";
    hide3.style.display = "inline";
    hide4.style.display = "none";
  } else {
    inp.type = "password";
    hide3.style.display = "none";
    hide4.style.display = "inline";
  }
}
function choosedAvatar(avatarID) {
  //   console.log(avatarID);
  document.getElementsByClassName("card-ava-clicked").className = "card-ava";
  document.getElementById(avatarID).className = "card-ava-clicked";
}

$(function () {
  var avaID = "";

  $(".card-ava").click(function () {
    $(".card-ava").removeClass("selected").addClass("unselected");
    $(this).removeClass("unselected").addClass("selected");
    avaID = this.id;
    var img = this.querySelector(".card-ava img");
    img = new URL(img.src).pathname;

    $(".btn-next").click(function (e) {
      const url = new URL(window.location.href);
      var email = url.searchParams.get("email");
      var username = url.searchParams.get("username");
      var handphone = url.searchParams.get("handphone");
      var password = url.searchParams.get("password");
      var confirmPassword = url.searchParams.get("confirmPassword");

      console.log({ email, username, handphone, password, confirmPassword });

      var requestBody = {
        email: email,
        username: username,
        handphone: handphone,
        password: password,
        password_confirmation: confirmPassword,
        avatar: img,
      };

      e.preventDefault();
      $.ajax({
        url: "https://sixmath.vnnyx.my.id/api/user/student",
        type: "post",
        contentType: "application/json; charset=utf-8",
        data: JSON.stringify(requestBody),
        success: function (result) {
          localStorage.setItem(
            "email",
            result.data.email,
            "username",
            result.data.username,
            "handphone",
            result.data.handphone,
            "password",
            result.data.email
          );
          window.location = "../login.html";
        },
        error: function () {
          console.log("ok");
        },
      });
    });
  });
});

$(function () {
  $(".btn").click(function () {
    var password = $("#password").val();
    var confirmPassword = $("#passwordConfirm").val();
    var username = $("#username").val();
    var email = $("#email").val();
    var hp = $("#handphone").val();

    if (password != confirmPassword) {
      $("#err").html("*Password tidak sama!");
      return false;
    }

    window.location = `pilih-ava.html?email=${email}&username=${username}&handphone=${hp}&password=${password}&confirmPassword=${confirmPassword}`;

    return false;
  });
});

// $("#btn-daftar").on("click", function (e) {
//   e.preventDefault();
//   $.ajax({
//     url: "https://sixmath.vnnyx.my.id/api/user/student",
//     type: "post",
//     dataType: "json",
//     data: {
//       email: $("#email").val(),
//       username: $("#username").val(),
//       handphone: $("#handphone").val(),
//       password: $("#myInput").val(),
//     },
//     success: function (result) {
//       localStorage.setItem(
//         "email",
//         result.data.email,
//         "username",
//         result.data.username,
//         "handphone",
//         result.data.handphone,
//         "password",
//         result.data.email
//       );
//       window.location = "login.html";
//     },
//     failure: function () {
//       console.log("ok");
//     },
//   });
// });
