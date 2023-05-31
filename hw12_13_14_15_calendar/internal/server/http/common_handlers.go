package internalhttp

import (
	"net/http"
)

func NotImplemented(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

func Hello(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	w.Write([]byte("hello"))
}

func Error(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, "internal server error", http.StatusInternalServerError)
}
