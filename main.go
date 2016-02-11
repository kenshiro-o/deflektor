package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", deflekt)
	http.ListenAndServe(":8081", nil)
}

func deflekt(w http.ResponseWriter, r *http.Request) {
	// Capture host:port and just take the host part
	hp := strings.Split(r.Host, ":")
	secureSite := fmt.Sprintf("https://%s%s", hp[0], r.RequestURI)
	http.Redirect(w, r, secureSite, http.StatusTemporaryRedirect)
}
