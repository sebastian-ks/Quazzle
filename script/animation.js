$(document).ready(function(){
  $("#add").click(function(){
    $(this).animate({left: "-=800"}, 300)
    .fadeOut(300);
    $("#recbtn").fadeIn(1000);
    $("#progress").fadeIn(1000);
    $("#recbtn").one("click",(function(){
      barTime();
    }));
  });
});

function barTime() {
  var recordTime = 5 //variable length of recording time (in seconds)
  var width = 0;
  bar = document.getElementById("bar");
  bar.style.width = "0%";
  var progressing = setInterval(function(){
    if(width >= 100){
      showplay();
      clearInterval(progressing);
    } else{
      width += 0.1;
      bar.style.width = width + '%';
    }
  },recordTime);
}

function showplay(){
  $("#recbtn").hide();
  $("#playbtn").show();
  $("#send").show();
  $("#playbtn").one("click",(function(){
    barTime();
  }));
}
