package main

import (
	"fmt"
	"os"
)

// Main function
func main() {
	//Use pointer or channel?
	var nrpeStatus NrpeStatus = checkNrpe()

	fmt.Println(nrpeStatus.Message)
	os.Exit(int(nrpeStatus.Code))
}

// This function do all work
func checkNrpe() NrpeStatus {
	var nrpeStatus NrpeStatus
	nrpeStatus.Message = "This message shows in Nagios"
	nrpeStatus.Code = OK

	return nrpeStatus
}
