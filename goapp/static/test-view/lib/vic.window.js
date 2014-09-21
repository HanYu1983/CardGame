var vic = vic || {};
vic.window = vic.window || {};
(function(){
	var Event = vic.events.Event;
	var EventDispatcher = vic.events.EventDispatcher;
	
	function HashManager(){
		EventDispatcher.call( this );
		window.onhashchange = this.onHashChange;
	}
	
	HashManager.prototype = {
		__proto__:EventDispatcher.prototype,
		push:function( hashstr ){
			hashstr = "\\" + hashstr;
			window.location.hash += hashstr;
		},
		pop:function(){
			var ary_strs = this.getHash().split('\\');
			ary_strs.pop();//pop last child
			ary_strs.shift();//pop #
			var newstrs = '';
			for( var i = 0; i < ary_strs.length; ++i ){
				newstrs += '\\' + ary_strs[i];
			}
			window.location.hash = newstrs;
		},
		getHash:function(){
			return window.location.hash;
		},
		onHashChange:function(){
			HashManagerInstance.dispatchEvent( new Event( 'onhashchange', window.location.hash ));
		}
	}
	
	var HashManagerInstance;
	HashManager.getInstance = function(){
		if( HashManagerInstance == null )	HashManagerInstance = new HashManager();
		return HashManagerInstance;
	}
	
	vic.window.HashManager = HashManager;
})();


