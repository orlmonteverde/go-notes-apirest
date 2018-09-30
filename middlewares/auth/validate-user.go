package auth

import (
	"fmt"
	"net/http"
)

func ValidateUser(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := validateToken(w, r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, err.Error())
			return
		}
		f(w, r)
	}
}

func ValidateAdmin(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role, err := validateToken(w, r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, err.Error())
			return
		}
		if role != "admin" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Permiso denegado")
			return
		}
		f(w, r)
	}
}
