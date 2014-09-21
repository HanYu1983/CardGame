/**
 * ...
 * @author vic
 
使用方法:
vic.ajax.makeRequest( 'GET', 'testXML.xml', function(){
	if( vic.ajax.checkReadyState() )	{
		vic.ajax.makeRequest( 'GET', 'testJson.txt', function(){
			if( vic.ajax.checkReadyState() )	console.log('b', vic.ajax.getResponse() );
		});
	}
});
*/

var vic = vic || {};
vic.ajax = vic.ajax || {};
vic.ajax.makeRequest = function( method, url, callback ){
	this.request = (window.XMLHttpRequest ) ? new XMLHttpRequest() : new ActiveXObject( 'MSXML2.XMLHTTP' );
	this.request.onreadystatechange = callback;
	this.request.open( method, url, true );
	this.request.send( url );
}
vic.ajax.checkReadyState = function( onReadyStateEvent ){
	if( onReadyStateEvent != undefined )	onReadyStateEvent( this.request.readyState );
	switch( this.request.readyState ){
		case 1:	case 2:	case 3:	return false;
		case 4:	return true;
		default:return false;
	}
}
vic.ajax.getResponse = function(){
	if( this.request.getResponseHeader( 'Content-Type' ).indexOf( 'xml' ) != -1 ){
		return this.request.responseXML.documentElement;
	}else{
		return this.request.responseText;
	}
}




