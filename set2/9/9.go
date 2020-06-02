/* Cryptopals Challenge #9
 * Implement PKCS#7 padding
 * Author: j5r5myk
 */
package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
)

func main() {
    var size int
    reader := bufio.NewReader(os.Stdin)

    fmt.Println("Welcome to PKSC#7 padding\nPlease enter a block:")
    // Read line from stdin 
    input, _ := reader.ReadString('\n')
    // Remove trailing newline
    input = strings.TrimSuffix(input, "\n")
    fmt.Println("Please enter padding size:")
    fmt.Scanln(&size)
    padded := pad([]byte(input), size)
    fmt.Printf("Result: %s length: %d\n", string(padded), len(padded))
}

func pad(input []byte, size int) []byte {
    if len(input) > size {
      fmt.Println("Padding error: block is larger than desired padding length")
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
