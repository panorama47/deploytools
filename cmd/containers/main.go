package main

import (
	"github.com/gin-gonic/gin"
	klog "k8s.io/klog/v2"
	"net/http"
)

type myForm struct {
	Colors []string `form:"colors[]"`
}

func start() error {
	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1.GET("/getLog", routers.GetNameTest(router))
	err := router.Run(":3050")
	if err != nil {
		return err
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	return nil
}

func main() {
	if err := start(); err != nil {
		klog.Fatal(err)
	}
}
