package containers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func checkBody(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
}

func GetNameTest(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname")
	c.JSON(http.StatusOK, gin.H{
		"result":  http.StatusOK,
		"message": fmt.Sprintf("Hello %s %s", firstname, lastname),
	})
}

func GetLog(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname")
	c.JSON(http.StatusOK, gin.H{
		"result":  http.StatusOK,
		"message": fmt.Sprintf("Hello %s %s", firstname, lastname),
	})
}
