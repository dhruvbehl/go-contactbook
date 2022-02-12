package main

import (
	"flag"
	"fmt"

	_ "github.com/dhruvbehl/address-book/docs"
	"github.com/dhruvbehl/address-book/pkg/router"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Contact Book
// @version 1.0
// @description Contact book written in golang with the purpose of learning golang rest framework gin
// @termsOfService http://swagger.io/terms/

// @contact.name Dhruv Behl
// @contact.email dhruvbhl@gmail.com

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	address := flag.String("address", ":8080", "http address of the server")
	flag.Parse()

	engine := gin.Default()
	url := ginSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", *address))
  	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	engine.GET("/contacts", router.GetAllContactHandler)
	engine.GET("/contact/:id", router.GetContactByIdHandler)
	engine.POST("/contact", router.CreateContactHandler)
	engine.PUT("/contact/:id", router.UpdateContactHandler)
	engine.DELETE("/contact/:id", router.DeleteContactByIdHandler)

	if err := engine.Run(*address); err != nil {
		log.Fatal().Err(err).Msg("unable to start the server")
	}
}