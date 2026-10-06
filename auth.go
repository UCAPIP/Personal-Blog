package main

import (
	"crypto/subtle"
	"net/http"
)

func basicAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		expectedUsername := "admin"
		expectedPassword := "admin"

		usernameMatch := subtle.ConstantTimeCompare([]byte(username), []byte(expectedUsername)) == 1
		passwordMatch := subtle.ConstantTimeCompare([]byte(password), []byte(expectedPassword)) == 1

		if !ok || !usernameMatch || !passwordMatch {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted Area"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
