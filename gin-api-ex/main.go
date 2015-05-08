package main

import (
	"github.com/daviskoh/go-playgrounds/gin-api-ex/healthcheck"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/healthcheck", healthcheck.Index)

	router.Run(":3000")
}
