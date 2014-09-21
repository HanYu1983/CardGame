window.api = {};
(function(){
	window.api.root = '../../test';
	
	function registerUser( account, password, type ){
		vic.utils.toURL( window.api.root, {cmd:'RegisterUser', account:account, password:password, type:type }, 'GET' )
	}
	
	function getAllUserData(){
		vic.utils.toURL( window.api.root, {cmd:'QueryUser' }, 'GET' )
	}
	
	function getUserData( name ){
		vic.utils.toURL( window.api.root, {cmd:'QueryUser', key:name }, 'GET' )
	}
	
	window.api.registerUser = registerUser;
	window.api.getAllUserData = getAllUserData;
	window.api.getUserData = getUserData;
})();