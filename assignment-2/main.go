package main

import "fmt"


type shape interface {
	getArea() float64
}

type square struct{
	sideLength float64
}
type triangle struct{
	base float64
	height float64
}

func main() {
	sq := square{
		sideLength: 10,
	}
	tr := triangle{
		base: 2,
		height: 10,
	}

	fmt.Println("Square:", printArea(sq))
	fmt.Println("Triangle:",  printArea(tr))

}

// Added
func printArea(sh shape) (float64){
	return sh.getArea()
}

func (sq square) getArea() float64 {
	return float64(sq.sideLength * sq.sideLength)
}

func (tr triangle) getArea() float64 {
	return float64(0.5 * tr.base * tr.height)
}