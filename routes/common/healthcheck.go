package common

import (
	"fmt"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello there, this is an example")
	w.WriteHeader(http.StatusOK)
}
