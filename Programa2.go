package main

import "fmt"

func prob_Num() {
	for i := 1; i <= 100; i++ {
		var pin int = 3
		var pan int = 5

		if i%pin == 0 && i%pan == 0 {
			fmt.Printf("%v são Pin e Pan", i)
			fmt.Println("")
		} else if i%pin == 0 {
			fmt.Printf("%v é Pin", i)
			fmt.Println("")
		} else if i%pan == 0 {
			fmt.Printf("%v é Pan", i)
			fmt.Println("")
		}
	}
}

func main() {
	prob_Num()
}
