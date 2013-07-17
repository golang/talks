// +build OMIT

package main

import "html/template"
import "net/http"

func rootHandler(w http.ResponseWriter, r *http.Request) {
	rootTemplate.Execute(w, listenAddr)
}

var rootTemplate = template.Must(template.New("root").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8" />
<script>

var input, output, websocket;

function showMessage(m) {
	var p = document.createElement("p");
	p.innerHTML = m;
	output.appendChild(p);
}

function onMessage(e) {
	showMessage(e.data);
}

function onClose() {
	showMessage("Connection closed.");
}

function sendMessage() {
	var m = input.value;
	input.value = "";
	websocket.send(m + "\n");
	showMessage(m);
}

function onKey(e) {
	if (e.keyCode == 13) {
		sendMessage();
	}
}

function init() {
	input = document.getElementById("input");
	input.addEventListener("keyup", onKey, false);

	output = document.getElementById("output");

	websocket = new WebSocket("ws://{{.}}/socket");
	websocket.onmessage = onMessage;
	websocket.onclose = onClose;
}

window.addEventListener("load", init, false);

</script>
</head>
<body>
<input id="input" type="text">
<div id="output"></div>
</body>
</html>
`))
