package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/scosky/myGoWeb/controller"
)

func main() {
	router := httprouter.New()
	router.POST("/login", controller.UserNamePwdLogin)
	router.GET("/get", controller.GetUserNameInfo)
	log.Fatal(http.ListenAndServe(":8080", router))
}
