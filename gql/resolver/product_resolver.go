package resolver

import (
	"go-graphql/repo"

	"github.com/graphql-go/graphql"
)

type ProductResolverEntity interface {
	ResolveProductListField(graphql.ResolveParams) (interface{}, error)
}
type ProductResolver struct {
	productRepo repo.ProductRepoEntity
}

func NewProductResolver(productRepo repo.ProductRepoEntity) ProductResolverEntity {
	return &ProductResolver{productRepo: productRepo}
}

func (p *ProductResolver) ResolveProductListField(graphql.ResolveParams) (interface{}, error) {
	return p.productRepo.GetAllProduct(), nil
}
