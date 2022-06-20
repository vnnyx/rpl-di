$(document).ready(function () {
    $('label').on('click', function () {
        $(this).siblings().removeClass('aktif')
        $(this).addClass('aktif')
    });
    $('.btn-next').on('click', function () {
        if ($('#btnradio1').prop("checked")) {
            window.location = '/frontend/html/student/register.html'
        }else if ($('#btnradio2').prop("checked")){
            window.location = '/frontend/html/student/login.html'
        }else{
            window.location = '/frontend/html/student/statistik.html'
        }
    })
});