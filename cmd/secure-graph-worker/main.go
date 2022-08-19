package main

import (
	osm "github.com/JesseleDuran/gograph/osm/pbf"
	"github.com/JesseleDuran/secure-graph-worker/config"
	"github.com/JesseleDuran/secure-graph-worker/inmem"
	"github.com/JesseleDuran/secure-graph-worker/inmem/crime"
	"github.com/JesseleDuran/secure-graph-worker/s3"
	"github.com/golang/geo/s2"
	"github.com/robfig/cron"
	"log"
	"os"
	"os/signal"
)

func main() {
	s3Manager := s3.NewS3Manager(s3.GetClient())
	// Initialize configs
	config.Initialize()

	// Init.
	mapsProvider := s3.MapsFileProvider{S3Manager: s3Manager}
	zoneProvider := s3.CrimeFileProvider{S3Manager: s3Manager}
	graphCreator := inmem.GraphCreator{S3Manager: s3Manager}

	c := cron.New()
	c.AddFunc("@midnight", func() {
		start(mapsProvider, zoneProvider, graphCreator)
	})
	go c.Start()
	start(mapsProvider, zoneProvider, graphCreator)

	// Handle graceful shutdowns via SIGINT
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

}
func start(mapsProvider s3.MapsFileProvider, crimeProvider s3.CrimeFileProvider, graph inmem.GraphCreator) {
	log.Print("Downloading files from S3.")
	osmPbfFile, _ := mapsProvider.Fetch()
	crimesFiles, _ := crimeProvider.Fetch()
	crimes := make([]crime.Crime, 0)
	points := make([]s2.Point, 0)

	for _, p := range crimesFiles {
		crimes = append(crimes, crime.FromCSVFile(p)...)
	}
	for _, c := range crimes {
		log.Println(c.Lat, c.Lng)
		points = append(points, s2.PointFromLatLng(s2.LatLngFromDegrees(c.Lat, c.Lng)))
	}
	coverage := s2.LoopFromPoints(points)
	tree := crime.IndexCrimes(crimes)

	log.Print("Generating driving distance graph.")
	err := graph.Create(osm.Filter{
		Path:      osmPbfFile,
		Mode:      osm.Driving,
		Coverage:  *coverage,
		SetWeight: nil,
	}, "distance")
	if err != nil {
		log.Println("err Generating driving distance graph.", err.Error())
	}

	log.Print("Generating driving crimes graph.")
	err = graph.Create(osm.Filter{
		Path:      osmPbfFile,
		Mode:      osm.Driving,
		Coverage:  *coverage,
		SetWeight: tree.SetWeight,
	}, "crimes")
	if err != nil {
		log.Println("err Generating driving crimes graph.", err.Error())
	}

	log.Print("Generating Cycling distance graph.")
	err = graph.Create(osm.Filter{
		Path:      osmPbfFile,
		Mode:      osm.Cycling,
		Coverage:  *coverage,
		SetWeight: nil,
	}, "distance")
	if err != nil {
		log.Println("err Generating Cycling distance graph.", err.Error())
	}

	log.Print("Generating Cycling crimes graph.")
	err = graph.Create(osm.Filter{
		Path:      osmPbfFile,
		Mode:      osm.Cycling,
		Coverage:  *coverage,
		SetWeight: tree.SetWeight,
	}, "crimes")
	if err != nil {
		log.Println("err Generating Cycling crimes graph.", err.Error())
	}
}
