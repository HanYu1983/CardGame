package tool

import (
	"net/http"
)

type AppEngineSystem struct {
	Request *http.Request
	Response http.ResponseWriter
}

func (s AppEngineSystem) GetRequest() *http.Request{
	return s.Request
}

func (s AppEngineSystem) GetResponse() http.ResponseWriter{
	return s.Response
}

func (s AppEngineSystem) Log(msg interface{}){
	var logger ILogger
	logger = &SimpleLogger{Writer: s.Response}
	logger.Log(msg)
}

