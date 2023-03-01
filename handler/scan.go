package handler

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thalerngsak/git-scanner/service"
	"gopkg.in/src-d/go-git.v4"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type scanHandler struct {
	repSrv service.RepositoryService
	resSrv service.ResultService
}

func NewScanHandler(resultSrv service.RepositoryService, resSrv service.ResultService) scanHandler {
	return scanHandler{repSrv: resultSrv, resSrv: resSrv}
}

func (h scanHandler) Scan(c *gin.Context) {
	var r service.ScanRequest
	if err := c.BindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repositories, err := h.repSrv.GetRepositoryByID(r.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	res := service.ResultRequest{
		ID:            uuid.New().String(),
		Status:        "Queued",
		RepositoryID:  repositories.ID,
		RepositoryURL: repositories.URL,
		EnqueuedAt:    time.Now(),
	}
	if err := h.resSrv.Create(res); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	scanRepository(&res, res.RepositoryURL, h.resSrv)

	response, _ := h.resSrv.GetResultByID(res.ID)

	c.JSON(http.StatusOK, response)
}

func scanRepository(res *service.ResultRequest, url string, resultStore service.ResultService) {
	res.Status = "In Progress"
	res.StartedAt = time.Now()
	if err := resultStore.UpdateResult(res); err != nil {
		log.Println(err)
	}
	// Clone repository
	repoDir := fmt.Sprintf("./repos/%s", res.ID)
	_, err := git.PlainClone(repoDir, false, &git.CloneOptions{
		URL: url,
	})
	if err != nil {
		res.Status = "Failure"
		res.FinishedAt = time.Now()
		res.Findings = append(res.Findings, service.Finding{Category: "Secret", Message: fmt.Sprintf("Failed to clone repository: %s", err)})
		if err := resultStore.UpdateResult(res); err != nil {
			log.Println(err)
		}
		return
	}
	// Scan repository for secrets
	var findings []service.Finding
	err = filepath.Walk(repoDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			if strings.HasSuffix(info.Name(), ".java") {
				file, err := os.Open(path)
				if err != nil {
					return err
				}
				defer file.Close()
				scanner := bufio.NewScanner(file)
				line := 1
				for scanner.Scan() {
					text := scanner.Text()
					if strings.Contains(text, "public_key") || strings.Contains(text, "private_key") {
						findings = append(findings, service.Finding{Category: "Secret", Message: fmt.Sprintf("Found secret in file %s, line %d", file.Name(), line)})
					}
					line++
				}
			}
		}
		return nil
	})

	if err != nil {
		res.Status = "Failure"
		res.FinishedAt = time.Now()
		res.Findings = append(res.Findings, service.Finding{Category: "Secret", Message: fmt.Sprintf("Failed to clone repository: %s", err)})
		if err := resultStore.UpdateResult(res); err != nil {
			log.Println(err)
		}
		return
	}

	res.Status = "Success"
	res.FinishedAt = time.Now()
	res.Findings = findings
	if err := resultStore.UpdateResult(res); err != nil {
		log.Println(err)
	}
}
