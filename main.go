package main

import (
	"log"
	"osm-graph-parser/file/json"
	"osm-graph-parser/parser"
)

func main() {
	g, err := parser.FromOSMFileV2("downloads/Bogota.osm")
	log.Println(len(g.Nodes), err)
	json.Write("osm-graph-v2-bogota.json", g.Nodes)
}
