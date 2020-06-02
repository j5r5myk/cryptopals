/* Cryptopals Challenge #5
 * Implement repeating-key XOR
 * Author: j5r5myk
 */
package main

import (
    "fmt"
//    "bufio"
    "os"
    "encoding/hex"
    "io/ioutil"
  //  "bytes"
)

func main() {
    key := "ICE"
    if len(os.Args) < 2 {
        fmt.Printf("Usage: 5 <path-to-ciphertext>\n")
        os.Exit(0)
    }
    // Pass 
    message, _ := ioutil.ReadFile(os.Args[1])
    // Perform repeating xor
    output := repeatXor(message, key)
    // Print result
    fmt.Printf("%s", hex.Dump(output))
}
func repeatXor(message []byte, key string) []byte {
    output := make([]byte, len(message) - 1)
    // Skip final newline
    for i := 0; i < len(message) - 1; i++ {
      output[i] = message[i] ^ key[i%len(key)]
    }
    return output
}

