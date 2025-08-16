package api

import (
	"zip_archive/controller"

	"github.com/gin-gonic/gin"
)

func Init() *API {
	router := gin.Default()

	return &API{router}
}

type handlers struct {
	controller controller.Controller
}

func New(controller controller.Controller) *handlers {
	return &handlers{
		controller: controller,
	}
}
func RegisterLoginHandlers(routerGroup *gin.RouterGroup, controller controller.Controller) {
	h := New(controller)
	//--------------------------------------------------------------------------------login
	{
		login := routerGroup.Group("/login")
		login.POST("", h.postLoginHandler)
	}
}

func RegisterSaveFileHandlers(routerGroup *gin.RouterGroup, controller controller.Controller) {
	h := New(controller)
	//--------------------------------------------------------------------------------task
	{
		save := routerGroup.Group("/save")
		save.POST("", h.postSaveFileHandler)
	}
}
