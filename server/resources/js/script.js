$(function () {
    //basedayとの日付差を求める関数
    function countDay() {
        var today = new Date();
        var baseday = new Date(2020, 5, 1);

        var day = Math.ceil((today-baseday)/(60*60*24*1000));
        return day
    };

    //いつもローカルストレージがクリアされていたらrankingを付与する
    if (localStorage.getItem("ranking") == null) {
        var ranking = {
            date: countDay(),
            age: 0,
            area: 99,
            fromCharaId: [],
            gender: 9,
            idAnswered: false,
            votingHistory: [],
            votingTodayHistory: []
        };
        localStorage.setItem("ranking", JSON.stringify(ranking));
    } else { // 日を跨いでいたら、votingTodayHistoryをクリア
        var ranking = JSON.parse(localStorage.getItem("ranking"));

        if(ranking.date != countDay()){
            ranking.date = countDay();
            ranking.votingTodayHistory = [];
            localStorage.setItem("ranking", JSON.stringify(ranking));
        }
    }

    //characterの投票ボタンを履歴に沿って無効化
    $(".btn-vote").each(function(i, obj) {
        var votingtodayHistory = JSON.parse(localStorage.getItem("ranking")).votingTodayHistory;
        if(votingtodayHistory.indexOf($(obj).val()) >= 0){
            $(obj).prop("disabled", true);
        }
    });

    // characterページの投票ボタン
    $(".btn-vote").on('click', function() {
        var ranking = JSON.parse(localStorage.getItem("ranking"));
        
        var character = $(this).val();
        ranking.fromCharaId = [character];
        var url = "/characters/"+character+"/vote"
        var isAnswered = ranking.idAnswered;
        if (isAnswered) {
            ranking.votingTodayHistory.push(character);
            ranking.votingHistory.push(character);
        }

        localStorage.setItem("ranking", JSON.stringify(ranking));

        var form = `<form method='post' action='${url}' id='refresh' style='display: none;'>
                        <input type='hidden' name='character' value='${character}'>
                    </form>`;
        $("body").append(form);
        $("#refresh").submit();
    });

    // フォームのsubmit時にローカルストレージに保存
    $("#questionnaire").submit(function() {
        var ranking = JSON.parse(localStorage.getItem("ranking"));
        var character = ranking.fromCharaId[0];

        var form = `<input type='hidden' name='character' value='${character}'>`;
        $(this).append(form);
        
        ranking.idAnswered = true;
        ranking.age = Number($("#form-age").val());
        ranking.gender = Number($("#form-gender").val());
        ranking.area = Number($("#form-address").val());
        ranking.votingTodayHistory.push(character);
        ranking.votingHistory.push(character);
        localStorage.setItem("ranking", JSON.stringify(ranking));
    });
});