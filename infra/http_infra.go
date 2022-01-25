package infra

import "github.com/gin-gonic/gin"

type HttpInfraEntity interface {
	GetHttpEngine() *gin.Engine
	RunEngine()
}

type HttpInfra struct {
	httpEngine *gin.Engine
}

func NewHttpInfra() HttpInfraEntity {
	ginEngine := gin.Default()
	return &HttpInfra{httpEngine: ginEngine}
}

func (h *HttpInfra) GetHttpEngine() *gin.Engine {
	return h.httpEngine
}

func (h *HttpInfra) RunEngine() {
	h.httpEngine.Run()
}
