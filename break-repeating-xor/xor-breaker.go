package main

import (
  "fmt"
  //"encoding/hex"
)

func main() {
  MIN_KEYSIZE := 2
  MAX_KEYSIZE := 40
  //str1 := "this is a test"
  //str2 := "wokka wokka!!!"
  // Read input
  input := "THISISATEST"
  // Find key size
  keysize := findKeySize(MIN_KEYSIZE, MAX_KEYSIZE, input)
  fmt.Printf("Likely key size: %d\n", keysize)
  // Divide input into keySized blocks
  blocks := createBlocks(input, keysize)
  tb := transposeBlocks(blocks)
  // Print blocks
  for pos, block := range blocks {
    fmt.Printf("%d: %s\n", pos, block)
  }
}
func findKeySize(min int, max int, c string) int {
  lowHam := 1000
  bestSize := 0
  // Check max vs input length
  if max > len(c) / 2 {
    max = len(c) / 2
  }
  for i := min; i <= max; i++ {
    fmt.Println("Testing keysize: ", i)
    // Hamming distance b/w first 2 chunks
    rawHam := calcHamming(c[0:i], c[i:2 * i])
    // Normalize and compare
    normHam := rawHam / i
    if normHam < lowHam {
      lowHam = normHam
      bestSize = i
    }
  }
  return bestSize
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
func createBlocks(c string, keysize int) []string{
  // Each array element is keySize # bytes
  // len(s) / keysize + 1 elements for remainder
  blocks := make([]string, (len(c) / keysize) + 1)
  for i := 0; i < len(blocks) - 1; i++ {
    // ith keysize block
    blocks[i] = c[0:keysize]
    // Remove that slice
    c = c[keysize:]
  }
  // Add remainder
  blocks[len(blocks) - 1] = c
  return blocks
}
func transposeBlocks(blocks []string, keysize int) []string {
  tb := make([]string, keysize)
  for i := 0; i < len(keysize); i++ {
    // something with modulo?
    for j := 0; j < len(blocks); j+=keysize {
      if i % keysize == j {
        append(tb[i], blocks[j]
      }
    }
  }
  return tb
}
