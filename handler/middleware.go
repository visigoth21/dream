package handler

import (
	"fmt"
	"net/http"
)

func WithUser(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// Do something with the user
		fmt.Println("from the with user middleware")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
