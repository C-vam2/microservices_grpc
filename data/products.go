package data

import (
	"encoding/json"
	"errors"
	"io"
	"time"

	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Product) Validate() error {
	validator := validator.New()
	validator.RegisterValidation("sku", validateSKU)
	return validator.Struct(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)

	return d.Decode(p)

}

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

func getProductID() int {
	lastProduct := productList[len(productList)-1]
	return lastProduct.ID + 1
}

func AddProduct(p *Product) {
	p.ID = getProductID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	for idx, prd := range productList {
		if prd.ID == id {
			p.ID = id
			productList[idx] = p
			return nil
		}
	}

	return errors.New("Product not found with given id")
}

// findIndex find the index oof a product in the database
// returns -1 when no product can be found
func findIndexByProductID(id int) int {
	for idx, prod := range productList {
		if prod.ID == id {
			return idx
		}
	}

	return -1
}

var productList = Products{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "abfjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
