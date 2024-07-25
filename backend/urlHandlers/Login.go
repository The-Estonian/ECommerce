package urlHandlers

import (
	"backend/helpers"
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login attempt!")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	var callback = make(map[string]string)
	callback["email"] = email
	callback["password"] = password
	writeData, err := json.Marshal(callback)
	helpers.CheckErr("handleLogin", err)
	w.Write(writeData)
}
