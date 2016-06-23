package main

import (
	"log"

	"github.com/black-banana/bee-hive/questions"
	"github.com/black-banana/bee-hive/rethink"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func main() {
	// LoadConfiguration from config.json
	LoadConfiguration()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api")
	questions.New(api)

	rethink.InitMasterSession(globalConfig.dbServer, globalConfig.dbName)
	defer rethink.CloseMasterSession()

	//go hub.Run()

	//e.GET("/ws", standard.WrapHandler(http.HandlerFunc(hub.ServeHub())))

	routes := e.Routes()
	for _, route := range routes {
		log.Println(route.Method, route.Path)
	}

	log.Println("Started with", globalConfig.listen)
	e.Run(standard.New(globalConfig.listen))
}
