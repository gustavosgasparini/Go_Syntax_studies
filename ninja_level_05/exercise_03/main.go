package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main()  {
	t := truck{
		vehicle: vehicle {
			doors: 4,
			color: "red",
		},
		fourWheel: true,
	}
	fmt.Println(t)
	fmt.Println("color:", t.color, "/ doors:", t.doors, "/ is four wheeled?", t.fourWheel)

	s := sedan{
		vehicle: vehicle {
			doors: 2,
			color: "Green",
		},
		luxury: true,
	}
	fmt.Println(s)
	fmt.Println("color:", s.color,"/ doors:", s.doors, "/ is luxury?", s.luxury)
}
