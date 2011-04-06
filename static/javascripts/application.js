// ====== Copy address blocks in group, Account and Location pages ==========
// AS dynamically builds form names based on the id of the record being
// edited. New records get a nil value in the name.
// Determine the record id and create the var names for the copy.
// Buggy for Ver 1.0 release - doesn't handle multiple instances of 
// create/update pages open at once. Namespace issues.

var ids;
var record_id;
var group;
var baseFields = $A(['contact_name', 'contact_title', 'contact_phone', 'contact_email',
    'address1', 'address2', 'city', 'state', 'zip', 'city_state_postal_code'])

getVars = function() {
    // Each form has a name field that we can pull the record id from.
    record_id = "_" + $A(document.forms[0].elements).detect(function(item){return item.id.include('record_name_');}).id.split("_").slice(2).join('_');
}

copyFields = function(from, to) {
    getVars();
    baseFields.each(function(field) {
        var toField   = $("record_" + to + "_" + field + record_id);
        var fromField = $("record_" + from + "_" + field + record_id);
        // Check the fields actually exist on the form - Display contacts are never present.
        if (toField != null ) {
            toField.value = fromField.value; // Do the copy.
        }
    });
    // Country seems to be special.
    $("record_" + to + "_country").value = $("record_" + from + "_country").value;

}  

// ========== Misc functions used by control files =============================

// Walk down the sortable_products object, select children noting the order and
// selecting those checked. Stuff results into record[products].
serializeControlFileProducts = function(event) {
    // Initialize products with beginning yaml.
    var products      = event.findElement().up('form').down("#products");
    if (!products){alert("missing hidden products element. please alert clarity IT.");}
    var controlFileUl = event.findElement().up('form').down("#sortable_products");
    if (!controlFileUl){alert("missing sortable_products ul. please alert clarity IT.");}
    products.value = "--- \n";
    controlFileUl.childElements().each(function(li) {
        if (li.down().checked) {
            products.value += "- " + li.down().value + "\n";
        }
    });
    // If products is empty, replace with empty array yaml.
    if (products.value == "--- \n") {
        products.value = "--- []\n\n";
    }
    return products.value;
};
// Need to extract data out of the Sortable items so they can be posted on submit.
serializeFilterFields = function(event) {
    var target = $A();
    var filters = event.findElement().up('form').down("#filters");
    var controlFilterDiv = event.findElement().up('form').down('#sortable_filters');
    if (controlFilterDiv){
      controlFilterDiv.select('div').each(function(row) {
          var rowNumber = row.id.split("_")[1];
          target.push([rowNumber,
                       row.down('#filter_' + rowNumber).value,
                       row.down('#new_filter_' + rowNumber).value,
                       row.down('#description_' + rowNumber).value]);
      });
      filters.setValue(target.toJSON());
    }
};

validateGroupAccountLocation = function(event){
  var focus_id = null;
  var message = '';
  var ids = new Array(new Array('#control_file_group_id','You must select a Group'),
                      new Array('#control_file_account_id','You must select an Account.'),
                      new Array('#control_file_location_id','You must select a specific Location or all locations for account level.'));
  var settings = event.findElement().up('form').down("#control_file_display_form #settings");
  if (!settings){alert("missing control file settings div. please alert clarity IT.");}
  ids.each(function(id){
    if (settings.down(id[0]).value == ''){
      focus_id = id[0];
      message  = id[1];
    }
  });
  if (message.length > 0){
    alert(message);
    Form.Element.focus(settings.down(ids[0][0]));
    Form.Element.focus(settings.down(focus_id));
    event.stop();
  }
};


// Add banding to filter rows.
function bandedFilters(id) {
    var element = $(id);
  if (!element){alert("missing bandedFilters "+id+" div. please alert clarity IT.");}
    element.select('div').each( function(item, index){
        if (index%2 == 1) {
            item.setStyle('background-color: #f6f6f6; border: 1px solid #d0d0d0;');
        } else {
            item.setStyle('background-color: #e5e5e5; border: 1px solid #d0d0d0;');
        }
    });
}

// Remove a new filter row.
removeFilter = function(cross) {
    // Find the parent div tag this cross is part of. And remove it.
    cross.up('div').remove();
}

