package main

import (
	"net/http"
)

func liveness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusGatewayTimeout)
}
