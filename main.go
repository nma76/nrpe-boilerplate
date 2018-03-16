package main

import (
	"fmt"
	"os"
)

func main() {
	var returnCode int = int(checkNrpe())
	os.Exit(returnCode)
}
func checkNrpe() ReturnStatusEnum {
	fmt.Println("This message shows in Nagios")
	return OK
}
