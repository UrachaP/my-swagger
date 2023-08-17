package handler

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

// UploadFile
// @Summary Upload a file
// @Produce json
// @Param file formData file true "File to upload"
// @Success 200 {string} string "File uploaded successfully"
// @Router /upload [post]
func UploadFile(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	// Save the uploaded file
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dstPath := filepath.Join("uploads", file.Filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.String(http.StatusOK, "File uploaded successfully")
}
