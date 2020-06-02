/*
 * Calculate Hamming distance between two strings
 * Author: j5r5myk
 */
package main

import (
    "fmt"
    "flag"
)

func main() {
    // Debug option for a nice visualization
    debugPtr := flag.Bool("d", false, "debug")
    flag.Parse()
    string1 := flag.Args()[0]
    string2 := flag.Args()[1]
    fmt.Printf("%d\n", calculateHammingDistance(string1, string2, *debugPtr))
}

func calculateHammingDistance(str1 string, str2 string, debug bool) int {
    result := make([]byte, len(str1))
    hamming := 0
    // XOR each byte
    for i := 0; i < len(str1); i++ {
        result[i] = str1[i] ^ str2[i]
        // Count differing bytes
        if debug {
            fmt.Printf("i=%d\n", i)
        }
        for result[i] > 0 { 
            if debug {
                fmt.Printf("\tStart loop result\t%08b (%d)\n", result[i], result[i])
            }
            hamming++
            /* Remove differing bits until zeroed out.
             * We want to count how many 1's exist in result[i].
             * This zeros 1 bit at a time, by the power of bitwise operations.
             * E.g. 80 (1010000) & 79 (1001111) = 1000000
             */
            result[i] = result[i] & (result[i] - 1)
            if debug {
                fmt.Printf("\tEnd loop result\t\t%08b (%d)\n\n", result[i], result[i])
            }
        }   
    }   
    return hamming
}
