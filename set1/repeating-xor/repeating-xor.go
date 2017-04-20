package main

import (
  "fmt"
  //"strconv"
  "bufio"
  "os"
  "encoding/hex"
  "bytes"
)

func main() {
  key := "ICE"
  //input := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
  // Read input
  scanner := bufio.NewScanner(os.Stdin)
  // Concat all the lines
  var buffer bytes.Buffer
  for scanner.Scan() {
    //fmt.Println(scanner.Text())
    buffer.WriteString(scanner.Text() + "\n")
  }
  // Perform repeating xor
  output := repeatXor(buffer.String(), key)
  // Print result
  fmt.Printf("%s", hex.Dump(output))
}
func repeatXor(s string, key string) []byte {
  output := make([]byte, len(s) - 1)
  // Skip final newline
  for i := 0; i < len(s) - 1; i++ {
    output[i] = s[i] ^ key[i%len(key)]
  }
  return output
}

