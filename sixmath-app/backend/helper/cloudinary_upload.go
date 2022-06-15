package helper

import (
	"context"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gofiber/fiber/v2"
	"mime/multipart"
	"os"
	"rpl-sixmath/config"
	"rpl-sixmath/exception"
	"strings"
)

func UploadToCloudinary(ctx *fiber.Ctx, file *multipart.FileHeader, path string, filename string) (*uploader.UploadResult, error) {
	configuration := config.New(".env")
	cld, err := config.NewCloudinary(configuration)
	exception.PanicIfNeeded(err)
	p := strings.Split(path, ":")
	err = ctx.SaveFile(file, p[0]+file.Filename)
	exception.PanicIfNeeded(err)

	res, err := cld.Upload.Upload(context.Background(), p[0]+file.Filename, uploader.UploadParams{
		PublicID: filename,
		Folder:   p[1],
	})

	err = os.Remove(p[0] + file.Filename)
	exception.PanicIfNeeded(err)

	return res, err
}
