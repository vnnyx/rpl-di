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
    url: 'https://sixmath.vnnyx.my.id/api/video/detail/1',
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
        <h6 class="desc mb-3 fw-bold text-muted">Deskripsi</h6>
        <p class="deskripsi ms-4 mb-5 px-5" style="text-indent: 30px; width: 100%;">`+ datas.description + `</p>
    `)
    }
})

$.ajax({
    url: 'https://sixmath.vnnyx.my.id/api/video/recommended?main=1',
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
                    <div class="card">
                        <iframe class="iframe-rek"
                            src="https://www.youtube.com/embed/`+ data.url_video + `" title="YouTube video player"
                            frameborder="0"
                            allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                            allowfullscreen></iframe>
                        <div class="card-body">
                            <p class="card-text">`+ data.title +`</p>
                        </div>
                    </div>
                </div>
            </div>
        `)
        })
    }
})