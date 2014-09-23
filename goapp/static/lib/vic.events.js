var vic = vic || {};
vic.events = vic.events || {};
(function(){
	function Event( type, data ){
		this._type = type;
		this._data = data;
	}
	
	Event.prototype = {
		getType:function(){
			return this._type;
		},
		getData:function(){
			return this._data;
		}
	}
	vic.events.Event = Event;
})();

(function(){
	var Event = vic.events.Event;
	function MouseEvent( type, data ){
		Event.call( this, type, data );
		this.localX;
		this.localY;
		this.stageX;
		this.stageY;
	}
	
	MouseEvent.prototype = {
		__proto__:Event.prototype
	}
	
	vic.events.MouseEvent = MouseEvent;
	vic.events.MouseEvent = MouseEvent;
	vic.events.MouseEvent.CLICK = 'click';
	vic.events.MouseEvent.MOUSE_MOVE = 'mousemove';
	vic.events.MouseEvent.MOUSE_OVER = 'mouseover';
	vic.events.MouseEvent.MOUSE_OUT = 'mouseout';
	vic.events.MouseEvent.MOUSE_DOWN = 'mousedown';
	vic.events.MouseEvent.MOUSE_UP = 'mouseup';
})();

(function(){
	var Event = vic.events.Event;
	function TouchEvent( type, data ){
		MouseEvent.call( this, type, data );
		this.localX;
		this.localY;
		this.stageX;
		this.stageY;
	}
	
	TouchEvent.prototype = {
		__proto__:Event.prototype
	}
	
	vic.events.TouchEvent = TouchEvent;
	vic.events.TouchEvent.TOUCH_START = 'touchstart';
	vic.events.TouchEvent.TOUCH_END = 'touchend';
	vic.events.TouchEvent.TOUCH_CANCEL = 'touchcancel';
	vic.events.TouchEvent.TOUCH_LEAVE = 'touchleave';
	vic.events.TouchEvent.TOUCH_MOVE = 'touchmove';
})();

(function(){
	var Event = vic.events.Event;
	function GestureEvent( type, data ){
		Event.call( this, type, data );
		this.phase;
		this.offsetX;
		this.offsetY;
		this.scaleX;
		this.scaleY;
		this.degreeOffset;
	}
	
	GestureEvent.prototype = {
		__proto__:Event.prototype
	}
	
	vic.events.GestureEvent = GestureEvent;
	vic.events.GestureEvent.ZOOM = 'GestureEvent.ZOOM';
	vic.events.GestureEvent.ONE_FINGER = 'GestureEvent.ONE_FINGER';
	vic.events.GestureEventPhase = {
		BEGIN:'GestureEventPhase.BEGIN',
		UPDATE:'GestureEventPhase.UPDATE',
		END:'GestureEventPhase.END'
	}
})();

(function(){
	function EventDispatcher(){ }
	
	EventDispatcher.prototype = {
		addEventListener:function( n, fn ){
			if( this._observers == undefined ){
				this._observers = {};
			}
			if( this._observers[n] == undefined ){
				this._observers[n] = [];
			}
			if( this._observers[n].indexOf( fn ) == -1 ){
				this._observers[n].push( fn );
			}
		},
		removeEventListener:function( n, fn ){
			if( this._observers == undefined )	return;
			if( this._observers[n] == undefined )	return;
			var ary = this._observers[n];
			var id = ary.indexOf( fn );
			if( id != -1 ){
				ary.splice( id, 1 );
			}
			if( ary.length == 0 ){
				delete this._observers[n];
			}
		},
		removeAllEventListener:function(){
			if( this._observers == undefined )	return;
			this._observers = undefined;
		},
		dispatchEvent:function( e ){
			if( this._observers == undefined )	return;
			var en = e.getType();
			if( this._observers[ en ] != undefined ){
				e.currentTarget = this;
				var k, ary = this._observers[ en ]
				for ( k in ary ){
					ary[k]( e );
				}
			}
		}
	}
	
	vic.events.EventDispatcher = EventDispatcher;
})();

