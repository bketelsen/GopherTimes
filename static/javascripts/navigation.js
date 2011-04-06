/****************************************/
function pointWithinBounds(point, bounds) {
    if (point.x < bounds.left || point.x > bounds.right) return false;
    if (point.y < bounds.top || point.y > bounds.bottom) return false;
    return true;
}


/****************************************/
jQuery.noConflict();
jQuery(document).ready(function($) {
    /****************************************/
    // centre the popover
    $('.nav > li > ul.popover').each(function(index, popover) {
        popover = $(popover);
        menuItem = popover.parent();
        menuItemOffset = menuItem.offset();
        menuItemLeft = menuItemOffset.left;
        menuItemWidth = menuItem.width();
        popoverWidth = popover.width() + 15;
        popoverLeft = menuItemLeft + (menuItemWidth / 2) - (popoverWidth / 2);
		//console.log('left:', menuItemOffset.left, 'top:', menuItemOffset.top);
        popover.offset({
            left: popoverLeft
        });
    });


    /****************************************/
    // popover manager
    $('.nav > li').hoverIntent(function(e) {
        var menuItem = $(e.target).parent();
        $('.popover:not(.active)').removeClass('active').slideUp().parent().removeClass('hover');
        var popover = $(menuItem.children('.popover').first());
        popover.addClass('active').slideDown();
        menuItem.addClass('hover');
        if (menuItem.hasClass('capped')) {
            var cap = menuItem.hasClass('leftcap') ? menuItem.prev() : menuItem.next();
            if (!cap.hasClass('selected')) cap.addClass('hover');
        }

        // improved handling of mouseOut events
        $('body').append('<div id="dehoveriser"></div>');
        $('#dehoveriser').hoverIntent(function(e) {
            $('.popover.active').each(function(index, item) {
                deactivatePopover($(item).parent());
            });
        },
        function() {});
    },
    function(e) {
        var bounds = $(e.target).offset();
        bounds.right = bounds.left + $(e.target).width();
        bounds.bottom = bounds.top + $(e.target).height();
        if (pointWithinBounds({
            x: e.clientX,
            y: e.clientY
        },
        bounds)) {
            return false;
        }

        deactivatePopover($(e.target).parent());
    });


    /****************************************/
    // popover deactivator
    function deactivatePopover(menuItem) {
        var popover = menuItem.children('.popover').first();
        popover.removeClass('active').slideUp().parent().removeClass('hover');
        menuItem.removeClass('hover');
        if (menuItem.hasClass('capped')) {
            var cap = menuItem.hasClass('leftcap') ? menuItem.prev() : menuItem.next();
            if (!cap.hasClass('selected')) cap.removeClass('hover');
        }
        $('#dehoveriser').remove();
    }
});