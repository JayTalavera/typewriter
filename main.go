package typewriter

import (
	"fmt"
	"strings"
	"time"
)

// provide a string to be typed out character by character and an int to specify the time in milliseconds between each typed character
func Println(x string, y int) string {
	// add newline to end of string
	var myText string = x + "\n"
	if y == 0 {
		// if y equals 0 it is defaulted to 50
		y = 50
	}

	for _, v := range myText {
		// loop through each rune
		fmt.Printf(string(v))
		time.Sleep(time.Duration(y) * time.Millisecond)
	}
	return myText
}

// provide a string to be typed out character by character and an int to specify the time in milliseconds between each typed character
func Printf(x string, y int) string {
	// does not have newline appended to end of string
	var myText string = x
	if y == 0 {
		// if y equals 0 it is defaulted to 50
		y = 50
	}

	for _, v := range myText {
		// loop through each rune
		fmt.Printf(string(v))
		time.Sleep(time.Duration(y) * time.Millisecond)
	}
	return myText
}

func Words(x string, y int) []string {
	myText := strings.Fields(x)
	if y == 0 {
		y = 50
	}

	for _, v := range myText {
		fmt.Printf(string(v + " "))
		time.Sleep(time.Duration(y) * time.Millisecond)
	}
	return myText
}
