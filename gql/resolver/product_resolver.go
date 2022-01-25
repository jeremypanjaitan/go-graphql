package resolver

import (
	"go-graphql/repo"

	"github.com/graphql-go/graphql"
)

type ProductResolverEntity interface {
	ResolveProductListField(p graphql.ResolveParams) (interface{}, error)
	ResolveProductDetailField(p graphql.ResolveParams) (interface{}, error)
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
