package main

import (
	"flag"

	"github.com/Bronsun/GogoSpace/config"
	"github.com/Bronsun/GogoSpace/server"
)

func main() {
	// environment flag to change env files - we can add custom gloabl variables in different env files
	environment := flag.String("e", "default", "")
	flag.Parse()

	config.Init(*environment)
	server.Init()

}
