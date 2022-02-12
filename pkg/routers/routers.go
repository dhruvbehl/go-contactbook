package routers

import (
	"fmt"

	_ "github.com/dhruvbehl/address-book/docs"
	"github.com/dhruvbehl/address-book/pkg/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup(address *string) *gin.Engine {
	engine := gin.Default()
	url := ginSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", *address))
  	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	engine.GET("/contacts", handlers.GetAllContactHandler)
	engine.GET("/contact/:id", handlers.GetContactByIdHandler)
	engine.POST("/contact", handlers.CreateContactHandler)
	engine.PUT("/contact/:id", handlers.UpdateContactHandler)
	engine.DELETE("/contact/:id", handlers.DeleteContactByIdHandler)

	return engine
}