$(function () {
    //いつもローカルストレージの中身を表示
    console.log(JSON.parse(localStorage.getItem("ranking")));

    //いつもローカルストレージがクリアされていたらrankingを付与する
    if (localStorage.getItem("ranking") == null) {
        var ranking = {
            age: 0,
            area: 99,
            fromCharaId: [],
            gender: 9,
            idAnswered: false,
            user: "",
            votingHistory: [],
            votingTodayHistory: []
        };
        localStorage.setItem("ranking", JSON.stringify(ranking));
    };

    // characterページの投票ボタン
    $(".btn-vote").on('click', function() {
        console.log("投票ボタンだよ");
    });

    // フォームのsubmit時にローカルストレージに保存
    $("#questionnaire").submit(function() {
        var ranking = JSON.parse(localStorage.getItem("ranking"));
        ranking.idAnswered = true;
        ranking.age = Number($("#form-age").val());
        ranking.gender = Number($("#form-gender").val());
        ranking.area = Number($("#form-address").val());
        localStorage.setItem("ranking", JSON.stringify(ranking));
    });
});