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
  // Read input
  // Find key size
  keysize := findKeySize(MIN_KEYSIZE, MAX_KEYSIZE, input)
  // Divide input into keySized blocks
  blocks := createBlocks(input, keysize)
  print(blocks)
}
func findKeySize(min int, max int, input string) int {
  lowHam := 1000
  bestSize := 0
  // TODO check min/max vs input length
  for i := min; i <= max; i++ {
    // Hamming distance b/w first 2 chunks
    rawHam := calcHamming(input[0:i - 1], input[i: (2 * i) - 1])
    // Normalize and compare
    normHam := rawHam / i
    if normHam < lowHam {
      lowHam = normHam
      bestSize = i
    }
  }
  return bestsize
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
func createBlocks(s cipher, keysize int) []string{
  // Array, each element is keySize # bytes
  // len(s) / keysize elements
  // I guess it will be an array of strings
  blocks := make([]string, len(s) / keysize)
  for i := 0; i < len(blocks); i++ {
    // ith keysize block
    blocks[i] = cipher[i:((i + 1) * keysize) - 1]
  }
  return blocks
}
func transposeBlocks() {

}
