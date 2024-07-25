package urlHandlers

import (
	"backend/helpers"
	"backend/validate"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func StartHandlers(r *mux.Router) {
	r.HandleFunc("/login", HandleLogin).Methods("POST")
	r.HandleFunc("/register", HandleRegistration).Methods("POST")
}

func handleSession(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("ecommerce")
	if err != nil || validate.ValidateUserSession(cookie.Value) == "0" {
		var callback = make(map[string]string)
		// check status
		sessionCookie := http.Cookie{
			Name:     "ecommerce",
			Value:    "",
			Expires:  time.Now(),
			Path:     "/",
			HttpOnly: true,
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
		}
		http.SetCookie(w, &sessionCookie)
		authCookie := http.Cookie{
			Name:     "ecommerce",
			Value:    "false",
			Expires:  time.Now(),
			Path:     "/",
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
		}
		http.SetCookie(w, &authCookie)
		callback["login"] = "fail"
		writeData, err := json.Marshal(callback)
		helpers.CheckErr("handleSession", err)
		w.Write(writeData)
		return false
	}
	return true
}
