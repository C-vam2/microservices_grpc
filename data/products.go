package data

import (
	"fmt"
)

var ErrProductNotFound = fmt.Errorf("Product not found")

// Product defines the structure for an API product
// swagger:model
type Product struct {
	// the id for the product
	//
	// required: false
	// min: 1
	ID int `json:"id"` // Unique identifier for the product

	// the name for this product
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// the description for this product
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`

	//the price fo the product
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price" validate:"gt=0"`

	// the SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"sku"`
}

// Products defines a slice of Product
type Products []*Product

// GetProducts returns all products from the database
func GetProducts() Products {
	return productList
}

// getProductID returns the next productID of the product
func getProductID() int {
	lastProduct := productList[len(productList)-1]
	return lastProduct.ID + 1
}

// GetProductByID returns a single product which matches the id from the
// database.
// If a product is not found this function returns a ProductNotFound error
func GetProductByID(id int) (*Product, error) {
	idx := findIndexByProductID(id)

	if idx == -1 {
		return nil, ErrProductNotFound
	}
	return productList[idx], nil
}

// AddProduct adds a new product to the database
func AddProduct(p *Product) {
	p.ID = getProductID()
	productList = append(productList, p)
}

// UpdateProduct replaces a product in the database with the given
// item.
// If a product with the given id does not exist in the database
// this function returns a ProductNotFound error
func UpdateProduct(p *Product) error {
	idx := findIndexByProductID(p.ID)
	if idx == -1 {
		return ErrProductNotFound
	}

	//update the product in the Database
	productList[idx] = p
	return nil
}

// DeleteProduct deletes a product from the database
func DeleteProduct(id int) error {
	idx := findIndexByProductID(id)
	if idx == -1 {
		return ErrProductNotFound
	}

	productList = append(productList[:idx], productList[(idx+1):]...)

	return nil
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
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "abfjd34",
	},
}
