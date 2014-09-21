var vic = vic || {};
vic.google = vic.google || {};
vic.google.api = vic.google.api || {};

(function(){
	var _clientId;
	var _apiKey;
	var _scopes;
	var _onApiReady;
	var _onNoLogin;
	
	function init( options ){
		_clientId = options.clientId;
		_apiKey = options.apiKey;
		_scopes = options.scopes;
		_onApiReady = options.onApiReady;
		_onNoLogin = options.onNoLogin;
		_handleGoogleClientLibrary();
	}
	
	function _handleGoogleClientLibrary(){
		gapi.client.setApiKey( _apiKey );
		window.setTimeout( _checkAuth, 1 );
	}
	
	function _checkAuth() {
        gapi.auth.authorize({client_id: _clientId, scope: _scopes, immediate: true}, _handleAuthResult);
    }
	
	function _handleAuthResult(authResult) {
		if (authResult && !authResult.error) {
			_onApiReady( authResult );
		} else {
			if( _onNoLogin != undefined )	_onNoLogin();
		}
	}
	
	function handleAuth() {
		gapi.auth.authorize({client_id: _clientId, scope: _scopes, immediate: false}, _handleAuthResult);
		return false;
	}
	
	vic.google.api.init = init;
	vic.google.api.handleAuth = handleAuth;
})();
//<script src="https://apis.google.com/js/client.js"></script>