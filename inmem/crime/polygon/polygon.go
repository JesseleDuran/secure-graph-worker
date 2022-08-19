package polygon

import (
	"github.com/JesseleDuran/secure-graph-worker/inmem/crime/cell"

	"github.com/golang/geo/s2"
)

// Polygon Represents a projection of coordinates to a set of points on a sphere.
// it should be noted that a point on the sphere is a vector in the
// three-dimensional plane.
type Polygon struct {
	Points []s2.Point
}

func MakeFromPoints(points []s2.Point) Polygon {
	return Polygon{Points: points}
}

// Tessellate retrieves the cell representation of a polygon.
func (p Polygon) Tessellate(min int) cell.LinkedList {
	rc := &s2.RegionCoverer{MaxLevel: 15, MinLevel: min}
	return cell.MakeCellListFromCellUnion(rc.Covering(s2.LoopFromPoints(p.Points)))
}
