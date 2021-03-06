package repo

import (
	"go-graphql/model"
)

var products = []model.Product{
	{
		ID:    1,
		Name:  "Chicha Morada",
		Info:  "Chicha morada is a beverage originated in the Andean regions of Perú but is actually consumed at a national level (wiki)",
		Price: 7.99,
	},
	{
		ID:    2,
		Name:  "Chicha de jora",
		Info:  "Chicha de jora is a corn beer chicha prepared by germinating maize, extracting the malt sugars, boiling the wort, and fermenting it in large vessels (traditionally huge earthenware vats) for several days (wiki)",
		Price: 5.95,
	},
	{
		ID:    3,
		Name:  "Pisco",
		Info:  "Pisco is a colorless or yellowish-to-amber colored brandy produced in winemaking regions of Peru and Chile (wiki)",
		Price: 9.95,
	},
}

type ProductRepoEntity interface {
	GetAllProduct() []model.Product
	GetDetailProduct(id int) model.Product
	CreateProduct(product model.Product) model.Product
	UpdateProduct(product model.Product, id int) model.Product
	DeleteProduct(id int) model.Product
}

type ProductRepo struct{}

func NewProductRepo() ProductRepoEntity {
	return &ProductRepo{}
}

func (p *ProductRepo) GetAllProduct() []model.Product {
	return products
}

func (p *ProductRepo) GetDetailProduct(id int) model.Product {
	for _, product := range products {
		if int(product.ID) == id {
			return product
		}
	}
	return model.Product{}
}

func (p *ProductRepo) CreateProduct(product model.Product) model.Product {
	products = append(products, product)
	return product
}

func (p *ProductRepo) UpdateProduct(product model.Product, id int) model.Product {
	for i, p := range products {
		if int64(id) == p.ID {
			products[i].Name = product.Name
			products[i].Info = product.Info
			products[i].Price = product.Price
			product = products[i]
			break
		}
	}
	product.ID = int64(id)
	return product
}

func (p *ProductRepo) DeleteProduct(id int) model.Product {
	var product model.Product

	for i, p := range products {
		if int64(id) == p.ID {
			product = products[i]
			products = append(products[:i], products[i+1:]...)
		}
	}
	return product
}
