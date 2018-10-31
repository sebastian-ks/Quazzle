function setActive(elem) {
  allOff();
  if (!elem.classList.contains("active")) {
    elem.classList.add("active");
  }
}

function allOff() {
  var anchors = document.getElementsByClassName("chat-item");
  for (var i = 0; i < anchors.length; i++) {
    var anchor = anchors[i];
    if (anchor.classList.contains("active")) {
      anchor.classList.remove("active");
    }
  }
}

window.onload = function() {
  var items = document.getElementsByClassName("chat-item");
  for (var i = 0; i < items.length; i++) (function(i) {
    items[i].onclick = function() {
      setActive(items[i]);
    }
  })(i);
}
