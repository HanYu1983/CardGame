var vic = vic || {};
vic.utils = vic.utils || {};

vic.utils.includeJS = function( ary_path ){
	for( var i = 0; i < ary_path.length; ++i ){
		document.write( '<script type="text/javascript" src="' + ary_path[i] + '"></script>' );
	}
}

vic.utils.includeCSS = function( ary_path ){
	for( var i = 0; i < ary_path.length; ++i ){
		document.write( '<link href="' + ary_path[i] + '" rel="stylesheet" type="text/css" />' );
	}
}

vic.utils.createElement = function( type, obj ){
	var element = document.createElement( type );
	for( prop in obj ){
		element[ prop ] = obj[ prop ];
	}
	return element;
}

vic.utils.getElement = function( id ){
	return document.getElementById( id );
}

vic.utils.appendChild = function(){
	if( arguments.length > 1 ){
		var a = arguments[0];
		for( i = 1; i < arguments.length; ++i ){
			if( arguments[i] ){
				a.appendChild( arguments[i] );
			}
		}
		return a;
	}else{
		return null;
	}
}

vic.utils.removeChildren = function( node ){
	if( node == null )	return;
	while( node.hasChildNodes() ){
		node.removeChild( node.firstChild );
	}
}

vic.utils.addEventListener = function( obj, eventName, listener ){
	if( obj.attachEvent )			obj.attachEvent( 'on' + eventName, listener );
	else if( obj.addEventListener )	obj.addEventListener( eventName, listener, false );
	else							return false;
	return true;
}

vic.utils.removeEventListener = function( obj, eventName, listener ){
	if( obj.detachEvent )				obj.detachEvent( 'on' + eventName, listener );
	else if( obj.removeEventListener )	obj.removeEventListener( eventName, listener, false );
	else return false;
	return true;
}

vic.utils.changeOpacity = function( opacity, id ){
	var styleobj = document.getElementById( id ).style;
	styleobj.opacity = opacity;
	styleobj.MozOpacity = opacity;
	styleobj.KhtmlOpacity = opacity;
	styleobj.filter = 'alpha(opacity=' + opacity + ')';
}

vic.utils.inherit = function(from, fields) {
	function Inherit() {} Inherit.prototype = from.prototype; var proto = new Inherit();
	for (var name in fields) proto[name] = fields[name];
	if( fields.toString !== Object.prototype.toString ) proto.toString = fields.toString;
	proto.uber = from.prototype;
	return proto;
}

vic.utils.toURL = function( url, data, method, onComplete ){
	$.ajax({
		url: url,
		data: data,
		type: method,
		dataType: 'json',
		success: function(msg){
			console.log( msg );
			if( msg.Error != undefined ){
				
			}
			if( onComplete != undefined )	onComplete( msg );
		},
		error:function(xhr, ajaxOptions, thrownError){ 
			alert(xhr.status); 
			alert(thrownError); 
		}
	});
}

