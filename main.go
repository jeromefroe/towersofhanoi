package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("ERROR: Invalid arguments.\n\n")
		exit()
	}

	disks, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("ERROR: Unable to parse argument as an integer: %v.\n\n", err)
		exit()
	}

	if disks < 0 {
		fmt.Printf("ERROR: Number of disks must be greater than zero.\n\n")
		exit()
	}

	move(disks, "left", "right", "middle")
}

func move(n int, source, destination, intermediate string) {
	if n == 0 {
		return
	}

	move(n-1, source, intermediate, destination)

	fmt.Printf("move %d from %s to %s\n", n, source, destination)

	move(n-1, intermediate, destination, source)
}

func exit() {
	fmt.Printf("Usage: %s [number of disks]\n", os.Args[0])
	os.Exit(1)
}
