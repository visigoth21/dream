package handler

import (
	"dream/view/home"
	"fmt"
	"net/http"
)

func HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {
	user := getAutenticatedUser(r)
	fmt.Printf("%+v\n", user.Email)
	return home.Index().Render(r.Context(), w)
}
