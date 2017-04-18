package main

import (
  "fmt"
  //"encoding/hex"
)

func main() {
  str1 := "this is a test"
  str2 := "wokka wokka!!!"
 // read input
 // calc hamming d
 fmt.Println(calcHamming(str1, str2))
}
func calcHamming(str1 string, str2 string) int {
  result := make([]byte, len(str1))
  hamming := 0
  // XOR each byte
  for i := 0; i < len(str1); i++ {
    result[i] = str1[i] ^ str2[i]
    // Count differieng bytes
    for result[i] > 0 {
      hamming++
      result[i] &= result[i] - 1
    }
  }
  return hamming
}
