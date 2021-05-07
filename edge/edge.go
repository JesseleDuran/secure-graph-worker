package edge

import (
  "osm-graph-parser/parser/resources"
  "osm-graph-parser/tag"

  "github.com/golang/geo/s2"
  "github.com/paulmach/osm"
)

type Edges []uint64

type Edges2 []Node

type Node struct {
  ID     int64
  CellId uint64
  Name   string
}

func FromOSMRelation(r osm.Relation, aux Edges) Edges {
  for i := 0; i < len(r.Members); i++ {
    //TODO: make this as a go routine.
    m := r.Members[i]
    if m.Type == "node" {
      if v, ok := resources.Nodes[m.Ref]; ok {
        aux = append(aux, toToken(v.Lat, v.Lon))
      }
    }
    if m.Type == "way" {
      w := resources.Ways[m.Ref]
      aux = append(aux, FromWay(w)...)
    }
    if m.Type == "relation" {
      if v, ok := resources.Relations[m.Ref]; ok {
        aux = append(aux, FromOSMRelation(v, Edges{})...)
      }
    }
  }
  return aux
}

func FromWay(w osm.Way) Edges {
  r := make(Edges, 0)
  for i := range w.Nodes {
    if v, ok := resources.Nodes[w.Nodes[i].ID.FeatureID().Ref()]; ok {
      r = append(r, toToken(v.Lat, v.Lon))
    }
  }
  return r
}

func FromOSMRelation2(r osm.Relation, aux Edges2) Edges2 {
  for i := 0; i < len(r.Members); i++ {
    //TODO: make this as a go routine.
    m := r.Members[i]
    if m.Type == "node" {
      if v, ok := resources.Nodes[m.Ref]; ok {
        aux = append(aux, Node{
          ID:     int64(v.ID),
          CellId: toToken(v.Lat, v.Lon),
        })
      }
    }
    if m.Type == "way" {
      w := resources.Ways[m.Ref]
      aux = append(aux, FromWay2(w)...)
    }
    if m.Type == "relation" {
      if v, ok := resources.Relations[m.Ref]; ok {
        aux = append(aux, FromOSMRelation2(v, Edges2{})...)
      }
    }
  }
  return aux
}

func FromWay2(w osm.Way) Edges2 {
  r := make(Edges2, 0)
  for i := range w.Nodes {
    if v, ok := resources.Nodes[w.Nodes[i].ID.FeatureID().Ref()]; ok {
      tagsNode := tag.FromOSMTags(v.Tags)
      name := ""
      if nameNode, ok := tagsNode["name"]; ok {
        name = nameNode
      } else {
        tagsWay := tag.FromOSMTags(w.Tags)
        if nameWay, ok := tagsWay["name"]; ok {
          name = nameWay
        }
      }
      r = append(r, Node{
        ID:     int64(v.ID),
        CellId: toToken(v.Lat, v.Lon),
        Name:   name,
      })
    }
  }
  return r
}

//TODO TO 17
func toToken(lat, lng float64) uint64 {
  return uint64(s2.CellFromPoint(s2.PointFromLatLng(
    s2.LatLngFromDegrees(lat, lng))).ID().Parent(17)) //19
}
