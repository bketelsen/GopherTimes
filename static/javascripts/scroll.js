$(document).ready(function() {
	
	var $panel = $('#gallery div.panels');
	
	var $tab   = $('#gallery ul.thumb');
	
	var $panelEles = $panel.find('p');
	
	var $tabEles   = $tab.find('li');
	
	var active = $panel.find('p.active');
	
	var length = $panelEles.size();
	
	var time = 5000;
	
	var timer;
	
	var flag = true;
	
	active = $panel.find('p').index(active);
	
	$panelEles.hide();
	
	// set the active item if not exists
	if ( active == -1 ) {
		active = 0;
		$panel.find('p').eq(0).addClass('active');
	}
	$panelEles.eq(active).show();
	$tabEles.eq(active).addClass('active');
	
	var play = function(next) {
		
		var _this = $panelEles.eq(next);
		
		var index = $panelEles.index(_this);

		if ( index == active ) return false;
		
		if ( flag == false ) return false;
		
		flag = false;
		
		clearTimeout(timer);
		
		// class
		$tab.find('li.active').removeClass('active');
		$tabEles.eq(index).addClass('active');
		
		// z-index
		$panelEles.eq(active).css('z-index', '1');
		_this.css('z-index', '20');
		
		_this.fadeIn(1000, function() {
			
			_this.addClass('active');
			
			$panelEles.eq(active).hide();
			
			$panelEles.eq(active).removeClass('active');
			
			active = index;
			
			flag = true;
			
			timer = setTimeout( function() {
				next = (active + 1) < length ? (active + 1) : 0;
				play(next)
			}, time );
			
		});

	};
	
	$tab.find('a').click(function() {
		var index = $tabEles.index( $(this).parent() );
		
		play(index);
		
		return false;
	});
	
	timer = setTimeout( function() {
		next = (active + 1) < length ? (active + 1) : 0;
		play(next)
	}, time );
	

});
