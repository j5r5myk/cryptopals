package main

import (
  "fmt"
  //"strconv"
  "encoding/hex"
)

func main() {
  key := "ICE"
  input := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
  // Read input
  // Perform repeating xor
  output := repeatXor(input, key)
  // Print result
  fmt.Printf("%s", hex.Dump(output))
}
func repeatXor(s string, key string) []byte {
  output := make([]byte, len(s))
  for i := 0; i < len(s); i++ {
    output[i] = s[i] ^ key[i%len(key)]
  }
  return output
}

