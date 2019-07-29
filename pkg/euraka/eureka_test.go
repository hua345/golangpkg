package euraka

import (
	"github.com/gin-gonic/gin"
	"github.com/hudl/fargo"
	"net/http"
	"testing"
	"time"
)

func TestEureka(t *testing.T) {
	c := fargo.NewConn("http://127.0.0.1:8000")
	apps, err := c.GetApps()
	if err != nil {
		panic(err)
	}
	t.Log(apps)
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
	err = router.Run(":8081")
	if err != nil {
		panic(err)
	}

	time.Sleep(100 * time.Second)
}
