package gql

import "github.com/graphql-go/graphql"

type GraphqlEntity interface {
	ExecuteQuery(query string) *graphql.Result
}

type Graphql struct {
	graphqlSchema graphql.Schema
}

func NewGraphql(
	queryType *graphql.Object,
) GraphqlEntity {
	graphqlSchema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
	return &Graphql{
		graphqlSchema: graphqlSchema,
	}
}

func (g *Graphql) ExecuteQuery(query string) *graphql.Result {
	return graphql.Do(graphql.Params{
		Schema:        g.graphqlSchema,
		RequestString: query,
	})
}
