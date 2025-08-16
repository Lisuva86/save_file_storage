package controller

import (
	"fmt"
	"os"
)

func (c *Controller) CreateFolder(fullPath string) error {
	if err := os.MkdirAll(fullPath, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	return nil
}
