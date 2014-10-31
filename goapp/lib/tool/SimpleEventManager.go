package tool

import (
	"reflect"
)

type SimpleEventManager struct{
	Senders map[string]func(mgr IEventManager)
	Receivers map[string][]interface{}
}

func (e *SimpleEventManager) AddSender(key string, fn func(mgr IEventManager)){
	if e.Senders == nil {
		e.Senders = map[string]func(mgr IEventManager){}
	}
	e.Senders[key] = fn
}

func (e *SimpleEventManager) RemoveSender(key string, fn func(mgr IEventManager)){
	
}

func (e *SimpleEventManager) AddReceiver(key string, fn interface{}){
	if e.Receivers == nil {
		e.Receivers = map[string][]interface{}{}
	}
	e.Receivers[key] = append(e.Receivers[key], fn)
}

func (e *SimpleEventManager) RemoveReceiver(key string, fn interface{}){
	
}

func (e *SimpleEventManager) Send(key string, params ... interface{}){
	for _, receiver := range e.Receivers[key] {
		f := reflect.ValueOf(receiver)
		in := make([]reflect.Value, len(params))
		for k, param := range params {
			in[k] = reflect.ValueOf(param)
		}
		f.Call(in)
	}
}

func (e *SimpleEventManager) SendAgain(){
	for _, sender := range e.Senders {
		sender(e)
	}
}