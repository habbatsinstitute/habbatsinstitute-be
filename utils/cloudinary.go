package utils

import (
	"fmt"
	"institute/config"

	"github.com/cloudinary/cloudinary-go"
)

func CloudinaryInstance(config config.ProgramConfig) *cloudinary.Cloudinary {
	var urlCloudinary = fmt.Sprintf("cloudinary://%s:%s@%s",
		config.CDN_API_Key,
		config.CDN_API_Secret,
		config.CDN_Cloud_Name)

	CDNService, err := cloudinary.NewFromURL(urlCloudinary)
	if err != nil {
		return nil
	}

	return CDNService
}