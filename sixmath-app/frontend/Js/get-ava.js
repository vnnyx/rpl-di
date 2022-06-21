$.ajax({
    url: 'https://sixmath.vnnyx.my.id/api/auth/login',
    type: 'post',
    dataType: 'json',
    contentType: "application/json",
    data: {
        'avatar' : 'avatar'
    },
    success: function (result) {
        localStorage.setItem(
            'access_token', result.data.access_token
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