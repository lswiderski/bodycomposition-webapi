package main

import (
    "net/http"
	"github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
	bc "github.com/lswiderski/bodycomposition-webapi"
)

func getInfo(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"name": "bodycomposition-webapi",
		"projectUrl": "https://github.com/lswiderski/bodycomposition-webapi",
		"requestEndpoint": "/upload"})
}

func postBodycomposition(c *gin.Context) {
	var request bc.BodyCompositionRequest
	if err := c.BindJSON(&request); err != nil {
        return
    }
	result := bc.UploadBodyComposition(request)
	if len(result) > 0 { 
		c.String(http.StatusBadRequest, result)
	 } else {
		c.String(http.StatusCreated, result)
	}
	
}

func main() {
	router :=gin.Default()
	router.Use(cors.Default())
	router.GET("/",getInfo)
	router.POST("/upload", postBodycomposition)
	router.Run(":8080")
}