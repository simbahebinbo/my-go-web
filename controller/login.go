package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserNamePwd struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

func UserNamePwdLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var u UserNamePwd
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("UserNamePwd:", u)
	json, _ := json.Marshal(u)
	w.Write(json)
}

func GetUserNameInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	values := r.URL.Query()
	for k, v := range values {
		fmt.Println(k, v)
	}

	u := UserNamePwd{"第三方都", "觉得舒服的"}
	fmt.Println("UserNamePwd:", u)
	json, _ := json.Marshal(u)
	w.Write(json)
}
