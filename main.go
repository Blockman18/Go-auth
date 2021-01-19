package main

import (
	"fmt"
	"time"
)

func main() {
	var Username string
	var Password string
	fmt.Print("Username?:")
	fmt.Scan(&Username)
	fmt.Print("\nPassowrd?:")
	fmt.Scan(&Password)

	if Username == Password {
		fmt.Print("invalid input!")
		time.Sleep(5 * time.Second)

	} else if Username == "Your username her" && Password == "Your password here" {

		fmt.Print("good job redirecting you!")
		time.Sleep(5 * time.Second)

	} else {
		fmt.Print("hey fag you fucked up!")
		time.Sleep(5 * time.Second)

	}
}
