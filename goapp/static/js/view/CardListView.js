window.app = window.app || {};
window.app.view = window.app.view || {};
window.app.view.CardListView = CardListView;

var EventDispatcher = vic.events.EventDispatcher;
var Event = vic.events.Event;

function CardListView( options ){
	EventDispatcher.call( this );
	this._root = options.root;
}

CardListView.prototype = {
	__proto__:EventDispatcher.prototype,
	addCard:function( index ){
		var itemroot = $('<div class="floatLeft item"></div>' );
		var cardroot = $('<div class="card"></div>' );
		var controlroot = $('<div class="itemControlContainer"></div>' );
		var checkbox = $('<input type="checkbox" id="chk_want_' + index + '" class="chk_want"/><label for="check">need</label>' );
		controlroot.append( checkbox);
		itemroot.append( cardroot );
		itemroot.append( controlroot );
		
		var self = this;
		checkbox.change( function(){
			var target = $(this);
			var id = target.attr( 'id' );
			var cid = id.substring( 'chk_want_'.length, id.length );
			var val = target[0].checked;
			self.dispatchEvent( new Event( 'onCardListViewChkNeedChange', {cid:cid, val:val } ));
		});
		this._root.append( itemroot );
	},
	setTargetNeed:function( id, val ){
		var target = this._root.find( '#chk_want_' + id );
		target.attr( 'checked', val );
	}
}