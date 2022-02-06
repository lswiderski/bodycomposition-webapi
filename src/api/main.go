package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
	bc "github.com/lswiderski/bodycomposition-webapi"
)

func getInfo(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "bodycomposition-webapi"})
}

func postBodycomposition(c *gin.Context) {
	var request bc.BodyCompositionRequest
	if err := c.BindJSON(&request); err != nil {
        return
    }
	result := bc.UploadBodyComposition(request)
	c.String(http.StatusCreated, result)
}

func main() {
	router :=gin.Default()
	router.GET("/",getInfo)
	router.POST("/upload", postBodycomposition)
	router.Run("localhost:4321")
}