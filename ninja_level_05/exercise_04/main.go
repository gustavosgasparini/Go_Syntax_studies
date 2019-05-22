package main

import "fmt"

func main() {
	annStruct := struct {
		firstName     string
		lastName      string
		isMyGirfriend bool
		isMyFamily    bool
		isItMe        bool
	}{
		firstName:     "Gustavo",
		lastName:      "Gasparini",
		isMyGirfriend: false,
		isMyFamily:    false,
		isItMe:        true,
	}

	fmt.Println(annStruct)
}
