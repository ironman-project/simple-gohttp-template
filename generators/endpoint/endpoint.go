package main

import (
	"fmt"
	"net/http"
)


var _ = registerEndpoint("{{.Values.endpoint}}", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New {{.Values.endpoint | default "Empty endpoint"}}")
})
