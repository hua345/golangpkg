package consul

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/welcome", func(c *gin.Context) {
		name := c.DefaultQuery("name", "Guest")
		c.String(http.StatusOK, "Hello %s", name)
	})
	// default listen and serve on 0.0.0.0:8080
	err := router.Run()
	if err != nil {
		panic(err)
	}
}
