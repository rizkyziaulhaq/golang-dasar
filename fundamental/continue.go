package main

import "fmt"

func sayContinue(){
	
	// Continue
	for i := 0; i <10; i++ {
		if i % 2 == 0{
			continue
		}

		fmt.Println("Perulangan ke", i)
	}
}