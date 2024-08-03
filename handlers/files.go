package handlers

import (
	"database/sql"
	"fmt"
	"images-go/database"
	"images-go/utils"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetImage(c *gin.Context) {
	resourceID := c.Param("resource_id")
	resolution := c.Param("resolution")

	image, err := database.GetImageFromDB(resourceID)

	if err != nil {
		if err == sql.ErrNoRows {
			c.Writer.WriteHeader(http.StatusNotFound)
		} else {
			c.Writer.WriteHeader(http.StatusInternalServerError)
		}
		c.File("./dot.gif")
		return
	}

	uploadsDir := os.Getenv("UPLOADS_DIR")
	filePath := filepath.Join(uploadsDir, resourceID)

	if resolution == "" {
		c.File(filePath)
		return
	}

	width, height, err := utils.ParseResolution(resolution)

	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.File("./dot.gif")
		return
	}

	imageType, err := database.GetImageTypeFromDB(image.TypeName)

	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.File("./dot.gif")
		return
	}

	if !utils.IsValidResolution(imageType.AvailableResolutions, width, height) {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.File("./dot.gif")
		return
	}

	resizedFilePath := fmt.Sprintf("%s_%dx%d", filePath, width, height)

	if _, err := os.Stat(resizedFilePath); os.IsNotExist(err) {
		// jest problem z generowaniem miniaturki
		fileFormat := strings.Split(image.MimeType, "/")[1]

		if err := utils.GenerateThumbnail(filePath, resizedFilePath+"."+fileFormat, width, height); err != nil {
			c.Writer.WriteHeader(http.StatusInternalServerError)
			c.File("./dot.gif")
			return
		}
	}

	c.File(resizedFilePath)
}
