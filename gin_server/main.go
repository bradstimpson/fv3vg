package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	filename := filepath.Base(os.Args[0])

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("Serving from "+filename+" model using Golang GIN."))
	})

	r.GET("/predict", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"y": "pong",
		})
	})

	err := r.Run(":8002")

	if err != nil {
		fmt.Errorf("%v", err)
	}
}
