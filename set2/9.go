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
  padded := pad([]byte("YELLOW SUBMARINE"), 14)
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
