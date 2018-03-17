package main

import (
	"fmt"
	"os"
)

// Main function
func main() {
	//Initialize variable that holds status
	var nrpeStatus NrpeStatus
	//Call function that does all work with a pointer
	checkNrpe(&nrpeStatus)

	//Print status message and exit with correct code
	fmt.Println(nrpeStatus.Message)
	os.Exit(int(nrpeStatus.Code))
}

// This function does all work and sets the status message and code
func checkNrpe(nrpeStatus *NrpeStatus) {
	nrpeStatus.Message = "This message shows in Nagios"
	nrpeStatus.Code = OK
}
