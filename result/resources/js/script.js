$(function () {
    // ユーザー作成確認画面での、パスワード表示
    $("#togglePassword").click(function () {
        let passwordForm = $("#showPassword");
        if (passwordForm.attr("type") == "password") {
            passwordForm.attr("type", "text");
            $(this).text("パスワードを非表示")
        } else {
            passwordForm.attr("type", "password");
            $(this).text("パスワードを表示")
        }
    });

    // ユーザー作成時に、ID重複を確認する
    $("#inputID").blur(function () {
        $.ajax({
            url: '/checkid',
            type: 'POST',
            data:{
                'userid':$(this).val(),
            }
        })
        .done(function(data, textStatus, jqXHR){  //data, success, 200
            if(Number(data)==0){
                $("#inputIDvalid").text("OK");
            } else {
                $("#inputIDvalid").text("すでにこのIDは使われています");
            }
        });
    });
});