window.app = window.app || {};
window.app.model = window.app.model || {};
window.app.model.DeckListModel = DeckListModel;

var EventDispatcher = vic.events.EventDispatcher;
var Event = vic.events.Event;

function DeckListModel( options ){
	EventDispatcher.call( this );
	
	this._ary_cards = [];
}
DeckListModel.prototype = {
	__proto__:EventDispatcher.prototype,
	addCard:function( cardModel ){
		if( this.hasCard( cardModel ) )	return;
		this._ary_cards.push( cardModel );
		this.dispatchEvent( new Event( 'onDeckListModelAddCard', cardModel ));
	},
	subCard:function( cardModel ){
		if( !this.hasCard( cardModel ) )	return;
		cardModel.setCount( 3 );
		this._ary_cards.splice( this._ary_cards.indexOf( cardModel ), 1 );
		this.dispatchEvent( new Event( 'onDeckListModelSubCard', cardModel ));
	},
	hasCard:function( cardModel ){
		for( var i in this._ary_cards ){
			var cm = this._ary_cards[i];
			if( cm.getId() == cardModel.getId() ){
				return true;
			}
		}
		return false;
	},
	getCards:function(){
		return this._ary_cards;
	}
}