package main

import (
	"fmt"

	"github.com/alvarollopi/encapsulation/course"
)

func main() {
	Go := course.New("GO desde cero", 12.34, false)

	Go.SetUserIDs([]uint{12, 56, 89})
	Go.SetClasses(map[uint]string{
		1: "Introducción",
		2: "Estructuras",
		3: "Maps",
	})

	/*css := Course{
		Name: "CSS desde cero", IsFree: true,
	}
	js := Course{}
	js.Name = "Curso JS"
	js.UserIDs = []uint{12, 67}
	fmt.Println(Go.Name)
	fmt.Printf("%+v \n", css)
	fmt.Printf("%+v", js)
	*/
	Go.SetName("POO con Go")
	fmt.Println(Go.Name())
	Go.PrintClasses()
	//Go.ChangePrice(67.12)
	//fmt.Println(Go.Price)

}
