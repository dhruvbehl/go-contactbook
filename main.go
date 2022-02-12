package main

import (
	"flag"

	"github.com/dhruvbehl/address-book/pkg/router"
	"github.com/gin-gonic/gin"
)

func main() {
	address := flag.String("address", ":8080", "http address of the server")
	flag.Parse()

	engine := gin.Default()

	engine.GET("/contacts", router.GetAllContactHandler)
	engine.GET("/contact/:id", router.GetContactByIdHandler)
	engine.POST("/contact", router.CreateContactHandler)
	engine.PUT("/contact/:id", router.UpdateContactHandler)
	engine.DELETE("/contact/:id", router.DeleteContactByIdHandler)

	engine.Run(*address)
}