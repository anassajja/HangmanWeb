var countdownNumberEl = document.getElementById('countdown-number');
var countdown =0;
var second = "s"
countdownNumberEl.textContent = countdown;
setInterval(function() {
  countdown = ++countdown <= 0 ? 60 : countdown;
  countdownNumberEl.textContent = countdown;
}, 1000);
