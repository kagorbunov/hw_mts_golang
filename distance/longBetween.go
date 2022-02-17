package distance

import (
	"GO_MTS/homework5/distance/point"
	"math"
)

type LineBetween struct {
	place1 point.PointEr
	place2 point.PointEr
}

const R = 6400

func (l LineBetween) Place1() point.PointEr {
	return l.place1
}

func (l LineBetween) Place2() point.PointEr {
	return l.place2
}

func NewLineBetween(place1 point.PointEr, place2 point.PointEr) *LineBetween {
	return &LineBetween{place1: place1, place2: place2}
}

func (l LineBetween) Distance() (distance float64) {
	a := math.Pow(math.Sin((l.place2.WidthR()-l.place1.WidthR())/2), 2)
	b := math.Cos(l.place1.WidthR())*math.Cos(l.place2.WidthR())*math.Pow(
		math.Sin(
		(l.place2.LongitudeR()-l.place1.LongitudeR())/2),
		2)
	distance = 2*R*math.Asin(math.Sqrt(a+b))
	return distance
}