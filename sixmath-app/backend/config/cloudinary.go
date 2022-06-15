package config

import (
	"github.com/cloudinary/cloudinary-go"
)

func NewCloudinary(cfg Config) (*cloudinary.Cloudinary, error) {

	return cloudinary.NewFromParams(cfg.Get("CLOUDINARY_NAME"), cfg.Get("CLOUDINARY_API_KEY"), cfg.Get("CLOUDINARY_API_SECRET"))
}
