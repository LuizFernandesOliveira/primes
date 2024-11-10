package main

import (
	"fmt"
	"github.com/LuizFernandesOliveira/primes/functions"
	"os"
	"strconv"
	"time"
)

func findMaxPrimeWithinDuration(duration time.Duration) int {
	startTime := time.Now()
	maxPrime := 2
	current := 2

	for {
		if time.Since(startTime) > duration {
			break
		}
		if functions.IsPrime(current) {
			maxPrime = current
		}
		current++
	}

	return maxPrime
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <duration_in_seconds>")
		return
	}

	seconds, err := strconv.Atoi(os.Args[1])
	if err != nil || seconds <= 0 {
		fmt.Println("Please provide a valid positive integer for the duration in seconds.")
		return
	}

	duration := time.Duration(seconds) * time.Second

	maxPrime := findMaxPrimeWithinDuration(duration)
	fmt.Printf("The largest prime found in %d seconds is: %d\n", seconds, maxPrime)
}
