package resolver

import (
	"go-graphql/model"
	"go-graphql/repo"
	"log"
	"math/rand"

	"github.com/graphql-go/graphql"
)

type ProductResolverEntity interface {
	ResolveProductListField(p graphql.ResolveParams) (interface{}, error)
	ResolveProductDetailField(p graphql.ResolveParams) (interface{}, error)
	ResolveCreateProduct(gp graphql.ResolveParams) (interface{}, error)
	ResolveUpdateProduct(gp graphql.ResolveParams) (interface{}, error)
	ResolveDeleteProduct(gp graphql.ResolveParams) (interface{}, error)
}
type ProductResolver struct {
	productRepo repo.ProductRepoEntity
}

func NewProductResolver(productRepo repo.ProductRepoEntity) ProductResolverEntity {
	return &ProductResolver{productRepo: productRepo}
}

func (p *ProductResolver) ResolveProductListField(gp graphql.ResolveParams) (interface{}, error) {
	return p.productRepo.GetAllProduct(), nil
}

func (p *ProductResolver) ResolveProductDetailField(gp graphql.ResolveParams) (interface{}, error) {
	id, ok := gp.Args["id"].(int)
	if ok {
		return p.productRepo.GetDetailProduct(id), nil
	}
	return nil, nil
}

func (p *ProductResolver) ResolveCreateProduct(gp graphql.ResolveParams) (interface{}, error) {
	product := model.Product{
		ID:    int64(rand.Intn(100000)),
		Name:  gp.Args["name"].(string),
		Info:  gp.Args["info"].(string),
		Price: gp.Args["price"].(float64),
	}
	return p.productRepo.CreateProduct(product), nil
}

func (p *ProductResolver) ResolveUpdateProduct(gp graphql.ResolveParams) (interface{}, error) {
	product := model.Product{}
	id, _ := gp.Args["id"].(int)
	product.Name, _ = gp.Args["name"].(string)
	product.Info, _ = gp.Args["info"].(string)
	product.Price, _ = gp.Args["price"].(float64)
	log.Println(product)
	return p.productRepo.UpdateProduct(product, id), nil
}

func (p *ProductResolver) ResolveDeleteProduct(gp graphql.ResolveParams) (interface{}, error) {
	id, _ := gp.Args["id"].(int)
	return p.productRepo.DeleteProduct(id), nil
}
