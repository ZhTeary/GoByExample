package function

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func T() {
	r := rect{10, 5}

	fmt.Println("area", r.area())
	fmt.Println("perm", r.perim())

	rp := &r
	fmt.Println("area", rp.area())
	fmt.Println("perm", rp.perim())

}
