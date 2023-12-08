package handlers

import (
	"io"
	"log"
	"net/http"
)

type home struct {
	log *log.Logger
}

func NewHome(logger *log.Logger) *home {
	return &home{log: logger}
}

func (h *home) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	h.log.Println("Hello world")
	io.WriteString(writer, "Welcome to Open Portal\n")
}
