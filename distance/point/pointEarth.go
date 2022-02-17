package point


type PointEr struct {
	width float64
	widthR float64 // Широта в радианах
	longitude float64
	longitudeR float64 // Долгота в радианах
}

func (p PointEr) Width() float64 {
	return p.width
}

func (p PointEr) WidthR() float64 {
	return p.widthR
}

func (p PointEr) Longitude() float64 {
	return p.longitude
}

func (p PointEr) LongitudeR() float64 {
	return p.longitudeR
}



func NewPointEr(width float64, longitude float64) *PointEr {
	widthR := (width*3.14)/180
	longitudeR := (longitude*3.14)/180
	return &PointEr{width: width, widthR: widthR, longitude: longitude, longitudeR: longitudeR}
}

