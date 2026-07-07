package main

import (
	"os"

	"log"
	"net/http"

	"github.com/microservices_grpc/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	sm := http.NewServeMux()

	sm.Handle("/", hh)

	http.ListenAndServe(":9090", sm)
}
