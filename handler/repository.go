package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/thalerngsak/git-scanner/service"
	"net/http"
)

type repositoryHandler struct {
	repSrv service.RepositoryService
}

func NewRepositoryHandler(repSrv service.RepositoryService) repositoryHandler {
	return repositoryHandler{repSrv: repSrv}
}

func (h repositoryHandler) NewRepository(c *gin.Context) {

	var r service.RepositoryRequest
	if err := c.BindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.repSrv.NewRepository(r)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, response)
}

func (h repositoryHandler) GetRepository(c *gin.Context) {
	repos, err := h.repSrv.GetRepository()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, repos)
}

func (h repositoryHandler) GetRepositoryByID(c *gin.Context) {
	id := c.Param("id")
	r, err := h.repSrv.GetRepositoryByID(id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	if r == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, r)
}

func (h repositoryHandler) UpdateRepository(c *gin.Context) {
	id := c.Param("id")
	var r *service.RepositoryRequest
	if err := c.BindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repSrv.UpdateRepository(id, r); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, r)
}

func (h repositoryHandler) DeleteRepository(c *gin.Context) {
	id := c.Param("id")
	if err := h.repSrv.DeleteRepository(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
