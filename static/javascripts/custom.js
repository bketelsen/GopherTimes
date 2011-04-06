$(document).ready(function() {
	
	var clamp_b = 0;

	var $bigone = $('#zoom-photo .show img');
	
	var $smallone = $('#zoom-photo .thumb .map');
	
	var $thumb = $('#zoom-photo .thumb');
	
	var $scale = $('input[name=scale]');
	
	var bigone = {};
	bigone.width = 180;
	bigone.height = 200;
	var smallone = {};
	smallone.width = 45;
	smallone.height = 50;
	
	var domleft = domtop = 0;
	
	// init
	$scale.val(1);
	// add alpha value
	$('#zoom-photo div.map').css({'opacity': '0.3', 'filter': 'Alpha(opacity=30)'});
	
	$('#zoom-photo .thumb .map').draggable({
		'containment': 'parent',
		'start': function() {
			clamp_b = 4 * $scale.val();
		},
		'drag': function(e, ui) {
			domleft = ui.position.left * clamp_b;
			if ( ($bigone.width() - domleft) < bigone.width ) domleft = $bigone.width() - bigone.width;
			domtop = ui.position.top * clamp_b;
			if ( ($bigone.height() - domtop) < bigone.height ) domtop = $bigone.height() - bigone.height;
			$bigone.css({'left': '-' + domleft + 'px', 'top': '-' + domtop + 'px'});
		}
	});
	
	$thumb.click(function(e) {
		
		var _this = $(this);
		
		var client = {};
		
		client.top = e.clientY - _this.get(0).offsetTop - _this.get(0).parentNode.offsetTop;
		
		client.left = e.clientX - _this.get(0).offsetLeft - _this.get(0).parentNode.offsetLeft;
		
		setCenter(client);
		
	});
	
	var setCenter = function(client) {
		
		if ( client.left + $smallone.width() / 2 > $thumb.width() ) client.left = $thumb.width() - $smallone.width() / 2;
		if ( client.left - $smallone.width() / 2 < 0 ) client.left = $smallone.width() / 2;
		if ( client.top + $smallone.height() / 2 > $thumb.height() ) client.top = $thumb.height() - $smallone.height() / 2;
		if ( client.top - $smallone.height() / 2 < 0 ) client.top = $smallone.height() / 2;
		
		var domleft = client.left - ($smallone.width() / 2);
		
		var domtop = client.top - ($smallone.height() / 2);
		
		$smallone.css({'left': domleft + 'px', 'top': domtop + 'px'});
		
		clamp_b = 4 * $scale.val();
		
		domleft = domleft * clamp_b;
		if ( ($bigone.width() - domleft) < bigone.width ) domleft = $bigone.width() - bigone.width;
		domtop = domtop * clamp_b;
		if ( ($bigone.height() - domtop) < bigone.height ) domtop = $bigone.height() - bigone.height;
		$bigone.css({'left': '-' + domleft + 'px', 'top': '-' + domtop + 'px' });
		
	};
	
	var getCenter = function(dom) {
		
		var client = {};
		
		client.left = (parseInt(dom.css('left')) ? parseInt(dom.css('left')) : 0) + dom.width() / 2;
		
		client.top = (parseInt(dom.css('top')) ? parseInt(dom.css('top')) : 0) + dom.height() / 2;
		
		return client;
		
	};
	
	$('#zoom-photo p.zoom-in a').click(function() {
		
		$(this).parent().hide();
		
		$thumb.show();
		
		$('#zoom-photo p.zoom-out').show();
		
	});
	
	$('#zoom-photo a.restore').click(function() {
		
		$(this).parent().hide();
		
		$thumb.hide();
		
		$('#zoom-photo p.zoom-in').show();
		
		$scale.val('1');
		
		$bigone.css({'left': '0', 'top': '0', 'width': bigone.width + 'px', 'height': bigone.height + 'px'});
		
		$smallone.css({'left': '0', 'top': '0', 'width': smallone.width + 'px', 'height': smallone.height + 'px'})
		
	});
	
	$('#zoom-photo a.plus').click(function() {
		
		if ( parseInt($scale.val()) + 1 > 8 ) return false;
		
		client = getCenter($smallone);
		
		$scale.val( parseInt($scale.val()) + 1 );
		
		$bigone.css({'width': bigone.width * $scale.val() + 'px', 'height': bigone.height * $scale.val() + 'px' });
		
		$smallone.css({'width': smallone.width / $scale.val() + 'px', 'height': smallone.height / $scale.val() + 'px'});
		
		setCenter(client);
		
	});
	
	$('#zoom-photo a.desc').click(function() {
		
		if ( parseInt($scale.val()) - 1 < 1 ) return false;
		
		client = getCenter($smallone);
		
		$scale.val( parseInt($scale.val()) - 1 );
		
		$bigone.css({'width': bigone.width * $scale.val() + 'px', 'height': bigone.height * $scale.val() + 'px' });
		
		$smallone.css({'width': smallone.width / $scale.val() + 'px', 'height': smallone.height / $scale.val() + 'px'});
		
		setCenter(client);
		
	});
	
});
