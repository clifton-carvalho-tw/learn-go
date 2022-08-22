package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/postAndRedirect", postAndRedirect)

	router.Run("localhost:8080")
}

// postAlbums adds an album from JSON received in the request body.
func postAndRedirect(c *gin.Context) {
	fmt.Println(c.PostForm("encryptedData"))
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(c.PostForm("encryptedData")+" This is awesome!!!!"))
}
