var sock = new WebSocket("ws://localhost:4000/");
sock.onmessage = function(m) { console.log("Received:", m.data); }
sock.send("Hello!\n")
