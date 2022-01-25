package types

import (
	"go-graphql/gql/resolver"

	"github.com/graphql-go/graphql"
)

type MutationTypeEntity interface {
	GetObject() *graphql.Object
}

type MutationType struct {
	object *graphql.Object
}

func NewMutationType(productType ProductTypeEntity, productResolver resolver.ProductResolverEntity) MutationTypeEntity {
	object := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"create": &graphql.Field{
				Type:        productType.GetObject(),
				Description: "Create new product",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"info": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Float),
					},
				},
				Resolve: productResolver.ResolveCreateProduct,
			},
		},
	})
	return &MutationType{
		object: object,
	}
}

func (p *MutationType) GetObject() *graphql.Object {
	return p.object
}
