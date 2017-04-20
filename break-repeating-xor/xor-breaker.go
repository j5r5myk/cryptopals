package main

import (
  "fmt"
  //"encoding/hex"
)

func main() {
  MIN_KEYSIZE := 2
  MAX_KEYSIZE := 40
  str1 := "this is a test"
  str2 := "wokka wokka!!!"
  // read input
  // Find key size
  keySize := findKeySize(MIN_KEYSIZE, MAX_KEYSIZE)
  // Divide input into keySized blocks
}
func findKeySize(min int, max int input string) int {
  lowHam := 1000
  for i := min; i <= max; i++ {
    // Hamming distance b/w first 2 chunks
    rawHam := calcHamming(input[0:keysize], input[keysize, 2 * keysize])
    // Normalize and compare
    normHam := rawHam / keySize
    if normHam < lowHam {
      lowHam = normHam
    }
  }
  return lowHam
}
func calcHamming(str1 string, str2 string) int {
  result := make([]byte, len(str1))
  hamming := 0
  // XOR each byte
  for i := 0; i < len(str1); i++ {
    result[i] = str1[i] ^ str2[i]
    // Count differing bytes
    for result[i] > 0 {
      hamming++
      result[i] &= result[i] - 1
    }
  }
  return hamming
}
func createBlocks(s string, keySize int) {
  // Array, each element is keySize # bytes
  // len(s) / keysize elements
  // I guess it will be an array of strings
  blocks
}
func transposeBlocks() {

}
