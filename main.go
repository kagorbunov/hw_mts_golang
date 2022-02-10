package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	var a,b,c *float64
	a,b,c = new(float64), new(float64),new(float64)
	fmt.Println("Введите a\n")
	fmt.Scan(a)
	fmt.Println("Введите b\n")
	fmt.Scan(b)
	fmt.Println("Введите c\n")
	fmt.Scan(c)

	var x1,x2 float64

	d, err := desc(*a,*b,*c)
	if err == nil {
		x1, x2 = findCorn(d, *a, *b)
		fmt.Printf("Корни уравнения x1 = %f, x2 = %f",x1,x2)
	} else {
		fmt.Println(err)
	}

}

func findCorn(d float64, a float64, b float64) (float64, float64){
	x1 := float64((math.Sqrt(d) - b) / (2 * a))
	x2 := float64((-math.Sqrt(d) - b) / (2 * a))
	return x1, x2
}

func desc(a float64, b float64, c float64) (float64,error){
	d := float64(b*b - 4*a*c)
	if d > 0 {
		return d, nil
	}
	return 0, errors.New("Дискриминант не действителен")
}
