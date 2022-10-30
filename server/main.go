package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const (
	addr = ":3000"
)

func main() {
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("../client/", true)))

	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "localhost:3000/api works",
			})
		})
	}

	log.Printf("server running at %s \n", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Error running server at %s \n %+v \n",
			addr,
			err)
	}
}
