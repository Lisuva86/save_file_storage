package main

import (
	"zip_archive/api"
	"zip_archive/config"
	"zip_archive/controller"
	"zip_archive/middleware"
)

func main() {
	controller := controller.New()
	a := api.Init()

	// логин
	login := a.Group("/api/v1")
	api.RegisterLoginHandlers(login, *controller)
	

	// авторизовано
	protected := a.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware())
	api.RegisterSaveFileHandlers(protected, *controller)
	port := config.GetPort()
	a.Run(":" + port)
}
