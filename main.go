package main

import (
	"log"

	ps "github.com/mitchellh/go-ps"
)

func main() {
	processList, err := ps.Processes()
	if err != nil {
		log.Println("ps.Processes() Failed")
		return
	}

	// map ages
	for x := range processList {
		var process ps.Process
		process = processList[x]
		log.Printf("%d\t%s\n", process.Pid(), process.Executable())
	}
	log.Printf("Total Processes: %d", len(processList))
}
