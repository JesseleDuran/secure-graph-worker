package crime

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	graph "github.com/JesseleDuran/gograph"
)

type Crime struct {
	Date     time.Time
	ID       int
	Lat, Lng float64
}

type Crimes struct {
	tree Tree
}

func IndexCrimes(crimes []Crime) Crimes {
	return Crimes{tree: MakeTree(crimes)}
}

func FromCSVFile(path string) []Crime {
	crimes := make([]Crime, 0)
	f, _ := os.Open(path)
	r := csv.NewReader(f)
	r.FieldsPerRecord = -1
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("error FromCSVFile: ", err.Error(), path)
			continue
		}
		crime, err := fromCsvValues(record)
		if err == nil {
			crimes = append(crimes, crime)
		}
	}
	return crimes
}

func fromCsvValues(record []string) (Crime, error) {
	values := strings.Split(record[0], ";")
	if len(values) >= 20 {
		t, _ := time.Parse("2006-01-02 15:04:05", values[0])
		lat, err := strconv.ParseFloat(values[2], 32)
		if err != nil || math.IsNaN(lat) {
			return Crime{}, fmt.Errorf("invalid lat")
		}
		lng, err := strconv.ParseFloat(values[3], 32)
		if err != nil || math.IsNaN(lng) {
			return Crime{}, fmt.Errorf("invalid lng")
		}
		if lat == 0 || lng == 0 {
			return Crime{}, fmt.Errorf("invalid lat and lng")
		}
		return Crime{
			Date: t,
			ID:   0,
			Lat:  lat,
			Lng:  lng,
		}, nil
	}
	return Crime{}, fmt.Errorf("not enough values")
}

// SetWeight sets the Weight between two nodes of the map graph.
// This function is going to be passed to the graph for using it in the
// construction.
func (cc Crimes) SetWeight(a, b graph.Coordinate) float32 {
	radius := float64(100)
	crimes := cc.tree.Search(a.Lat, a.Lng, radius)
	crimes1 := cc.tree.Search(b.Lat, b.Lng, radius)
	return float32(len(crimes) + len(crimes1))
}
