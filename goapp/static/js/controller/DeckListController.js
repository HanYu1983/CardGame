window.app = window.app || {};
window.app.controller = window.app.controller || {};
window.app.controller.DeckListController = DeckListController;

var EventDispatcher = vic.events.EventDispatcher;
var Event = vic.events.Event;

function DeckListController( options ){
	EventDispatcher.call( this );
	var self = this;
	
	this.model = new window.app.model.DeckListModel();
	this.view = new window.app.view.DeckListView({
		root:options.root
	});
	
	this.model.addEventListener( 'onDeckListModelAddCard', function( e ){
		var model = e.getData();
		self.view.addCard( model );
	});
	
	this.model.addEventListener( 'onDeckListModelSubCard', function( e ){
		var model = e.getData();
		self.view.subCard( model );
		self.dispatchEvent( e );
	});
	
	this.view.addEventListener( 'onDeckListViewSprCountChange', function( e ){
		self.dispatchEvent( e );
	});
	
	this.view.addEventListener( 'onDeckListViewBtnDeleteClick', function( e ){
		self.dispatchEvent( e );
	});
}

DeckListController.prototype = {
	__proto__:EventDispatcher.prototype,
	
	addCard:function( model ){
		this.model.addCard( model );
	},
	subCard:function( model ){
		this.model.subCard( model );
	},
	getCards:function(){
		return this.model.getCards();
	}
}
