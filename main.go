package main

import (
	"io"

	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		log.Println("hello world")
		data, err := io.ReadAll(req.Body)

		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("An error occured"))
			return
		}
		log.Println(string(data))
	})

	http.ListenAndServe(":9090", nil)
}
