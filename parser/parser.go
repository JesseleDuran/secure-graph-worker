package parser

import (
	"context"
	"os"
	"osm-graph-parser/edge"
	"osm-graph-parser/graph"
	"osm-graph-parser/parser/resources"
	"osm-graph-parser/tag"

	"github.com/paulmach/osm"
	"github.com/paulmach/osm/osmxml"
)

func FromOSMFile(path string) (graph.Graph, error) {
	f, err := os.Open(path)
	if err != nil {
		return graph.Graph{}, err
	}
	defer f.Close()
	scanner := osmxml.New(context.Background(), f)
	defer scanner.Close()

	for scanner.Scan() {
		o := scanner.Object()
		switch o.ObjectID().Type() {
		case "node":
			n := *o.(*osm.Node)
			resources.Nodes[n.ID.FeatureID().Ref()] = n

		//case "relation":
		//  r := *o.(*osm.Relation)
		//  auxTags := tag.FromOSMTags(r.Tags)
		//  if _, ok := auxTags["building"]; !ok {
		//    resources.Relations[r.ID.FeatureID().Ref()] = r
		//  }

		case "way":
			w := *o.(*osm.Way)
			auxTags := tag.FromOSMTags(w.Tags)
			v, ok := auxTags["highway"]
			//v1, ok1 := auxTags["motor_vehicle"];
			//v2, ok2 := auxTags["motorcar"];
			if ok && (v == "motorway" ||
				v == "motorway_link" ||
				v == "trunk" ||
				v == "trunk_link" ||
				v == "primary" ||
				v == "primary_link" ||
				v == "secondary" ||
				v == "secondary_link" ||
				v == "tertiary" ||
				v == "tertiary_link" ||
				v == "residential" ||
				v == "service" ||
				v == "unclassified" ||
				v == "living_street") {
				resources.Ways[w.ID.FeatureID().Ref()] = w
			}

		default:
			continue
		}
	}
	err = scanner.Err()
	if err != nil {
		return graph.Graph{}, err
	}

	g := graph.Graph{Edges: make([]edge.Edges, 0), Edges2: make([]edge.Edges2, 0)}

	//for _, v := range resources.Relations {
	//  g.Edges = append(g.Edges, edge.FromOSMRelation(v, edge.Edges{}))
	//}
	//for _, v := range resources.Ways {
	//  g.Edges = append(g.Edges, edge.FromWay(v))
	//}

	//for _, v := range resources.Relations {
	//  g.Edges2 = append(g.Edges2, edge.FromOSMRelation2(v, edge.Edges2{}))
	//}
	for _, v := range resources.Ways {
		g.Edges3 = append(g.Edges3, edge.FromWay3(v))
	}

	resources.Ways = nil
	resources.Relations = nil
	resources.Nodes = nil

	return g, nil
}

func FromOSMFileV2(path string) (graph.GraphV2, error) {
	f, err := os.Open(path)
	if err != nil {
		return graph.GraphV2{}, err
	}
	defer f.Close()
	scanner := osmxml.New(context.Background(), f)
	defer scanner.Close()

	for scanner.Scan() {
		o := scanner.Object()
		switch o.ObjectID().Type() {
		case "node":
			n := *o.(*osm.Node)
			resources.Nodes[n.ID.FeatureID().Ref()] = n

		//case "relation":
		//  r := *o.(*osm.Relation)
		//  auxTags := tag.FromOSMTags(r.Tags)
		//  if _, ok := auxTags["building"]; !ok {
		//    resources.Relations[r.ID.FeatureID().Ref()] = r
		//  }

		case "way":
			w := *o.(*osm.Way)
			auxTags := tag.FromOSMTags(w.Tags)
			v, ok := auxTags["highway"]
			//v1, ok1 := auxTags["motor_vehicle"];
			//v2, ok2 := auxTags["motorcar"];
			if ok && (v == "motorway" ||
				v == "motorway_link" ||
				v == "trunk" ||
				v == "trunk_link" ||
				v == "primary" ||
				v == "primary_link" ||
				v == "secondary" ||
				v == "secondary_link" ||
				v == "tertiary" ||
				v == "tertiary_link" ||
				v == "residential" ||
				v == "service" ||
				v == "unclassified" ||
				v == "living_street") {
				resources.Ways[w.ID.FeatureID().Ref()] = w
			}

		default:
			continue
		}
	}
	err = scanner.Err()
	if err != nil {
		return graph.GraphV2{}, err
	}

	g := graph.GraphV2{Nodes: make(map[uint64]graph.Node)}

	//for _, v := range resources.Relations {
	//  g.Edges = append(g.Edges, edge.FromOSMRelation(v, edge.Edges{}))
	//}
	//for _, v := range resources.Ways {
	//  g.Edges = append(g.Edges, edge.FromWay(v))
	//}

	//for _, v := range resources.Relations {
	//  g.Edges2 = append(g.Edges2, edge.FromOSMRelation2(v, edge.Edges2{}))
	//}
	for _, v := range resources.Ways {
		g.RelateNodesFromWay(v)
	}

	resources.Ways = nil
	resources.Relations = nil
	resources.Nodes = nil

	return g, nil
}
