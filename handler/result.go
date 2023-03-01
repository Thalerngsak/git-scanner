package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/thalerngsak/git-scanner/service"
	"net/http"
)

type resultHandler struct {
	resultSrv service.ResultService
}

func NewResultHandler(resultSrv service.ResultService) resultHandler {
	return resultHandler{resultSrv: resultSrv}
}

func (h resultHandler) GetResult(c *gin.Context) {
	results, err := h.resultSrv.GetResult()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, results)
}

func (h resultHandler) GetResultByRepositoryID(c *gin.Context) {
	id := c.Param("id")
	results, err := h.resultSrv.GetResultByRepositoryID(id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, results)
}
