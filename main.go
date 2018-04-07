package main

import (
	"fmt"
	"os"
)

// Main function
func main() {
	//Initialize variable that holds status
	var nrpeStatus NrpeStatus
	//Call function that does all work with
	checkNrpe(&nrpeStatus)

	//Print status message and exit with correct code
	fmt.Println(nrpeStatus.Message)
	os.Exit(int(nrpeStatus.Code))
}

// This function does all work and sets the status message and code
func checkNrpe(nrpeStatus *NrpeStatus) {
	//TODO: add logic here

	//set a message that will show in Nagios
	nrpeStatus.Message = "This message shows in Nagios"
	//set status. This shows as gree, yellow, red or gray in nagios
	nrpeStatus.Code = OK
}
