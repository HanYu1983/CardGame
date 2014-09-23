window.app = window.app || {};
window.app.view = window.app.view || {};
window.app.view.CardListView = CardListView;

var EventDispatcher = vic.events.EventDispatcher;
var Event = vic.events.Event;

function CardListView( options ){
	EventDispatcher.call( this );
	this._root = options.root;
	this._cardTemplete = options.cardTemplete;
}

CardListView.prototype = {
	__proto__:EventDispatcher.prototype,
	addCard:function( index, model ){
		var itemroot = $('<div class="floatLeft item"></div>' );
		//var controlroot = $('<div class="itemControlContainer"></div>' );
		//var checkbox = $('<input type="checkbox" id="chk_want_' + index + '" class="chk_want"/><label for="check">need</label>' );
		var clickborder = $('<div class="clickborder"></div>');
		clickborder.attr('id', 'btn_border_' + index );
		
		var cardroot = this.getCardTemplate( model );
		//controlroot.append( checkbox);
		itemroot.append( cardroot );
		itemroot.append( clickborder );
		//itemroot.append( controlroot );
		var self = this;
		clickborder.css( 'cursor', 'pointer' );
		clickborder.click( function(){
			var target = $(this);
			var id = target.attr( 'id' );
			var cid = id.substring( 'btn_border_'.length, id.length );
			var val;
			if( target.hasClass( 'redOutline' )){
				$( this ).removeClass( 'redOutline' );
				val = false;
			}else{
				$( this ).addClass( 'redOutline' );
				val = true;
			}
			self.dispatchEvent( new Event( 'onCardListViewChkNeedChange', {cid:cid, val:val } ));
		});
		/*
		var self = this;
		checkbox.change( function(){
			var target = $(this);
			var id = target.attr( 'id' );
			var cid = id.substring( 'chk_want_'.length, id.length );
			var val = target[0].checked;
			self.dispatchEvent( new Event( 'onCardListViewChkNeedChange', {cid:cid, val:val } ));
		});
		*/
		this._root.append( itemroot );
	},
	setTargetNeed:function( id, val ){
		var target = this._root.find( '#chk_want_' + id );
		target.attr( 'checked', val );
	},
	getCardTemplate:function( card ){
		var temp = this._cardTemplete.clone();
		temp.find( '.name' ).html( card.Name );
		temp.find( '.content' ).html( card.Content );
		temp.find( '.action' ).html( card.Action );
		temp.find( '.power' ).html( card.Power );
		temp.find( '.magic' ).html( card.Magic );
		temp.find( '.type' ).html( card.AtkType );
		temp.find( '.card_type' ).html( card.Type );
		temp.find( '.action_cost' ).html( card.ActionCost );
		temp.find( '.action_type' ).html( card.ActionType );
		temp.find( '.action_content' ).html( card.ActionContent );
		temp.find( '.weight' ).html( '重' + card.Weight );
		temp.find( '.id' ).html( card.Key );
		temp.find( '.level' ).html( '星' + card.Level );
		temp.find( '.tag' ).html( card.Tag );
		return temp;
	}
}

