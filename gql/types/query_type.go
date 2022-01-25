package types

import (
	"go-graphql/gql/resolver"

	"github.com/graphql-go/graphql"
)

type QueryTypeEntity interface {
	GetObject() *graphql.Object
}

type QueryType struct {
	object *graphql.Object
}

func NewQueryType(productType ProductTypeEntity, productResolver resolver.ProductResolverEntity) QueryTypeEntity {
	object := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"product": &graphql.Field{
				Type:        productType.GetObject(),
				Description: "Get product by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: productResolver.ResolveProductDetailField,
			},
			"productList": &graphql.Field{
				Type:        graphql.NewList(productType.GetObject()),
				Description: "Get product list",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: productResolver.ResolveProductListField,
			},
		},
	})
	return &QueryType{
		object: object,
	}
}

func (p *QueryType) GetObject() *graphql.Object {
	return p.object
}
