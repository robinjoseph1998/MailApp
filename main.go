package main

import (
	"MailAoo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.SetUpRputes(router)
	router.Run(":8080")

}
