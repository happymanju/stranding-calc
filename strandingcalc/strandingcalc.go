package strandingcalc

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Run(args []string) int {
	if len(args) < 3 {
		log.Println("Not enough arguments; Exiting...")
		return 1
	}

	target := strings.ToLower(args[0])
	currentAmnt, err := strconv.Atoi(args[1])
	if err != nil {
		log.Printf("Could not convert to a number: %s\n", args[1])
		return 1
	}
	maxAmnt, err := strconv.Atoi(args[2])
	if err != nil {
		log.Printf("Could not convert to a number: %s\n", args[2])
		return 1
	}

	matContainers, ok := mats[target]
	if !ok {
		log.Println("Material not found; Exiting...")
		return 1
	}
	currentContainer := 0
	remaining := maxAmnt - currentAmnt
	containersNeeded := map[int]int{}

	for currentContainer != len(matContainers) {
		_, ok := containersNeeded[matContainers[currentContainer]]
		if !ok {
			containersNeeded[matContainers[currentContainer]] = 0
		}

		if remaining >= matContainers[currentContainer] {
			remaining -= matContainers[currentContainer]
			containersNeeded[matContainers[currentContainer]]++
		} else {
			currentContainer++
		}
	}

	for k, v := range containersNeeded {
		fmt.Printf("%d: %d\n", k, v)
	}
	fmt.Printf("Remainder: %d\n", remaining)

	return 0
}
