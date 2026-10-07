package handler

import (
	"fmt"
	"net/http"
)

func Status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "API está funcionando!")
}
