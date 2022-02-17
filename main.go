package main

import (
	"GO_MTS/homework5/distance"
	"GO_MTS/homework5/distance/point"
	"fmt"
)

func main() {
	city1 := "Москва"
	city2 := "Санкт-Петербург"
	LongBetween := distance.NewLineBetween(
		*point.NewPointEr(55.75,37.62),
		*point.NewPointEr(59.94,30.31))
	distanceBetween := LongBetween.Distance()
	fmt.Printf(
		"Расстояние между городом - %v и городом - %v равно %.2f километров.",
		city1,
		city2,
		distanceBetween)
}