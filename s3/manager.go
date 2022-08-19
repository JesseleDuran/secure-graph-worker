package s3

import (
	"fmt"
	worker "github.com/JesseleDuran/secure-graph-worker"
	"github.com/JesseleDuran/secure-graph-worker/config"
)

const (
	downloadPath = "downloads/"
)

type Manager struct {
	client worker.S3Client
}

func NewS3Manager(client worker.S3Client) worker.FileManager {
	return Manager{
		client: client,
	}
}

// Download builds a Memory Image from a S3 file snapshot of Stores.
func (s3 Manager) Download(source, destination string) (string, error) {
	destinationPath := downloadPath + destination
	err := s3.client.Get(
		config.GetS3BucketName(),
		source,
		destinationPath,
	)
	if err != nil {
		return "", fmt.Errorf("[S3Manager:Download][s3 downloading][err: %w]", err)
	}
	return destinationPath, nil
}

func (s3 Manager) Upload(source, destination string) error {
	_, err := s3.client.Put(config.GetS3BucketName(), source, destination)
	if err != nil {
		return fmt.Errorf("[S3Manager:Upload][s3 uploading][err: %w]", err)
	}
	return nil
}

func (s3 Manager) AllKeys() []string {
	return s3.client.GetAllObjectKeys(config.GetS3BucketName())
}
