package controller

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"zip_archive/config"
)

func (c *Controller) SaveFile(env, subDir, filename string, file io.Reader) (string, error) {
	// Проверка окружения
	if env != config.EnvDev && env != config.EnvProd {
		return "", fmt.Errorf("invalid environment: %s", env)
	}

	// Путь на диске: ./storage/dev/subDir/
	fullPath := filepath.Join(config.StorageRoot, env, subDir)

	// Создаём папку, если её нет
	err := c.CreateFolder(fullPath)
	if err != nil {
		return "", err
	}

	// Полный путь к файлу
	filePath := filepath.Join(fullPath, filename)

	// Создаём файл
	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer dst.Close()

	// Копируем содержимое
	if _, err := io.Copy(dst, file); err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	// Формируем URL: используем только /, добавляем /storage
	urlPath := path.Join(config.StorageRoot, env, subDir, filename)
	return config.URL + urlPath, nil
}
