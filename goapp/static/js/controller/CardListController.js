window.app = window.app || {};
window.app.controller = window.app.controller || {};
window.app.controller.CardListController = CardListController;

var EventDispatcher = vic.events.EventDispatcher;
var Event = vic.events.Event;
var CardProxyModel = window.app.model.CardProxyModel;
var CardListModel = window.app.model.CardListModel;
var CardListView = window.app.view.CardListView;

function CardListController( options ){
	EventDispatcher.call( this );
	var self = this;
	var ary_cardModelProxy = options.ary_cardModelProxy;
	var model = new CardListModel( {
		ary_model:options.ary_model
	} );
	this.view = new CardListView( {
		root:options.root,
	} );
	
	model.addEventListener( 'onCardListModelAddCard', function( e ){
		var retIndex = e.getData().index;
		var retModel = e.getData().model;
		retModel.cid = retIndex;
		self.view.addCard( retIndex );
		ary_cardModelProxy.push( new CardProxyModel( {cid:retIndex, model:retModel} ));
	});
	var self = this;
	this.view.addEventListener( 'onCardListViewChkNeedChange', function( e ){
		var retcid = e.getData().cid;
		var retval = e.getData().val;
		self.dispatchEvent( e );
	});
	
	model.init();
}

CardListController.prototype = {
	__proto__:EventDispatcher.prototype,
	setTargetNeed:function( id, val ){
		this.view.setTargetNeed( id, val );
	}
}
