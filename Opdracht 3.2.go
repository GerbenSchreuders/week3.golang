package main

import (
	"fmt"
	"math"
)



func main() {

	var x int

	fmt.Println(x * 10)

	c := math.Pow10(2)					//0 = math.Pow. Machtheffing
	fmt.Printf("%.1f", c)

	p := math.Round(10.5)				//1 = math.Round. Afronding naar significant getal
	fmt.Printf("%.1f\n", p)

	n := math.Round(-10.5)
	fmt.Printf("%.1f\n", n)

	a2 := math.Floor(1.51)				//2 = math.Floort. Naar beneden afronden
	fmt.Printf("%.1f", a2)

	a3 := math.Sqrt(3*3 + 4*4) 			//3 = math.Sqrt. Worteltrekken
	fmt.Printf("%.1f", a3)

	p := math.Abs(-2)
	fmt.Printf("%.1f\n", p)

	y := math.Abs(2)
	fmt.Printf("%.1f\n", y)



}