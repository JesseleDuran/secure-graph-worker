package main

import (
  "log"
  "osm-graph-parser/file/json"
  "osm-graph-parser/parser"
)

func main() {
  g, err := parser.FromOSMFile("downloads/medellin.osm")
  log.Println(len(g.Edges2), err)
  json.Write("osm-graph-medellin-name-17.json", g.Edges2)
}
