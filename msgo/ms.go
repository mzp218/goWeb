package msgo

import (
	"log"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type router struct {
	handlerMap map[string]HandlerFunc
}

func (r *router) Add(name string, handlerFunc HandlerFunc) {
	r.handlerMap[name] = handlerFunc
}

type Engine struct {
	router
}

func New() *Engine {
	return &Engine{
		router{handlerMap: make(map[string]HandlerFunc)},
	}
}

func (e *Engine) Run() {
	for key, value := range e.handlerMap {
		http.HandleFunc(key, value)
	}
	err := http.ListenAndServe(":8111", nil)
	if err != nil {
		log.Fatal(err)
	}
}