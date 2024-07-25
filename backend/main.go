package main

import (
	"backend/middleware"
	"backend/urlHandlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	urlHandlers.StartHandlers(mux)
	server := &http.Server{
		Addr: ":8080",
		// Addr:    ":443",
		Handler: middleware.CorsMiddleware(mux),
	}
	fmt.Println("Backend running on port 8080")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
