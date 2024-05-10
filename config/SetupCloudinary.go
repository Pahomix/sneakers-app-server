package config

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"os"
)

func SetupCloudinary() (*cloudinary.Cloudinary, error) {
	CloudinaryCloudName := os.Getenv("CLD_NAME")
	CloudinaryApiKey := os.Getenv("API_KEY")
	CloudinaryApiSecret := os.Getenv("API_SECRET")

	cld, err := cloudinary.NewFromParams(CloudinaryCloudName, CloudinaryApiKey, CloudinaryApiSecret)
	if err != nil {
		return nil, err
	}

	return cld, nil
}
