// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

(function() {
  "use strict";

  var websocket, outputs = {};

  function showMessage(o, m, className) {
    var span = document.createElement("span");
    m = m.replace(/&/g, "&amp;");
    m = m.replace(/</g, "&lt;");
    span.innerHTML = m;
    span.className = className;
    o.appendChild(span);
  }

  function onMessage(e) {
    var m = JSON.parse(e.data);
    var o = outputs[m.Id];
    if (o === null) {
      return;
    }
    if (m.Kind === "stdout" || m.Kind === "stderr") {
      showMessage(o, m.Body, m.Kind);
    }
    if (m.Kind === "end") {
      var s = "Program exited";
      if (m.Body !== "") {
        s += ": " + m.Body;
      } else {
        s += ".";
      }
      s += "\n";
      showMessage(o, s, "system");
    }
  }

  function onClose() {
    window.alert('websocket connection closed');
  }

  function sendMessage(m) {
    websocket.send(JSON.stringify(m));
  }

  var count = 0;

  function getId() {
    return "code" + (count++);
  }

  function text(node) {
    var s = "";
    for (var i = 0; i < node.childNodes.length; i++) {
      var n = node.childNodes[i];
      if (n.nodeType === 1 && n.tagName === "PRE") {
        var innerText = n.innerText === undefined ? "textContent" : "innerText";
        s += n[innerText] + "\n";
        continue;
      }
      if (n.nodeType === 1 && n.tagName !== "BUTTON") {
        s += text(n);
      }
    }
    return s;
  }

  function init(code) {
    var id = getId();

    var output = document.createElement('div');
    var outpre = document.createElement('pre');
    
    function onRun() {
      outpre.innerHTML = "";
      output.style.display = "block";
      sendMessage({Id: id, Kind: "run", Body: text(code)});
    }

    function onKill() {
      sendMessage({Id: id, Kind: "kill"});
    }

    function onClose() {
      onKill();
      output.style.display = "none";
    }

    var run = document.createElement('button');
    run.contenteditable = false;
    run.innerHTML = 'Run';
    run.addEventListener("click", onRun, false);
    var run2 = document.createElement('button');
    run2.innerHTML = 'Run';
    run2.addEventListener("click", onRun, false);
    var kill = document.createElement('button');
    kill.innerHTML = 'Kill';
    kill.addEventListener("click", onKill, false);
    var close = document.createElement('button');
    close.innerHTML = 'Close';
    close.addEventListener("click", onClose, false);

    var buttons = document.createElement('div');
    buttons.classList.add('buttons');
    buttons.appendChild(run2);
    buttons.appendChild(kill);
    buttons.appendChild(close);

    output.classList.add('output');
    output.appendChild(buttons);
    output.appendChild(outpre);
    output.style.display = "none";

    code.appendChild(run);
    code.parentNode.appendChild(output);

    outputs[id] = outpre;
  }

  var play = document.querySelectorAll('div.playground');
  for (var i = 0; i < play.length; i++) {
    init(play[i]);
  }
  if (play.length > 0) {
    // TODO(adg): pass the host and port through from gopresent
    websocket = new WebSocket("ws://localhost:3999/socket");
    websocket.onmessage = onMessage;
    websocket.onclose = onClose;
  }

})();
