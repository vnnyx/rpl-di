window.addEventListener("DOMContentLoaded", (_) => {
    var roomname = getParameterByName("room");
    var username = getParameterByName("name");
    let websocket = new WebSocket(
      "ws://3e69-103-57-36-42.ngrok.io/websocket?name=" + username
    );
    let room = document.getElementById("sender-chat");
  
    websocket.addEventListener("open", (e) => {
      websocket.send(JSON.stringify({ message: roomname, action: "join-room" }));
  
      let form = document.getElementById("input-form");
      form.addEventListener("submit", function (event) {
        event.preventDefault();
        let text = document.getElementById("input-text");
        check = text.value;
  
        websocket.send(
          JSON.stringify({
            action: "send-message",
            message: text.value,
            target: roomname,
            sender: {
              name: username,
            },
          })
        );
        text.value = "";
      });
    });
  
    websocket.addEventListener("message", function (e) {
      let data = JSON.parse(e.data);
  
      let justify = "justify-content-end";
      let clasz = "'talk-bubble talk-bubble-sender tri-right round right-in '";
      if (data.sender.name == "teacher") {
        justify = "justify-content-start";
        clasz = "'talk-bubble talk-bubble-recipient tri-left round left-in '";
      }
  
      let chatContent = `<div class='d-flex ${justify}' id='sender-chat'>
      <div class=${clasz}>
          <div class='talktext text-white'>
              <p class='mb-2'>
                  ${data.message}
              </p>
              <p class='text-end mx-1 fw-bold'>
                  ${data.time}
              </p>
          </div>
      </div>
  </div>`;
      room.innerHTML += chatContent;
      room.scrollTop = room.scrollHeight; // Auto scroll to the bottom
    });
  });
  
  function getParameterByName(name, url = window.location.href) {
    name = name.replace(/[\[\]]/g, "\\$&");
    var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)"),
      results = regex.exec(url);
    if (!results) return null;
    if (!results[2]) return "";
    return decodeURIComponent(results[2].replace(/\+/g, " "));
  }