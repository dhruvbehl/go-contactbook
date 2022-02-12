package main

import (
	"flag"
	"github.com/dhruvbehl/address-book/pkg/routers"
	"github.com/rs/zerolog/log"
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

	engine := routers.Setup(address)

	if err := engine.Run(*address); err != nil {
		log.Fatal().Err(err).Msg("unable to start the server")
	}
}
