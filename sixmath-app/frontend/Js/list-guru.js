$(document).ready(function () {
  let sidebar = document.querySelector(".sidebar");
  $("#btn").on("click", function () {
    sidebar.classList.toggle("act");
  });
  $("li").on("click", function () {
    $(this).siblings().removeClass("akt");
    $(this).addClass("akt");
    console.log("ok");
  });
});
$(".btn-tanya").on("click", function () {
  window.location = "chat.html?room=rqwew&name=student";
});
$(".btn-tanya2").on("click", function () {
  window.location = "chat.html?room=rqwew&name=student";
});

$(document).ready(function () {
  $.ajax({
    url: "https://sixmath.vnnyx.my.id/api/user/teacher/all",
    type: "post",
    dataType: "json",
    headers: {
      Authorization: "Bearer " + localStorage.getItem("access_token"),
    },
    data: {
      username: "username",
      avatar: "avatar",
    },
    success: function (result) {
      let datas = result.data;
      console.log(datas);
      $(".box2").append(
        `
                <img src="../assets/logo-2.png" alt="" class="img2">
                <img src="../assets/` +
          datas.avatar +
          `" alt="" class="img1">
                <h1 class="nama">` +
          datas.username +
          `</h1>
                <p class="umur">24 Tahun</p>
                <button class="btn-tanya2">Tanya Guru</button>
            `
      );
    },
    failure: function () {
      console.log("ok");
    },
  });
});
