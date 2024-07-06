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

	// print image
	fmt.Println(image)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Image not found:", resourceID)
			c.Writer.WriteHeader(http.StatusNotFound)
			c.File("./errors/404_not_found.webp")
		} else {
			fmt.Println("Error getting image from DB:", err)
			c.Writer.WriteHeader(http.StatusInternalServerError)
			c.File("./errors/500_internal_error.webp")
		}
		return
	}

	uploadsDir := os.Getenv("UPLOADS_DIR")
	filePath := filepath.Join(uploadsDir, resourceID)

	if resolution == "" {
		fmt.Println("Serving image without resizing:", filePath)
		c.File(filePath)
		return
	}

	width, height, err := utils.ParseResolution(resolution)

	if err != nil {
		fmt.Println("Error parsing resolution:", err)
		c.Writer.WriteHeader(http.StatusBadRequest)
		c.File("./errors/400_invalid_resolution.webp")
		return
	}

	imageType, err := database.GetImageTypeFromDB(image.TypeName)

	if err != nil {
		fmt.Println("Error getting image type from DB:", err)
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.File("./errors/500_internal_error.webp")
		return
	}

	if !utils.IsValidResolution(imageType.AvailableResolutions, width, height) {
		fmt.Println("Invalid resolution:", resolution)
		c.Writer.WriteHeader(http.StatusBadRequest)
		c.File("./errors/400_invalid_resolution.webp")
		return
	}

	resizedFilePath := fmt.Sprintf("%s_%dx%d", filePath, width, height)

	if _, err := os.Stat(resizedFilePath); os.IsNotExist(err) {
		fmt.Println("Generating thumbnail:", resizedFilePath)

		parts := strings.Split(image.TypeName, "/")

		fmt.Println("parts", parts)

		if len(parts) < 2 {
			fmt.Println("Invalid image type format:", image.TypeName)
			c.Writer.WriteHeader(http.StatusInternalServerError)
			c.File("./errors/500_internal_error.webp")
			return
		}

		fileFormat := parts[1]

		if err := utils.GenerateThumbnail(filePath, resizedFilePath+"."+fileFormat, width, height); err != nil {
			c.Writer.WriteHeader(http.StatusInternalServerError)
			c.File("./errors/500_internal_error.webp")
			fmt.Println("Error generating thumbnail:", err)
			return
		}
	}

	c.File(resizedFilePath)
}
