package handlers

import (
	"io"
	"log"
	"net/http"
)

type employee struct {
	log *log.Logger
}

func NewEmployee(logger *log.Logger) *employee {
	return &employee{log: logger}
}

func (e *employee) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	e.log.Println("Request time")
	username, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "An error occurred", http.StatusBadGateway)
		io.WriteString(writer, "An error occurred")
	}
	log.Printf("%s", username)
	io.WriteString(writer, "Welcome employee\n")
}
