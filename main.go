package main

import (
	"fmt"
	"math"
	"time"
)

// List of known Prime counts under various values
// Source: https://primes.utm.edu/howmany.html
var expectedCount = map[int]int{
	10:   4,         // <1ms
	100:  25,        // <1ms
	1000: 168,       // ~1ms
	1e4:  1229,      // ~1ms
	1e5:  9592,      // ~1ms
	1e6:  78498,     // ~<1s
	1e7:  664579,    // ~<1s
	1e8:  5761455,   // ~<1s
	1e9:  50847534,  // ~8s
	1e10: 455052511, // ~2m  - Feel free to go higher than 10 billion (I don't think you need/want to though)
}

type PrimeSieve struct {
	bitArray []bool
	size     int
	calcEnd  int
}

// We'll calculate all Primes less than this value
// Refer to the estimated time comments in the 'expectedValues' slice
// before changing this to some mad value like 1e10000
const sieveSize = 1e6 // Default: 1e6

// Entry point
func main() {

	if expectedCount[sieveSize] > 0 {

		primeSieve := PrimeSieve{
			bitArray: make([]bool, sieveSize),
			size:     sieveSize,
			calcEnd:  int(math.Sqrt(float64(sieveSize))),
		}

		primeSieve.prepBitArray()

		fmt.Printf("Finding Primes less than: %d...\n", primeSieve.size)

		timerStart := time.Now()

		primeSieve.findPrimes()
		primeSieve.countPrimes()

		timerElapsed := time.Since(timerStart)

		fmt.Printf("Duration: %f second(s)", float64(timerElapsed.Seconds()))
	} else {
		fmt.Printf("Your chosen 'sieveSize' of %d, should exist within the 'expectedValues' slice. Won't continue as we can't validate the result - sorry!", int64(sieveSize))
	}
}

// Calculate primes upto the given sieve size,
// we're using the "Sieve of Eratosthenes" algorithm
func (sieve *PrimeSieve) findPrimes() {
	var factor int = 2

	// Loop through all values less than the square root
	// of the total sieve size, as long as we've eliminated
	// all multiples of the Primes less than that value, we
	// don't need to process anymore
	for num := factor; num <= sieve.calcEnd; num++ {

		if sieve.bitArray[num] {
			factor = num

			// Eliminate the multiples of our new prime number
			latestPrimeSquared := int(math.Pow(float64(factor), 2))

			for x := latestPrimeSquared; x < sieve.size; x += factor {
				sieve.bitArray[x] = false
			}
		}
	}
}

// Set all bits in Bit Array to true, we can then mark all none primes as false.
// I'm using a pointer to resolve the memory address due to the fact that we're
// making changes we want to carry over to the next stage
func (sieve *PrimeSieve) prepBitArray() {

	for x := range sieve.bitArray {

		if x >= 2 {
			sieve.bitArray[x] = true
		}
	}
}

// Count all of the items in the bit array, if they are True then we have found a prime, so count it!
// We don't need to use a pointer, as we're not transforming the data in any way - only reading
func (sieve PrimeSieve) countPrimes() {
	primesFound := 0

	for x := range sieve.bitArray {

		if sieve.bitArray[x] {
			primesFound++
		}
	}

	primesExpected := expectedCount[sieve.size]

	if primesFound == primesExpected {
		fmt.Printf("Success! We found %d Primes!\n", primesFound)
	} else {
		fmt.Printf("Failed! We found %d Primes, expected %d\n", primesFound, primesExpected)
	}
}
