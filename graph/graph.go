package graph

import "osm-graph-parser/edge"

type Graph struct {
  Edges []edge.Edges `json:"edges"`
  Edges2 []edge.Edges2 `json:"edges"`
}
