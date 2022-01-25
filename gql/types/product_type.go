package types

import "github.com/graphql-go/graphql"

type ProductTypeEntity interface {
	GetObject() *graphql.Object
}

type ProductType struct {
	object *graphql.Object
}

func NewProductType() ProductTypeEntity {
	object := graphql.NewObject(graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"info": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Float,
			},
		},
	})
	return &ProductType{
		object: object,
	}
}

func (p *ProductType) GetObject() *graphql.Object {
	return p.object
}
