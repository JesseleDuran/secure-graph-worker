package worker

import osm "github.com/JesseleDuran/gograph/osm/pbf"

//go:generate mockery --name S3Client
type S3Client interface {
	Get(bucketName, objectName, fileName string) error
	Put(bucketName, objectName, filePath string) (int64, error)
	GetAllObjectKeys(bucketName string) []string
}

//go:generate mockery --name S3FileManager
type FileManager interface {
	Download(source, destination string) (string, error)
	Upload(source, destination string) error
	AllKeys() []string
}

type FileProvider interface {
	Fetch() (string, error)
}

type Graph interface {
	Create(filter osm.Filter, content string) error
}
