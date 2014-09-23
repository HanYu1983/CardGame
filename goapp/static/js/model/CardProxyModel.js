window.app = window.app || {};
window.app.model = window.app.model || {};
window.app.model.CardProxyModel = CardProxyModel;

function CardProxyModel( options ){
	this._cid = options.cid;
	this._model = options.model;
	this._count = 3;
}
CardProxyModel.prototype = {
	add:function(){
		this._count++;
	},
	sub:function(){
		this._count--;
	},
	setCount:function( val ){
		this._count = val;
	},
	count:function(){
		return this._count;
	},
	getId:function(){
		return this._cid;
	},
	getName:function(){
		return this._model.name;
	},
	getModel:function(){
		return this._model;
	}
}