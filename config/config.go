package config

import (
	"os"
)

var Config config

type config struct {
	Country        string
	S3BucketName   string
	S3DownloadPath string
}

func Initialize() {
	Config.Country = GetCountry()
	Config.S3BucketName = GetS3BucketName()
	Config.S3DownloadPath = GetS3DownloadPath()
}

func GetCountry() string {
	return os.Getenv("COUNTRY")
}

func GetS3BucketName() string {
	return os.Getenv("AWS_BUCKET_NAME")
}

func GetS3DownloadPath() string {
	path := os.Getenv("S3_DOWNLOAD_PATH")
	if path == "" {
		return "downloads/"
	}
	return path
}
