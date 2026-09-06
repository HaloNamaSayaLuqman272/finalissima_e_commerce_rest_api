package fileupload

import (
	"context"
	"log"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

type CloudinaryConfig struct {
	CloudinaryURL string
}

type Uploader interface {
	UploadFile(ctx context.Context, file any) (string, error)
}

type CloudinaryUploader struct {
	Cld     *cloudinary.Cloudinary
	File    any
	Options uploader.UploadParams
}

func (c *CloudinaryConfig) InitCloudinary() *cloudinary.Cloudinary {
	// kita disini tidak menambahkan parameter input
	// karena memang tidak dibutuhkan untuk keadaan fungsi ini
	// fungsi ini kita hanya untuk meng-koneksikan aplikasi Go ini
	// ke API Cloudinary
	cld, err := cloudinary.NewFromURL(c.CloudinaryURL)
	//code ini adalah baris untuk membuat koneksi ke layanan Cloudinary
	if err != nil {
		log.Fatalf("error when connceting to the cloudinary: %s\n", err)
	}

	log.Println("cloudinary initialization succed")
	return cld
}

func (c *CloudinaryUploader) UploadFile(ctx context.Context, file any) (string, error) {
	// code ini untuk kebutuhan upload file ke Cloudinary
	resp, err := c.Cld.Upload.Upload(ctx, file, c.Options)
	if err != nil {
		return "", err
	}

	return resp.SecureURL, nil
}
