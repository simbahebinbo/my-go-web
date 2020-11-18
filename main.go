package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/scosky/myGoWeb/handler"
)

func main() {
	router := httprouter.New()

	user := "admin"
	pass := "admin"

	router.GET("/login/", handler.BasicAuth(Protected, user, pass))
	router.GET("/get/:name", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("employee name:" + p.ByName("name")))
	})
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Protected(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Protected!\n")
}
