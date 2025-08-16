package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handlers) postSaveFileHandler(c *gin.Context) {
	// 1. Парсим multipart/form-data
	err := c.Request.ParseMultipartForm(32 << 20) // Максимум 32 MB
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse form data",
		})
		return
	}

	// 2. Читаем поля
	env := c.Request.FormValue("env")
	subDir := c.Request.FormValue("sub_dir")

	// 3. Получаем файл по ключу "file"
	file, handler, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No file is received",
		})
		return
	}
	defer file.Close()

	// 4. Проверяем обязательные поля
	if env == "" || subDir == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing required fields: 'env' or 'sub_dir'",
		})
		return
	}

	// 5. Передаём в контроллер
	result, err := h.controller.SaveFile(env, subDir, handler.Filename, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 6. Успех
	c.JSON(http.StatusCreated, result)
}
