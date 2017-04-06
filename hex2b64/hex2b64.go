package main

import (
  "fmt"
)

func main() {
  readHexLine()
}

func readHexLine() {
  var i int
  fmt.Scanf("%X", &i)
  fmt.Printf("saw hex %x\n", i)
  fmt.Printf("saw dec %d\n", i)
}

func isHex(hex string) bool {
  fmt.Println("string length: ", len(hex))
  for i := 0; i < len(hex) - 1; i++ {
    if (hex[i] < '0' || hex[i] > '9') {
      return false
    }
    if (hex[i] < 'a' || hex[i] > 'f') {
      return false
    }
  }
  return true
}

func hexToB64(hex string) {
  // read 3 octets
  // convert into 3 b64 chars
  for i := 0; i < len(hex); i+=3 {
    //hex[i]+hex[i+1]+hex[i+2]
  }
}
