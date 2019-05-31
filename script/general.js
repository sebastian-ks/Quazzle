var recordTime = 5; //global Time to record voice messages

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
    };
  })(i);
};

window.onload = function(){
  let constraint = {
    audio: true,
    video: false
  };

  navigator.mediaDevices.getUserMedia(constraint)
  .then(function(mediaStreamObj) {
    let audio = document.getElementById("recAudio");
    if("srcObject" in audio) {
      audio.srcObject = mediaStreamObj;
    } else {
      audio.src = window.URL.createObjectURL(mediaStreamObj);
    }
    let start = document.getElementById("recbtn");
    let save = document.getElementById("playAudio");
    let play = document.getElementById("playbtn");
    let mediaRecorder = new MediaRecorder(mediaStreamObj);
    let chunks = [];

    start.addEventListener("click", (ev)=>{
      mediaRecorder.start();
      setTimeout(function(){
        mediaRecorder.stop();
      },recordTime*1000);
    });

    play.addEventListener("click", (ev)=>{
      save.play();
    })

    mediaRecorder.ondataavailable = function(ev) {
      chunks.push(ev.data);
    }
    mediaRecorder.onstop = (ev)=>{
      let blob = new Blob(chunks, {'type' : 'audio/wav;'});
      chunks = [];
      let audioURL = window.URL.createObjectURL(blob);
      save.src = audioURL;
    }
  })
  .catch(function(err) {
    console.log(err.name, err.message)
  });
}
