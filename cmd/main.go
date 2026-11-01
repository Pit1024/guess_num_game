package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	min, max := 1, 100
	num := rand.IntN(max-min+1) + min
	userNum := 0
	cntr := 0
	fmt.Println("Hello! Try guess a number!")
	_, err := fmt.Scanln(&userNum)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	if userNum > num {
		fmt.Println("too high, try again")
		cntr += 1
		fmt.Scanln(&userNum)
	} else if userNum < num {
		fmt.Println("too low, try again")
		cntr += 1
		fmt.Scanln(&userNum)
	} else {
		fmt.Printf("Congratulations! You win on %d number of try!", cntr)
	}

}
