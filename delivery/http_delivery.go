package delivery

import (
	"go-graphql/infra"
	"go-graphql/model"
	"go-graphql/usecase"
	"go-graphql/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpDeliveryEntity interface {
	RunHttpDelivery()
}

type HttpDelivery struct {
	httpInfra      infra.HttpInfraEntity
	graphqlUseCase usecase.GraphqlUsecaseEntity
}

func NewHttpDelivery(
	httpInfra infra.HttpInfraEntity,
	graphqlUseCase usecase.GraphqlUsecaseEntity,
) HttpDeliveryEntity {
	httpDelivery := &HttpDelivery{httpInfra: httpInfra, graphqlUseCase: graphqlUseCase}
	api := httpInfra.GetHttpEngine().Group("/api")
	api.POST("/graphql", httpDelivery.HandleGraphql)
	return httpDelivery
}

func (h *HttpDelivery) HandleGraphql(c *gin.Context) {
	var gqlReq model.GraphqlRequest
	err := c.ShouldBindJSON(&gqlReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.CreateHTTPRespond(http.StatusBadRequest, "Bad Request", err.Error()))
		return
	}
	result := h.graphqlUseCase.GraphqlDoQuery(gqlReq.Query)
	if result.HasErrors() {
		c.JSON(http.StatusBadRequest, utils.CreateHTTPRespond(http.StatusBadRequest, "Bad Request", result.Errors))
		return
	}
	c.JSON(http.StatusOK, utils.CreateHTTPRespond(http.StatusOK, "ok", result.Data))
}

func (h *HttpDelivery) RunHttpDelivery() {
	h.httpInfra.RunEngine()
}
