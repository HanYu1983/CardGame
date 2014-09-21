var vic = vic || {};
vic.mvc = vic.mvc || {};
(function(){
	function Command( cname ){
		this._name = cname;
	}
	Command.prototype = {
		execute:function( view, from ){
			
		},
		getName:function(){return this._name; }
	}
	vic.mvc.Command = Command;
})();

(function(){
	var Event = vic.events.Event;
	var EventDispatcher = vic.events.EventDispatcher;
	//var HashManager = vic.window.HashManager;
	function QueryView( query, pname, view ){
		this.isShow = true;
		this._query = query;
		this._name = pname;
		this._view = view;
		this._ary_buttonQuerys = [];
		//this._hashManager = HashManager.getInstance();
		this.hide();
	}
	
	QueryView.prototype = {
		__proto__:EventDispatcher.prototype,
		init:function(){
			this.parseQuery( this._query );
		},
		show:function(){
			if( this.isShow )	return;
			this.isShow = true;
			this._query.show();
			this.addListener();
			//if( withHash != undefined && withHash == false )	return;
			//else this._hashManager.push( this._name );
		},
		hide:function(){
			this.isShow = false;
			this._query.hide();
			this.removeListener();
			//this._hashManager.pop();
		},
		getQuery:function(){
			return this._query;
		},
		getView:function(){
			return this._view;
		},
		getName:function(){
			return this._name;
		},
		addListener:function(){
			var b;
			var tempthis = this;
			for( var i = 0; i < this._ary_buttonQuerys.length; ++i ){
				b = this._ary_buttonQuerys[i];
				b.click( function(){
					tempthis._view.execute( $( this ).attr( 'commandName') );
				});
			}
		},
		removeListener:function(){
			var b;
			for( var i = 0; i < this._ary_buttonQuerys.length; ++i ){
				b = this._ary_buttonQuerys[i];
				b.off( 'click' );
			}
		},
		parseQuery:function( query ){
			var str_result = query.attr( 'vicmvc' );
			var ary_result;
			var type;
			var data;
			if( str_result != undefined ){
				ary_result = str_result.split( ':' );
				type = ary_result[0];
				data = ary_result[1];
				switch( type ){
					case 'button':
						var commandName = data
						query.attr( 'commandName', commandName );
						query.css( 'cursor', 'pointer' );
						this._ary_buttonQuerys.push( query );
						this.dispatchEvent( new Event( 'addCommand', commandName ));
						break;
				}
				
			}
			if( query.children().length > 0 ){
				for( var i = 0, max = query.children().length; i < max; ++i ){
					this.parseQuery( query.children().eq(i) );
				}
			}
		}
	}
	
	vic.mvc.QueryView = QueryView;
})();

(function(){
	//var HashManager = vic.window.HashManager;
	function View( hashName ){
		this._hashName = hashName;
		this._ary_pages = [];
		this._ary_commands = [];
		//this._hashManager = HashManager.getInstance();
		var self = this;
		/*
		this._hashManager.addEventListener( 'onhashchange', function( e ){
			console.log( 'hashstr', e.getData() );
			var ary_strs = e.getData().split( '\\' );
			self.openPage( ary_strs.pop(), false );
		});*/
	}
	
	View.prototype = {
		execute:function( cname, from ){
			var c = this.getCommand( cname );
			if( c == undefined )	throw 'not have this command: ' + cname;
			c.execute( this, from );
		},
		addCommand:function( cmd ){
			if( this.getCommand( cmd.getName()) != undefined )	return;
			this._ary_commands.push( cmd );
		},
		getCommand:function( cname ){
			var c;
			for( var i = 0; i < this._ary_commands.length; ++i ){
				c = this._ary_commands[i];
				if( c.getName() == cname )	return c;
			}
			return undefined;
		},
		addPage:function( page ){
			this._ary_pages.push( page );
		},/*
		pushHash:function(){
			var p = this.getPage( pname );
			if( p != undefined )	this._hashManager.push( pname );
		},
		popHash:function(){
			var p = this.getPage( pname );
			if( p != undefined )	this._hashManager.pop();
		},*/
		openPage:function( pname, withHash ){
			var p = this.getPage( pname );
			if( p != undefined )	p.show( withHash );
		},
		closePage:function( pname ){
			var p = this.getPage( pname );
			if( p != undefined )	p.hide();
		},
		getPage:function( pname ){
			var p;
			for( var i = 0; i < this._ary_pages.length; ++i ){
				p = this._ary_pages[i];
				if( p.getName() == pname )	return p;
			}
			return undefined;
		},
		getPages:function(){
			return this._ary_pages;
		}
	}
	
	vic.mvc.View = View;
})();

//<div id="mainPage" vicmvc="page:MainPage">
//<button id="btn_doA" vicmvc="button:DoACommand">next</button>
(function(){
	var View = vic.mvc.View;
	function MVCParser( factory ){
		var v = new View();
		var category = $( "div[vicmvc]" );
		var str_result;
		var ary_result;
		var type;
		var data;
		var queryName, className;
		var targetQuery;
		var page;
		for( var i = 0, max = category.length; i < max; ++i ){
			targetQuery = category.eq( i );
			str_result = targetQuery.attr( 'vicmvc' );
			ary_result = str_result.split( ':' );
			type = ary_result[0];
			data = ary_result[1];
			switch( type ){
				case 'page':
					queryName = targetQuery.attr( 'id' );
					className = data;
					if( className == 'QueryView' )	factory[className] = vic.mvc.QueryView;
					page = new factory[className]( $('#'+ queryName ), queryName, v );
					v.addPage( page );
					page.addEventListener( 'addCommand', onAddCommandEvent );
					page.init();
					break;
			}
		}
		
		function onAddCommandEvent( e ){
			var className = e.getData();
			v.addCommand( new factory[className]( className ));
		}
		
		var pages = v.getPages();
		for( var i = 0, max = pages.length; i < max; ++i ){
			page = pages[i];
			page.removeEventListener( 'addCommand', onAddCommandEvent );
		}
		return v;
	}
	vic.mvc.MVCParser = MVCParser;
})();