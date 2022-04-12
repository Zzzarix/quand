package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	rep "quand/app/repository"
)

func (h *Handler) updateServer(c *gin.Context) {

}

func (h *Handler) getApiInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Version": h.Config.Api.Version})
}

func (h *Handler) getQuestion(c *gin.Context) {
	c.JSON(http.StatusOK, rep.GetQuestion())
}

func (h *Handler) getTodayQuestion(c *gin.Context) {
	c.HTML(http.StatusOK, "question.html", nil)
}

func (h *Handler) apiDocs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
