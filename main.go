package main

import (
	"log"
	"urlShorter/config"
	"urlShorter/handler"
	"urlShorter/storage/redis"
	
	"fmt"
	"github.com/valyala/fasthttp"
)

func main() {
	configuration, err := config.FromFile("./configuration.json")
	if err != nil {
		log.Fatal(err)
	}
	

	service, err := redis.New(configuration.Redis.Host, configuration.Redis.Port, configuration.Redis.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer service.Close()
	fmt.Println("Hello World")
	
	router := handler.New(configuration.Options.Schema, configuration.Options.Prefix, service)
	fmt.Println("Hello World")
	log.Print("Hello Pig")
	log.Fatal(fasthttp.ListenAndServe(":"+configuration.Server.Port, router.Handler))
	fmt.Println("Hello World")
}
