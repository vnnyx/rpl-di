$(document).ready(function () {
    let sidebar = document.querySelector(".sidebar")
    $('#btn').on('click', function () {
        sidebar.classList.toggle("act");
    });
    $('li').on('click', function () {
        $(this).siblings().removeClass('akt')
        $(this).addClass('akt')
    });
});
$.ajax({
    url: 'https://sixmath.vnnyx.my.id/api/video',
    type: 'GET',
    contentType: 'json',
    headers: {
        'Authorization': "Bearer " + localStorage.getItem('access_token')
    },
    success: function (result) {
        let datas = result.data;
        $('#main-video').append(`
        <iframe width="591" height="328" src="https://www.youtube.com/embed/`+ datas.url_video + `" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
    `)
        $('#detail').append(`
        <h1 class="judul-main fw-bold mb-4" style="margin-left: 75px;">`+ datas.title + `</h1>
        <h6 class="ms-4 mb-3 fw-bold text-muted">Deskripsi</h6>
        <p class="deskripsi ms-4 mb-5" style="text-indent: 30px; width: 96%;">`+ datas.deskripsi + `</p>
    `)
    }
})

$.ajax({
    url: 'https://sixmath.vnnyx.my.id/api/video/recommended',
    type: 'GET',
    contentType: 'json',
    headers: {
        'Authorization': "Bearer " + localStorage.getItem('access_token')
    },
    success: function (result) {
        let datas = result.data.rows;
        $.each(datas, function (i, data) {
            $('.video-rekomendasi').append(`
            <div class="row">
                <div class="col mb-4 ms-3 me-3">
                    <div class="card" style="width: 14rem; height: 14rem;">
                        <iframe width="223" height="" style="border-radius: 4px;"
                            src="https://www.youtube.com/embed/`+ data.url_video + `" title="YouTube video player"
                            frameborder="0"
                            allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                            allowfullscreen></iframe>
                        <div class="card-body">
                            <p class="card-text">Judul</p>
                        </div>
                    </div>
                </div>
            </div>
        `)
        })
    }
})