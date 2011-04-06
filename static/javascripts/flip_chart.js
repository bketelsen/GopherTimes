$ = jQuery;
$(document).ready(function(){
  // update real-time stats every 15 seconds
  $('body').everyTime(15*1000, 'poll_real_time_stats', function(){
    $.get('/marketing/refresh_real_time_stats', function(data, status, xhr){
      $('embed.flip_chart_tick').remove();
      $('#real_time_stats').fadeOut('fast').html(data).fadeIn('fast');
    });
  });
  
  // update latest denial reasons every two minutes
  $('body').everyTime(160*1000, 'poll_denial_stats', function(){
    var redis_key = 'marketing_stats_latest_denial_reasons';
    var stats_uri = '/marketing/refresh_denial_reasons?redis_key=' + redis_key;
    $('#'+redis_key+'_marquee').load(stats_uri);
  });
});

// animate the transition of a number div to its new value
function setStripPosition(divID, digit) {
  var targetPos = (digit % 10) * -2640;
  var stripDiv = $('#'+divID);
  stripDiv.css('top', targetPos-2640).animate({
    'top': ((digit == 0) ? 0 : targetPos)
  }, 'slow');
}
