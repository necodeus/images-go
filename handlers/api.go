package handlers

import (
	"images-go/database"
	"images-go/types"
	"images-go/utils"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadImage(c *gin.Context) {
	typeName := c.PostForm("type_name")
	file, err := c.FormFile("file")

	if err != nil || typeName == "" || !utils.GetExtension(file.Filename) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	resourceID := uuid.New().String()
	uploadsDir := os.Getenv("UPLOADS_DIR")
	filePath := filepath.Join(uploadsDir, resourceID)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	mimeType := file.Header.Get("Content-Type")
	fileInfo, err := os.Stat(filePath)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get file info"})
		return
	}

	newImage := types.Image{TypeName: typeName, ResourceID: resourceID, MimeType: mimeType, Size: fileInfo.Size()}
	inserts, err := database.SaveImageToDB(newImage)

	if err != nil || inserts == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image info"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"resource_id": resourceID})
}

// func getImageTypes(c *gin.Context) {
// 	query := "SELECT name, available_resolutions FROM image_types"
// 	rows, err := db.Query(query)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve image types"})
// 		return
// 	}

// 	defer rows.Close()

// 	var imageTypes []ImageType
// 	for rows.Next() {
// 		var imageType ImageType
// 		var resolutions string

// 		if err := rows.Scan(&imageType.Name, &resolutions); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan image type"})
// 			return
// 		}

// 		if err := json.Unmarshal([]byte(resolutions), &imageType.AvailableResolutions); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal resolutions"})
// 			return
// 		}

// 		imageTypes = append(imageTypes, imageType)
// 	}

// 	c.JSON(http.StatusOK, imageTypes)
// }

// func addImageType(c *gin.Context) {
// 	var imageType ImageType
// 	if err := c.ShouldBindJSON(&imageType); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
// 		return
// 	}

// 	resolutions, err := json.Marshal(imageType.AvailableResolutions)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal resolutions"})
// 		return
// 	}

// 	query := "INSERT INTO image_types (name, available_resolutions) VALUES (?, ?)"
// 	if _, err := db.Exec(query, imageType.Name, resolutions); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image type"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Image type added"})
// }

// func updateImageType(c *gin.Context) {
// 	name := c.Param("name")

// 	var resolutions []string

// 	if err := c.ShouldBindJSON(&resolutions); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
// 		return
// 	}

// 	resolutionsJSON, err := json.Marshal(resolutions)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal resolutions"})
// 		return
// 	}

// 	query := "UPDATE image_types SET available_resolutions = ? WHERE name = ?"

// 	if _, err := db.Exec(query, resolutionsJSON, name); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update image type"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Image type updated"})
// }
