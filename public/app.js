$(function () {
    let websocket = new WebSocket("ws://localhost:8080/ws");
    let chatBox = $("#chat-box");
    websocket.addEventListener("message", function (e) {
        let data = JSON.parse(e.data);
        let chatContent = `<p><strong>${data.username}</strong>: ${data.message}</p>`;
        chatBox.append(chatContent);
        window.scrollTo(0, document.body.scrollHeight)// Auto scroll to the bottom
    });
    $("#input-form").on("submit", function (event) {
        event.preventDefault();
        let username = $("#input-username")[0].value;
        let message = $("#input-message")[0].value;
        websocket.send(
            JSON.stringify({
                username: username,
                message: message,
            })
        );
        $("#input-message")[0].value = "";
    });
  });