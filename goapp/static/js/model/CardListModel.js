window.app = window.app || {};
window.app.model = window.app.model || {};
window.app.model.CardListModel = CardListModel;

var EventDispatcher = vic.events.EventDispatcher;
var Event = vic.events.Event;

function CardListModel( options ){
	EventDispatcher.call( this );
	this._ary_model = options.ary_model;
}

CardListModel.prototype = {
	__proto__:EventDispatcher.prototype,
	init:function(){
		for( var i = 0; i < this._ary_model.length; ++i ){
			var model = this._ary_model[i];
			this.dispatchEvent( new Event( 'onCardListModelAddCard', {index:i, model:model} ));
		}
	}
}