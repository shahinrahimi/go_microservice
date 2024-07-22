package data

import "fmt"

// ErrProductNotFound is an error raised when the product not found in the database
var ErrProductNotFound = fmt.Errorf("Product not found")

// Product defines the structure for an API product
// swagger:model
type Product struct {
	// the id for the product
	//
	// required: false
	// min: 1
	ID int `json:"id"` // Unique identifier for the product

	// the name for this poduct
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// the description for this poduct
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`

	// the price for the product
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price" validate:"required,gt=0"`

	// the SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"sku"`
}

// Products define a slice of Product
type Products []*Product

var productList = []*Product{
	{
		ID:          1,
		Name:        "Espresso",
		Description: "Strong and bold espresso coffee.",
		Price:       2.50,
		SKU:         "coffee-espresso-basic",
	},
	{
		ID:          2,
		Name:        "Cappuccino",
		Description: "Espresso with steamed milk and foam.",
		Price:       3.00,
		SKU:         "coffee-cappuccino-foam",
	},
	{
		ID:          3,
		Name:        "Latte",
		Description: "Espresso with steamed milk and a light layer of foam.",
		Price:       3.50,
		SKU:         "coffee-latte-smooth",
	},
	{
		ID:          4,
		Name:        "Americano",
		Description: "Espresso with added hot water.",
		Price:       2.75,
		SKU:         "coffee-americano-water",
	},
	{
		ID:          5,
		Name:        "Mocha",
		Description: "Espresso with chocolate, steamed milk, and whipped cream.",
		Price:       4.00,
		SKU:         "coffee-mocha-choco",
	},
	{
		ID:          6,
		Name:        "Macchiato",
		Description: "Espresso with a small amount of steamed milk and foam.",
		Price:       3.25,
		SKU:         "coffee-macchiato-steam",
	},
	{
		ID:          7,
		Name:        "Flat White",
		Description: "Espresso with steamed milk, similar to a latte but with less foam.",
		Price:       3.75,
		SKU:         "coffee-flatwhite-creamy",
	},
}

// GetProducts return all products from the DB
func GetProducts() Products {
	return productList
}

// GetProductByID returns a single Product which matches the id from the DB
// if a product is not found this function returns a ProductNotFoundError
func GetProductByID(id int) (*Product, error) {
	i := findIndexByProductID(id)
	if i == -1 {
		return nil, ErrProductNotFound
	}
	return productList[i], nil
}

// AddProduct adds a new product to the database
func AddProduct(p Product) {
	maxID := productList[len(productList)-1].ID
	p.ID = maxID + 1
	productList = append(productList, &p)
}

// UpdateProduct replaces a product in the database with the given
// item.
// If a product with the given id does not exist in the database
// this function returns a ProductNotFound error
func UpdateProduct(p Product) error {
	i := findIndexByProductID(p.ID)
	if i == -1 {
		return ErrProductNotFound
	}
	productList[i] = &p
	return nil
}

// DeleteProduct deletes a product from the database
func DeleteProduct(id int) error {
	i := findIndexByProductID(id)
	if i == -1 {
		return ErrProductNotFound
	}
	productList = append(productList[:i], productList[i+1])
	return nil
}

// findIndex finds the index of a product in the database
// returns -1 when no product can be found
func findIndexByProductID(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}
	return -1
}
