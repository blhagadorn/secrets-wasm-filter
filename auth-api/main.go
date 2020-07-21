package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You've reached the endpoint!")
	})

	router.HandleFunc("/infosecure", func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		split := strings.Split(token, "Bearer ")
		fmt.Println(token)
		if len(split) > 1 {
			if split[1] == "secretpassword" {
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, "Authorized")
			} else {
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, "Unauthorized request")
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Unknown request")
		}
	})

	http.ListenAndServe(":8080", router)
}
