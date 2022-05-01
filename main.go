package main

import (
	"fmt"
	"log"
)

func main() {
	config, err := ReadConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("%+v\n", config)

	dbconfig := config.Database
	fmt.Printf("%+v\n", dbconfig)

	apiconfig := config.API
	fmt.Printf("%+v\n", apiconfig)
}
