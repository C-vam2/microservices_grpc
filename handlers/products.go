package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/microservices_grpc/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		reg := regexp.MustCompile(`/([0-9]+)`)

		matches := reg.FindAllStringSubmatch(r.URL.Path, -1)
		fmt.Print(r.URL.Path)
		fmt.Print(len(matches))
		for _, match := range matches {
			fmt.Println(match)
		}
		if len(matches) != 1 {
			http.Error(rw, "ID not found in the URL", http.StatusBadRequest)
			return
		}

		lastItem := matches[0][1]

		id, err := strconv.Atoi(lastItem)

		if err != nil {
			http.Error(rw, "Unable to parse ID", http.StatusInternalServerError)
			return
		}
		p.updateProduct(rw, r, id)
		return

	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("POST request triggered!!")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to process request body.", http.StatusBadRequest)
		return
	}
	data.AddProduct(prod)
}

func (p *Products) updateProduct(rw http.ResponseWriter, r *http.Request, id int) {
	p.l.Println("PUT request triggered!!")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to process request body.", http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(id, prod)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

}
