$ = jQuery;
$(document).ready(function(){
  if($.browser.msie && (parseInt($.browser.version) < 8)) {
  	$('body').prepend('<div id="browserWarning"><a href="#">Some features may not work until you upgrade to a supported browser.</a></div>');
  	$('#browserWarning a').click(function(){
  		alert("We noticed you're using an older version of Internet Explorer. Unfortunately, we only support Internet Explorer 8 or later. You can upgrade for free at: http://microsoft.com/ie");
  	});
  }
});