// Construct URL query string for checking filter (id plus filter value)
getFilter = function(checker) {
    var id = checker.up('div').id.split('_')[1]
    var qString = "id=" + id
    qString += "&filter=" + encodeURIComponent($('filter_' + id).value)
    return qString
}

// Class that auto-resizes a text area - used for entering multi-line filters.
// A textarea with a dynamic id uses this class like this:
//   new ResizingTextArea($("filter_<%= row %>"));
//   or
//   new ResizingTextArea($("filter_102"));
var ResizingTextArea = Class.create();

ResizingTextArea.prototype = {
    defaultRows: 1,

    initialize: function(field)
    {
        this.defaultRows = Math.max(field.rows, 1);
        this.resizeNeeded = this.resizeNeeded.bindAsEventListener(this);
        Event.observe(field, "click", this.resizeNeeded);
        Event.observe(field, "keyup", this.resizeNeeded);
    },

    resizeNeeded: function(event)
    {
        var t = Event.element(event);
        var lines = t.value.split('\n');
        var newRows = lines.length;
        var oldRows = t.rows;
        for (var i = 0; i < lines.length; i++)
        {
            var line = lines[i];
            if (line.length >= t.cols) newRows += Math.floor(line.length / t.cols);
        }
        if (newRows > t.rows) {
            t.rows = newRows;
        //            t.style.height = (15 * newRows) + "px";

        }
        if (newRows < t.rows) {
            t.rows = Math.max(this.defaultRows, newRows);
        //            t.style.height = (15 * Math.max(this.defaultRows, newRows)) + "px";
        }
    }
}

// See how many filters we already have and return the next filter row number.
// Ignore any _existing_ filter rows - those with id > 99.
nextRow = function(id) {
    var element = $(id);
    //return element.select('div').length + 1;

    if ($$('#sortable_filters div').length == 0) {
        return 1;
    }
    else {
        return $$('#sortable_filters div').max(function(row){
            var rowVal=1;
            if (row.id.split('_')[1] < 99) {
                rowVal = row.id.split('_')[1]
            }
            return (parseInt(rowVal) + 1)
        })
    }
}

// ========== Love this one ====================================================
copyToClipboard = function(text) {
    if (window.clipboardData)
    {
        window.clipboardData.setData('text',text);
    }
    else
    //    Commented out works for Opera and Safari but only with Flash version < 10. Let's focus on FF.
    //    And for Flash 10 use: http://code.google.com/p/zeroclipboard/
    //        {
    //      var clipboarddiv = document.getElementById('divclipboardswf');
    //      if (clipboarddiv == null)
    //      {
    //         clipboarddiv = document.createElement('div');
    //         clipboarddiv.setAttribute("name", "divclipboardswf");
    //         clipboarddiv.setAttribute("id", "divclipboardswf");
    //         document.body.appendChild(clipboarddiv);
    //      }
    //      clipboarddiv.innerHTML = '<embed src="/images/clipboard.swf" FlashVars="clipboard=' +
    //        encodeURIComponent(text)+ '" width="0" height="0" type="application/x-shockwave-flash"></embed>';
    //    }
    {
        // Firefox browser only. Opera and Safari are untested.
        //  Requires user to set permissions locally.
        // Open new tab. Go to about:config
        // Right click anywhere and select New -> Boolean
        // Enter: signed.applets.codebase_principal_support
        // Ensure its value is set to true.
        // Now FF will display a dialog box for security warnings. You have to accept each time it comes up.
        // LDB: commented out below IE does not like this and not sure what is it for?
        //      am sure will find out right after we release to production
//        netscape.security.PrivilegeManager.enablePrivilege('UniversalXPConnect');
//        const gClipboardHelper = Components.classes["@mozilla.org/widget/clipboardhelper;1"].
//        getService(Components.interfaces.nsIClipboardHelper);
//        gClipboardHelper.copyString(text);
    }
    // alert('Text copied to your clipboard: ' + text);
    return false;
}

// If the sortable object is empty, add the 'empty' class to display a blank target.
checkEmpty = function(list) {
    methodStart = list.down('li') ? 'remove' : 'add';
    list[methodStart + 'ClassName']('empty');
}

// =============================================================================

disableDateRange = function() {
// Start with the date range selector disabled.
  $('spreadsheet_start').disable();
  $('spreadsheet_stop').disable();
}
