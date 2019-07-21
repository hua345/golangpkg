package consul

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// default listen and serve on 0.0.0.0:8080
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
