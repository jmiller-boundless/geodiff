package shpfile

import (
	"github.com/jonas-p/go-shp"
	"log"
	"fmt"
	"reflect"
	"math"
)

func Open(path string){
	// open a shapefile for reading
	shape, err := shp.Open(path)
	if err != nil { log.Fatal(err) }
	defer shape.Close()

	// fields from the attribute table (DBF)
	fields := shape.Fields()

	// loop through all features in the shapefile
	for shape.Next() {
		n, p := shape.Shape()

		// print feature
		fmt.Println(reflect.TypeOf(p).Elem(), p.BBox())

		// print attributes
		for k, f := range fields {
			val := shape.ReadAttribute(n, k)
			fmt.Printf("\t%v: %v\n", f, val)
		}
		fmt.Println()
	}

}

type PolygonWrapper struct {
	Polygon shp.Polygon
}

func (p PolygonWrapper) EqualsExact(other PolygonWrapper,tolerance float64) bool {
	polygon := p.Polygon
	otherpolygon := other.Polygon
	if polygon.NumParts!=otherpolygon.NumParts{
		return false
	}
	if polygon.NumPoints!=otherpolygon.NumPoints{
		return false
	}
	for k,v :=range polygon.Points{
		if !equalPoints(v,otherpolygon.Points[k],tolerance){
			return false
		}
	}
	return true
}

func equalPoints(point1 shp.Point, point2 shp.Point, tolerance float64) bool{
	return pointDistance(point1,point2) <= tolerance
}

func pointDistance(point1 shp.Point, point2 shp.Point)float64{
	dx := point1.X - point2.X
	dy := point1.Y - point2.Y
	return math.Sqrt(dx*dx+dy*dy)
}
