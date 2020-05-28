$(function () {
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
});