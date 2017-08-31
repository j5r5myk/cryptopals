/* Cryptopals Challenge #9
 * Implement PKCS#7 padding
 * Author: j5r5myk@gmail.com
 */
package main

import (
  "fmt"
  "os"
)

func main() {
  fmt.Println("Welcome to PKSC#7 padding\nPlease enter a block:")
  var input string
  fmt.Scanln(&input)
  fmt.Println("Please enter padding size:")
  var size int
  fmt.Scanln(&size)
  padded := pad([]byte(input), size)
  fmt.Printf("Result: %s length: %d\n", string(padded), len(padded))
}

func pad(input []byte, size int) []byte {
  if len(input) > size {
    fmt.Println("Padding error: block is larger than length")
    os.Exit(1)
  } else if len(input) == size {
    return input
  } else {
    tail := make([]byte, size - (len(input)))
    for i, _ := range tail {
      tail[i] = 0x04
    }
    out := append(input, tail...)
    return out
  }
  return []byte("bogus")
}
