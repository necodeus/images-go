package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"images-go/types"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var connection *sql.DB

// SetDB ustawia połączenie z bazą danych dla pakietu
func SetDB(database *sql.DB) {
	connection = database
}

func Close() {
	connection.Close()
}

func InitDB() error {
	err := godotenv.Load()

	if err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	connection, err = sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	))

	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	err = connection.Ping()

	if err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	log.Println("Connected to database")

	return nil
}

func GetImageFromDB(resourceID string) (types.Image, error) {
	row := connection.QueryRow("SELECT id, type_name, resource_id, mime_type, size FROM images WHERE resource_id = ?", resourceID)

	var image types.Image
	err := row.Scan(&image.ID, &image.TypeName, &image.ResourceID, &image.MimeType, &image.Size)

	return image, err
}

func GetImageTypeFromDB(name string) (types.ImageType, error) {
	row := connection.QueryRow("SELECT name, available_resolutions FROM image_types WHERE name = ?", name)

	var imageType types.ImageType
	var resolutions string
	err := row.Scan(&imageType.Name, &resolutions)

	if err != nil {
		return imageType, err
	}

	err = json.Unmarshal([]byte(resolutions), &imageType.AvailableResolutions)

	return imageType, err
}

func SaveImageToDB(image types.Image) (int64, error) {
	result, err := connection.Exec("INSERT INTO images (type_name, resource_id, mime_type, size) VALUES (?, ?, ?, ?)", image.TypeName, image.ResourceID, image.MimeType, image.Size)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
