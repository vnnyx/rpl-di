$(document).ready(function () {
    let sidebar = document.querySelector(".sidebar")
    $('#btn').on('click', function () {
        sidebar.classList.toggle("act");
    });
    $('li').on('click', function () {
        $(this).siblings().removeClass('akt')
        $(this).addClass('akt')
        console.log('ok')
    });
    $('.btn-mengerjakan').on('click', function () {
        window.location = 'pelajaran.html'
    })
});