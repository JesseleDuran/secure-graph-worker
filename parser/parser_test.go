package parser

import (
	"fmt"
	"log"
	"osm-graph-parser/coordinates"
	"osm-graph-parser/graph"
	"osm-graph-parser/graph/shortest_path/dijkstra"
	"testing"
)

func TestFromOSMFileV2(t *testing.T) {
	g, _ := FromOSMFileV2("test-w-w-w.osm")
	rr, _ := graph.RestrictionsFromFile("test-w-w-w.osm", g)
	g.AddRestrictions(rr)
	log.Println(len(g.Nodes))
	for _, n := range g.Nodes {
		log.Println(n.ID, n.Edges)
	}
}

func TestFromOSMFile(t *testing.T) {
	g, _ := FromOSMFileV2("testdata/Bogota.osm")
	restrictions, err := graph.RestrictionsFromFile("testdata/Bogota.osm", g)
	if err != nil {
		t.Fatal(err)
	}
	log.Println("restrictions", len(restrictions))
	g.AddRestrictions(restrictions)
	log.Println("starting dijsktra")
	d := dijkstra.DijkstraV2{Graph: g}
	r := d.FromCoordinates(coordinates.Coordinates{
		Lat: 4.649383,
		Lng: -74.106763,
	}, coordinates.Coordinates{
		Lat: 4.617690,
		Lng: -74.145522,
	})
	log.Println("termino", r.TotalWeight, len(r.Polyline))
	for _, c := range r.Polyline {
		fmt.Println("[", c[0], ",", c[1], "],")
	}
}

func TestRestrictions(t *testing.T) {
	tests := []struct {
		name             string
		graphFile        string
		originID         int64
		destinyID        int64
		expectedLenNodes int
		restrictionFile  string
	}{
		{
			name:             "Diff from multiple restrictions",
			graphFile:        "testdata/diff_from_multiple_restrictions.osm",
			originID:         1,
			destinyID:        9,
			expectedLenNodes: 8,
			restrictionFile:  "testdata/diff_from_multiple_restrictions.osm",
		},
		{
			name:             "Diff from multiple restrictions",
			graphFile:        "testdata/diff_from_multiple_restrictions.osm",
			originID:         4,
			destinyID:        8,
			expectedLenNodes: 3,
			restrictionFile:  "testdata/diff_from_multiple_restrictions.osm",
		},
		{
			name:             "Simple way way way",
			graphFile:        "testdata/simple_way_way_way.osm",
			originID:         1,
			destinyID:        6,
			expectedLenNodes: 7,
			restrictionFile:  "testdata/simple_way_way_way.osm",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := FromOSMFileV2(tt.graphFile)
			if err != nil {
				t.Fatal(err)
			}
			restrictions, err := graph.RestrictionsFromFile(tt.restrictionFile, g)
			if err != nil {
				t.Fatal(err)
			}
			g.AddRestrictions(restrictions)
			for _, n := range g.Nodes {
				log.Println(n.ID, n.Edges)
			}
			d := dijkstra.DijkstraV2{Graph: g}

			_, got := d.FromNodeIDs(tt.originID, tt.destinyID)
			log.Println(got)
			if len(got) != tt.expectedLenNodes {
				t.Fatalf("expect: %d, got: %d", tt.expectedLenNodes, len(got))
			}

		})
	}

}
