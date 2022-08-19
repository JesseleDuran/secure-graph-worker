package s3

import (
	"fmt"
	worker "github.com/JesseleDuran/secure-graph-worker"
	"github.com/JesseleDuran/secure-graph-worker/config"
)

type MapsFileProvider struct {
	S3Manager worker.FileManager
}

func (p MapsFileProvider) Fetch() (string, error) {
	// Get osm.pbf from S3 Bucket
	return p.S3Manager.Download(fmt.Sprintf("%s.osm.pbf", config.Config.Country), fmt.Sprintf("%s.osm.pbf", config.Config.Country))
}
