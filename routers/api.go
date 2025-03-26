package routers

import (
    "net/http"
    "tree-visualization/models"
    "tree-visualization/utils"
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.POST("/generate", func(c *gin.Context) {
        var request struct {
            Nodes  []models.Node `json:"nodes"`
            Format string        `json:"format" binding:"required"` // e.g., "png", "svg"
        }

        if err := c.ShouldBindJSON(&request); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // Generate DOT content
        dotContent := utils.GenerateDOT(request.Nodes)
        
        // Generate image
        imgBytes, err := utils.GenerateImage(dotContent, request.Format)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        // Set content type based on format
        contentType := "image/" + request.Format
        c.Data(http.StatusOK, contentType, imgBytes)
    })

    return r
}