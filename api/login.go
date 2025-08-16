package api

import (
	"github.com/gin-gonic/gin"
)

func (h *handlers) postLoginHandler(c *gin.Context) {
	h.controller.Login(c)
}
