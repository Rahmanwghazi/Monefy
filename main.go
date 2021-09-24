package main

import (
	"github.com/Rahmanwghazi/Monefy/config"
	"github.com/Rahmanwghazi/Monefy/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Start(":8000")
}
