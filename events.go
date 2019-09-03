package events

import (
	"errors"
)

type Events struct {
	Listeners map[string][]func(data interface{})
	Limit int
}
func (e *Events) Init() {
	e.Listeners = map[string][]func(data interface{}){}
}
func (e *Events) On(name string,cb func(data interface{})) {
	_,found := e.Listeners[name]
	if !found {
		e.Listeners[name] = []func(data interface{}){cb}
		return
	}else{
		e.Listeners[name] = append(e.Listeners[name], cb)
	}

}
func (e *Events) Emit(name string,data interface{}) {
	Listeners,found := e.Listeners[name]
	if !found {
		//console.Log("no listener")
	}else{
		for _,item := range Listeners {
			item(data)
		}
	}
}
func (e *Events) Once(name string,cb func(data interface{})) {
	e.On(name,cb)
	err := e.Clear(name)
	if err != nil {
		panic(err)
	}
}
func (e *Events) Clear(name string) (err error) {
	_,found := e.Listeners[name]
	if !found {
		err = errors.New("Error event name:" + name)
		return
	}
	delete(e.Listeners,name)
	return
}
