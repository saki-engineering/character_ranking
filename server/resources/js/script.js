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
        console.log($(this).val());
        var ranking = JSON.parse(localStorage.getItem("ranking"));
        
        var character = $(this).val();
        ranking.fromCharaId = [character];
        localStorage.setItem("ranking", JSON.stringify(ranking));

        var url = "/characters/"+character+"/vote"
        var isAnswered = ranking.idAnswered;
        var form = `<form method='post' action='${url}' id='refresh' style='display: none;'>
                        <input type='hidden' name='character' value='${character}'>
                        <input type='hidden' name='isAnswered' value='${isAnswered}'>
                    </form>`;
        $("body").append(form);
        $("#refresh").submit();
    });

    // フォームのsubmit時にローカルストレージに保存
    $("#questionnaire").submit(function() {
        var ranking = JSON.parse(localStorage.getItem("ranking"));

        var form = `<input type='hidden' name='character' value='${ranking.fromCharaId[0]}'>`;
        $(this).append(form);
        
        ranking.idAnswered = true;
        ranking.age = Number($("#form-age").val());
        ranking.gender = Number($("#form-gender").val());
        ranking.area = Number($("#form-address").val());
        localStorage.setItem("ranking", JSON.stringify(ranking));
    });
});