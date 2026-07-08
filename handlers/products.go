package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/microservices_grpc/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	lp := data.GetProducts()
	// lp.ToJSON();
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(res, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	res.Write(d)
}
