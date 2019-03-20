package main

import (
	"net/http"
	"os"
)

type handler struct{}

func (handler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	filename := "file.zip"
	if _, err := os.Open(filename); err != nil {
		responseWriter.WriteHeader(http.StatusNotFound)
		return
	}
	responseWriter.Header().Set("Content-Type", "applicaiton/octet-stream")
	responseWriter.Header().Set("Content-Disposition", "attachment; filename=client.ovpn")
	http.ServeFile(responseWriter, request, filename)
}

func main() {
	http.ListenAndServe(":10001", handler{})
}
