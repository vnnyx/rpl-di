$(document).ready(function () {
    $('label').on('click', function () {
        $(this).siblings().removeClass('aktif')
        $(this).addClass('aktif')
    });
    $('.btn-next').on('click', function () {
        if ($('#btnradio1').prop("checked")) {
            window.location = '../html/teacher/guru-registrasi.html'
        }else if ($('#btnradio2').prop("checked")){
            window.location = '../html/student/register.html'
        }else{
            window.location = '../ortu/parent-regist.html'
        }
    })
});