package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tkanos/gonfig"
	"log"
	. "quand/app"
)

func main() {
	gin.DisableConsoleColor()
	// gin.SetMode(gin.ReleaseMode)

	config := Config{}
	err := gonfig.GetConf("configs/config.json", &config)
	if err != nil {
		panic("Invalid config provided!: " + err.Error())
	}

	app := new(App)
	for {
		func() {
			defer func() {
				if err := recover(); err != nil {
					log.Print("Server was successfully restarted")
				}
			}()
			if err := app.Run(config); err != nil {
				log.Fatal(err)
			}
		}()
	}
}
