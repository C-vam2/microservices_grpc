package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	h.l.Println("hello world")
	data, err := io.ReadAll(req.Body)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("An error occured"))
		return
	}
	fmt.Println(string(data))
}
