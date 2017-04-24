package main

import (
  "fmt"
  "encoding/hex"
  "b64"
)

func main() {
  MIN_KEYSIZE := 2
  MAX_KEYSIZE := 40
  //str1 := "this is a test"
  //str2 := "wokka wokka!!!"
  
  // Read input
  // convert from b64
  // run this:

  //input := "THISISATEST"
  // Find key size
  keysize := findKeySize(MIN_KEYSIZE, MAX_KEYSIZE, input)
  fmt.Printf("Likely key size: %d\n", keysize)
  // Divide input into keySized blocks
  blocks := createBlocks(input, keysize)
  tb := transposeBlocks(blocks, keysize)
  // Print blocks
  println("Blocks:")
  for pos, block := range blocks {
    fmt.Printf("%d: %s\n", pos, block)
  }
  println("Transposed:")
  for pos, tblock := range tb {
    fmt.Printf("%d: %s\n", pos, tblock)
  }
  key := ""
  // Single XOR tranposed blocks
  for i := 0; i < keysize; i++ {
    key += singleXOR(tb[i])
  }
  fmt.Printf("key: %s\n", key)
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
  // Each array element is keysize # bytes
  // len(s) / keysize + 1 elements for remainder
  blocks := make([]string, (len(c) / keysize) + 1)
  for i := 0; i < len(blocks) - 1; i++ {
    // ith keysize block
    blocks[i] = c[0:keysize]
    // Remove processed slice
    c = c[keysize:]
  }
  // Add remainder
  blocks[len(blocks) - 1] = c
  return blocks
}
func transposeBlocks(blocks []string, keysize int) []string {
  tb := make([]string, keysize)
  for i := 0; i < len(blocks); i++ {
    for j := 0; j < len(blocks[i]); j++ {
      //fmt.Printf("[%d][%d]\n",i, j)
      tb[j] += string(blocks[i][j])
    }
  }
  return tb
}
func singleXOR(input string) string {
  reverseLetters := "zjqxkvbpgwyfmculdhrsnioate"
  hiscore := 0
  hichar := ""
  var score int
  hiholder := make([]byte, len(input) / 2)
  // Create hashmap
  freq := make(map[byte]int)
  for pos, char := range reverseLetters {
    freq[byte(char)] = pos
  }
  for i := 0; i < 128; i++ {
    // Decode using the given byte and calculate frequency score
    decoded := decode(byte(i), input)
    score = calcScore(decoded, freq)
    // Greedily compare to previous attempts
    if score > hiscore {
      hiscore = score
      hichar = string(i)
      copy(hiholder, decoded)
    }
  }
  fmt.Printf("Winner:\nChar: %s Score: %d\n%s", hichar, hiscore, hex.Dump(hiholder))
  return hichar
}
func decode(b byte, input string) []byte {
  result := make([]byte, len(input))
  for i := 0; i < len(input); i++ {
     result[i] = input[i] ^ b
  }
  return result
}
func calcScore(s []byte, freq map[byte]int) int {
  score := 0
  for _, char := range s {
    score += freq[char]
  }
  return score
}