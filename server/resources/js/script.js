$(function () {
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
    
    console.log(JSON.parse(localStorage.getItem("ranking")));

    $(".btn-vote").on('click', function() {
        console.log("投票ボタンだよ");
    });
});