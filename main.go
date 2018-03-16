package main

import (
	"fmt"
	"os"
)

// Main function
func main() {
	var returnCode int = int(checkNrpe())
	os.Exit(returnCode)
}

// This function do all work
func checkNrpe() ReturnStatusEnum {
	fmt.Println("This message shows in Nagios")
	return OK
}
