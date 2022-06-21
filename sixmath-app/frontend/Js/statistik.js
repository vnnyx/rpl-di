


$(document).ready(function () {
    var data
    var config
    var myChart
    $.ajax({
        url: 'https://sixmath.vnnyx.my.id/api/user/registration-statistics',
        type: 'GET',
        contentType: 'json',
        headers: {
            'Authorization': "Bearer " + localStorage.getItem('access_token')
        },
        success: function (result) {
            let datas = result.data;
            let arr = []
            let arrMonth = []
            for (i = 0; i < datas.length; i++) {
                arr.push(datas[i].amount)
                arrMonth.push(datas[i].month)
            }
            const data = {
                labels: arrMonth.reverse(),
                datasets: [{
                    data: arr.reverse(),
                    backgroundColor: ['#6A35FF', '#EBC314'],
                    borderRadius: 8,
                    borderSkipped: false,
                    barPercentage: 0.7
                }]
            };
            const config = {
                type: 'bar',
                data: data,
                options: {
                    plugins: {
                        legend: {
                            display: false
                        }
                    },
                    scales: {
                        y: {
                            ticks: {
                                callback: (value, index, values)=> {
                                    return value;
                                }
                            },
                            beginAtZero: true,
                            grid: {
                                display: false,
                                drawOnChartArea: false,
                                drawBorder: false
                            }
                        },
                        x: {
                            grid: {
                                display: false,
                                drawOnChartArea: false,
                                drawBorder: false
                            }
                        }
                    }
                },
            };
            myChart = new Chart(
                document.getElementById('myChart'),
                config
            );
            console.log(datas)
        }
    })
    $("select.combo-box").change(function () {
        myChart.destroy()
        var bar = $(this).children("option:selected").val();
        $.ajax({
            url: "https://sixmath.vnnyx.my.id/api/user/registration-statistics?month=" + bar,
            type: 'GET',
            contentType: 'json',
            headers: {
                'Authorization': "Bearer " + localStorage.getItem('access_token')
            },
            success: function (result) {
                let datas = result.data;
                let arr = []
                let arrMonth = []
                for (i = 0; i < datas.length; i++) {
                    arr.push(datas[i].amount)
                    arrMonth.push(datas[i].month)
                }
                const labels = Samples.utils.months({ count: datas.length });
                const data = {
                    labels: arrMonth.reverse(),
                    datasets: [{
                        data: arr.reverse(),
                        backgroundColor: ['#6A35FF', '#EBC314'],
                        borderRadius: 8,
                        borderSkipped: false,
                        barPercentage: 0.7
                    }]
                };
                const config = {
                    type: 'bar',
                    data: data,
                    options: {
                        plugins: {
                            legend: {
                                display: false
                            }
                        },
                        scales: {
                            y: {
                                beginAtZero: true,
                                grid: {
                                    display: false,
                                    drawOnChartArea: false,
                                    drawBorder: false
                                },
                                type: 'linear',
                            },
                            x: {
                                grid: {
                                    display: false,
                                    drawOnChartArea: false,
                                    drawBorder: false
                                }
                            }
                        }
                    },
                };
                myChart = new Chart(
                    document.getElementById('myChart'),
                    config
                );
            }
        })

    });
    const dough = {
        labels: [
            'Matematika',
            'Penalaran'
        ],
        datasets: [{
            data: [300, 250],
            backgroundColor: ['#6A35FF', '#EBC314'],
            hoverOffset: 4
        }]
    };
    const cfg = {
        type: 'doughnut',
        data: dough,
        options: {
            plugins: {
                legend: {
                    position: 'bottom',
                    labels: {
                        boxHeight: 20,
                        boxWidth: 20
                    }
                }
            }
        }
    };

    var doughChart = new Chart(
        document.getElementById('doughChart'),
        cfg
    );
})