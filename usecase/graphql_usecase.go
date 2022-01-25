package usecase

import (
	"go-graphql/gql"

	"github.com/graphql-go/graphql"
)

type GraphqlUsecaseEntity interface {
	GraphqlDoQuery(query string) *graphql.Result
}

type GraphqlUsecase struct {
	graphql gql.GraphqlEntity
}

func NewGraphqlUsecase(graphql gql.GraphqlEntity) GraphqlUsecaseEntity {
	return &GraphqlUsecase{graphql: graphql}
}

func (g *GraphqlUsecase) GraphqlDoQuery(query string) *graphql.Result {
	return g.graphql.ExecuteQuery(query)
}
