package s3

import (
	worker "github.com/JesseleDuran/secure-graph-worker"
	"log"
	"path/filepath"
)

type CrimeFileProvider struct {
	S3Manager worker.FileManager
}

func (p CrimeFileProvider) Fetch() ([]string, error) {
	files := p.S3Manager.AllKeys()
	result := make([]string, 0)
	for _, file := range files {
		log.Println("Downloading file:", file)
		if filepath.Ext(file) == ".csv" && file != "output1.csv" {
			_, err := p.S3Manager.Download(file, file)
			if err != nil {
				log.Printf("couldnt download file %s, err:", file)
				continue
			}
			result = append(result, "downloads/"+file)
		}
	}
	return result, nil
}
