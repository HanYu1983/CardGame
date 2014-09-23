window.app = window.app || {};
window.app.view = window.app.view || {};
window.app.view.DeckListView = DeckListView;

var EventDispatcher = vic.events.EventDispatcher;
var Event = vic.events.Event;

function DeckListView( options ){
	EventDispatcher.call( this );
	this._root = options.root;
}

DeckListView.prototype = {
	__proto__:EventDispatcher.prototype,
	addCard:function( cmp ){
		var itemroot = $('<div class="deckItem"></div>' );
		itemroot.attr( 'id', 'deckItem_' + cmp.getId() );
		
		var div_name = $('<div id="name" class="name floatLeft"></div>' );
		div_name.html( cmp.getModel().Name + ':' + cmp.getModel().Action );
		
		var spr_count = $( '<input class="spr_count" value="3" />' );
		spr_count.attr( 'id', 'spr_count_' + cmp.getId() );
		
		var btn_delete = $( '<button id="delete">delete</button>' );
		btn_delete.attr( 'id', 'btn_delete_' + cmp.getId() );
		
		itemroot.append( div_name );
		itemroot.append( spr_count );
		itemroot.append( btn_delete );
		
		this._root.append( itemroot );
		
		spr_count.spinner( {min:0} );
		spr_count.on( "spinchange", function( event, ui ) {
			var target = $(this);
			var cid = target.attr('id');
			var val = target.val();
			var id = cid.substring( 'spr_count_'.length, cid.length );
			self.dispatchEvent( new Event( 'onDeckListViewSprCountChange', {index:id, val:val} ));
		} );
		var self = this;
		btn_delete.click( function(){
			var target = $(this);
			var id = target.attr('id');
			var cid = id.substring( 'btn_delete_'.length, id.length );
			self.dispatchEvent( new Event( 'onDeckListViewBtnDeleteClick', {index:cid} ));
		});
	},
	subCard:function( cmp ){
		var targetname = '#deckItem_' + cmp.getId();
		this._root.find( targetname ).remove();
	}
}