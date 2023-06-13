package internalhttp

import (
	"net/http"
)

func notImplemented(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

func hello(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	w.Write([]byte("hello"))
}

func serverError(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, "internal server error", http.StatusInternalServerError)
}
