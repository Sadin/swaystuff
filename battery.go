package main

import (
	"fmt"

	"github.com/distatus/battery"
)

func main() {
	batteries, err := battery.GetAll()
	if err != nil {
		fmt.Println("Could no getbattery info")
		return
	}

	for i, battery := range batteries {
		fmt.Printf("Bat%d:", i)
		fmt.Printf("state: %s, ", battery.State)
	}
}
