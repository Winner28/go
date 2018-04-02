package interfaces

import "fmt"

// Renderer type with Render method
type Renderer interface {
	Render()
}

type myString string
type myInt int

func (str myString) Render() {
	fmt.Println(str)
}

func (in myInt) Render() {
	fmt.Println(in)
}

// Display Render
func Display(renderer []Renderer) {
	for _, rend := range renderer {
		rend.Render()
	}
}

// Run inter
func Run() {
	var integer myInt = 10
	var str myString = "stringer!"
	renderers := []Renderer{
		integer,
		str,
	}
	Display(renderers)
}
