package manager

import (
	"go-graphql/gql"
	"go-graphql/gql/resolver"
	"go-graphql/gql/types"
	"go-graphql/repo"
)

type GraphqlManagerEntity interface {
	GetGraphql() gql.GraphqlEntity
}

type GraphqlManager struct {
	graphql gql.GraphqlEntity
}

func NewGraphqlManager(productRepo repo.ProductRepoEntity) GraphqlManagerEntity {
	productResolver := resolver.NewProductResolver(productRepo)
	productType := types.NewProductType()
	queryType := types.NewQueryType(productType, productResolver)
	mutationType := types.NewMutationType(productType, productResolver)
	graphql := gql.NewGraphql(queryType, mutationType)
	return &GraphqlManager{graphql: graphql}
}

func (g *GraphqlManager) GetGraphql() gql.GraphqlEntity {
	return g.graphql
}
