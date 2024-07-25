package urlHandlers

import (
	"backend/helpers"
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRegistration(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register attempt!")

	var callback = make(map[string]string)

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		helpers.CheckErr("backend/urlHandlers/HandleRegistration -> MultipartForm", err)
		callback["error"] = "Error parsing register data"
	} else {
		email := r.FormValue("email")
		password := r.FormValue("password")
		firstName := r.FormValue("firstName")
		lastName := r.FormValue("lastName")

		callback["email"] = email
		callback["password"] = password
		callback["firstName"] = firstName
		callback["lastName"] = lastName
	}

	writeData, err := json.Marshal(callback)
	helpers.CheckErr("backend/urlHandlers/HandleRegistration -> Marshal", err)
	w.Write(writeData)
}
