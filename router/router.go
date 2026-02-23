package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ai-api/ai"
	"ai-api/data"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/api/status", func(c *gin.Context) { ServerStatus(c) })
	router.POST("/api/v1/text", func(c *gin.Context) { RequestText(c) })
}

func ServerStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "life",
	})
}

func RequestText(c *gin.Context) {
    var req data.Request

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var response string
    var err error

    switch req.Provider {
    case "gemini":
        response, err = ai.RequestToGemini(&req)
    case "gigachat":
        response, err = ai.RequestToGigaChat(&req)
    default:
        c.JSON(http.StatusBadRequest, gin.H{"error": "unknown provider"})
        return
    }

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"response": response})
}
