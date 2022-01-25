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
			"update": &graphql.Field{
				Type:        productType.GetObject(),
				Description: "Update product by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"info": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
				},
				Resolve: productResolver.ResolveUpdateProduct,
			},

			"delete": &graphql.Field{
				Type:        productType.GetObject(),
				Description: "Delete product by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: productResolver.ResolveDeleteProduct,
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
