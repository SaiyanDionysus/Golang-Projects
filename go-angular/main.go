package main

import (
	"github.com/gin-gonic/gin"
	"path"
	"path/filepath"
)

func main() {
	r := gin.Default()
	r.NoRoute(func(c *gin.Context)) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File(".")
		}
	}
}
