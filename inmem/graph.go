package inmem

import (
	"fmt"
	osm "github.com/JesseleDuran/gograph/osm/pbf"
	worker "github.com/JesseleDuran/secure-graph-worker"
	"github.com/JesseleDuran/secure-graph-worker/config"
	"log"
)

type GraphCreator struct {
	S3Manager worker.FileManager
}

func (c GraphCreator) Create(filter osm.Filter, content string) error {
	graph := osm.MakeGraphFromFile(filter)
	graphName := fmt.Sprintf("%s-%s-%s.gob", content, filter.Mode.ToString(), config.Config.Country)
	errGraph := graph.Serialize(graphName)
	if errGraph != nil {
		return fmt.Errorf("[Create][serialize][err: %w]", errGraph)
	}

	//Upload .gob to S3 bucket
	log.Print("Uploading graph to S3")
	errUpload := c.S3Manager.Upload(fmt.Sprintf("%s", graphName), graphName)
	if errUpload != nil {
		return fmt.Errorf("[Create][serialize][err: %w]", errUpload)
	}
	return nil
}
