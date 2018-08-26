package main

import "fmt"

/*CharReplace ...*/
func CharReplace() string {
	myString := "Hello"
	myChars := []rune(myString)                    // char array (runes)
	myChars[0] = 'c'                               // get first element
	myFinalString := string(myChars) + ", World\n" // concat strings
	// fmt.Printf(myFinalString + "\n")
	return myFinalString
}

/*Main Fucntion*/
func main() {
	helloWorld := CharReplace()
	fmt.Printf(helloWorld)
}
