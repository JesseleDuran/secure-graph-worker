package graph

import (
	"github.com/golang/geo/s2"
	"github.com/paulmach/osm"
	"github.com/umahmood/haversine"
	"osm-graph-parser/edge"
	"osm-graph-parser/parser/resources"
	"osm-graph-parser/tag"
)

type Graph struct {
	Edges  []edge.Edges  `json:"edges"`
	Edges2 []edge.Edges2 `json:"edges"`
	Edges3 []edge.Edges3 `json:"edges"`
}

type GraphV2 struct {
	Nodes Nodes
}

//can be a cell id or a osmID
type Nodes map[uint64]Node

type Node struct {
	ID       uint64
	Location s2.CellID
	Edges    Edges
	IsVia    bool
	WayID    int
}

type Edge struct {
	Weight float64
}

type Edges map[uint64]Edge

func (g *GraphV2) RelateNodesFromWay(way osm.Way) {
	oneway := false
	auxTags := tag.FromOSMTags(way.Tags)
	v, ok := auxTags["oneway"]
	if ok && v == "yes" {
		oneway = true
	}
	v, ok = auxTags["junction"]
	if ok && v == "roundabout" {
		oneway = true
	}
	for i := 0; i < len(way.Nodes)-1; i++ {

		a := resources.Nodes[way.Nodes[i].ID.FeatureID().Ref()]
		b := resources.Nodes[way.Nodes[i+1].ID.FeatureID().Ref()]
		nodeA := g.FindNodeOrCreate(a)
		nodeB := g.FindNodeOrCreate(b)
		weight := Distance(nodeA.Location, nodeB.Location)

		nodeA.Edges[nodeB.ID] = Edge{Weight: weight}

		if !oneway {
			nodeB.Edges[nodeA.ID] = Edge{Weight: weight}
		}
	}
}

func (g *GraphV2) FindNodeOrCreate(n osm.Node) *Node {
	node, ok := g.Nodes[uint64(n.ID.FeatureID().Ref())]
	if !ok {
		node = Node{
			ID:       uint64(n.ID.FeatureID().Ref()),
			Location: s2.CellID(edge.ToToken(n.Lat, n.Lon)),
			Edges:    make(map[uint64]Edge),
		}
		g.AddNodes(node)
	}
	return &node
}

func (g *GraphV2) AddNodes(nodes ...Node) {
	for _, n := range nodes {
		g.Nodes[n.ID] = n
	}
}

func Distance(a, b s2.CellID) float64 {
	_, km := haversine.Distance(
		haversine.Coord{Lat: a.LatLng().Lat.Degrees(), Lon: a.LatLng().Lng.Degrees()},
		haversine.Coord{Lat: b.LatLng().Lat.Degrees(), Lon: b.LatLng().Lng.Degrees()},
	)
	return km * 1000
}
