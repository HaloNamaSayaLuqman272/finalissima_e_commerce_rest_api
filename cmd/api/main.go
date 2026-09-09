package main

import (
	"finalissima_e_commerce_rest_api/api/middlewares"
	"finalissima_e_commerce_rest_api/database/drivers"
	"finalissima_e_commerce_rest_api/package/constant"
	"finalissima_e_commerce_rest_api/package/fileupload"
	"finalissima_e_commerce_rest_api/package/utils"
	"log"
	"strconv"
)

// dalam file "main.go" kita akan jadikan menjadi tempat penghubungan ke
// database "GetConfigurance", ingat kita hanya perlu melakukan penghubungan
// ini hanya sekali
func main() {
	dbConfig := drivers.DBConfig{
		Username: utils.GetConfigurance(constant.DB_USERNAME),
		Password: utils.GetConfigurance(constant.DB_PASSWORD),
		Database: utils.GetConfigurance(constant.DB_NAME),
		Host:     utils.GetConfigurance(constant.DB_HOST),
		Port:     utils.GetConfigurance(constant.DB_PORT),
	}

	cloudinaryConfig := fileupload.CloudinaryConfig{
		CloudinaryURL: utils.GetConfigurance(constant.CLOUDINARY_URL),
	}

	expireDuration, err := strconv.Atoi(utils.GetConfigurance(constant.JWT_EXPIRE_DURATION))
	if err != nil {
		log.Fatalf("error when parsing expire duration: %v\n", err)
	}

	jwtConfig := middlewares.JWTConfig{
		SecretKey:      utils.GetConfigurance(constant.JWT_SECRET_KEY),
		ExpireDuration: expireDuration,
	}

	var (
		repository = dbConfig.InitDB()
		cloudinary = cloudinaryConfig.InitCloudinary()
		e = 
	)
}
