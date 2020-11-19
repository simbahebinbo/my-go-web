package main

import "github.com/julienschmidt/httprouter"

type HttpRouter struct {
	Path   string
	Method string
	Handle httprouter.Handle
}

func AddRouter(routers []HttpRouter, router *httprouter.Router) {
	if len(routers) <= 0 {
		return
	}

	for _, value := range routers {
		router.Handle(value.Method, value.Path, value.Handle)
	}
}
