package main

import (
	"net/http"

	"github.com/camcast3/SimpleCalculatorApi/arithmetic"

	"github.com/gin-gonic/gin"
)

func main() {
	/*port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}*/

	r := setupRouter()
	//r.Run(":" + port)
	r.Run()
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/summation", findSummation)
	return router
}

func findSummation(c *gin.Context) {
	var sumSet arithmetic.Summation

	if err := c.BindJSON(&sumSet); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "field validation failed"})
		return
	}

	sum := Add(sumSet.Value1, sumSet.Value2)
	c.IndentedJSON(http.StatusOK, sum)
}

func Add(x float64, y float64) float64 {
	return x + y
}
