package main

import (
  "fmt"
  "encoding/hex"
  "os"
  "bufio"
  "io/ioutil"
  "encoding/base64"
)

func main() {
  MIN_KEYSIZE := 2
  MAX_KEYSIZE := 40
  // Open file
  file, err := os.Open("6.txt")
  if err != nil {
    os.Exit(1)
  }
  // Read 1 line for keysize calculations
  defer file.Close()
  scanner := bufio.NewScanner(file)
  scanner.Scan()
  line := scanner.Text()
  fmt.Println(line)
  // Convert from b64
  //input := []byte("ae00ff1235889901fee11247981275981725987febc7ca7a7c7a8c79a87ca8fe8b9c0a7e635a412")
  lineD, err := base64.StdEncoding.DecodeString(line)
  fmt.Printf("Decoded: %s\n", lineD)
  fmt.Printf("Decoded: %v\n", []byte(lineD))
  for i := 0; i < len(lineD); i++ {
    fmt.Printf("%x ", lineD[i])
  }
  println()
  fmt.Printf("%s\n", hex.Dump([]byte(lineD)))
  // Find key size
  keysize := findKeySize(MIN_KEYSIZE, MAX_KEYSIZE, lineD)
  fmt.Printf("Likely key size: %d\n", keysize)
  // Read whole file into memory
  input, err := ioutil.ReadFile("6.txt")
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
    key += string(solveSingleXOR([]byte(tb[i])))
  }
  // Print likely key
  fmt.Printf("key: %v (%s)\n", []byte(key), key)
  // Then actually decrypt it... 
  scanner = bufio.NewScanner(file)
  for scanner.Scan() {
    decodeSingleXOR(key, scanner.Text())
  }
}
func findKeySize(min int, max int, c []byte) int {
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
func calcHamming(str1 []byte, str2 []byte) int {
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
func createBlocks(c []byte, keysize int) [][]byte{
  // Each array element is keysize # bytes
  // len(s) / keysize + 1 elements for remainder
  blocks := make([][]byte, (len(c) / keysize) + 1)
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
func transposeBlocks(blocks [][]byte, keysize int) []string {
  tb := make([]string, keysize)
  for i := 0; i < len(blocks); i++ {
    for j := 0; j < len(blocks[i]); j++ {
      //fmt.Printf("[%d][%d]\n",i, j)
      tb[j] += string(blocks[i][j])
    }
  }
  return tb
}
func solveSingleXOR(input []byte) int {
  reverseLetters := "zjqxkvbpgwyfmculdhrsnioate"
  hiscore := 0
  hichar := 0
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
      hichar = i
      copy(hiholder, decoded)
    }
  }
  fmt.Printf("Winner:\nChar: %s (%d) Score: %d\n%s", string(hichar), hichar, hiscore, hex.Dump(hiholder))
  return hichar
}
func decode(b byte, input []byte) []byte {
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
func decodeSingleXOR(key string, line string) []byte {
  result := make([]byte, len(line))
  for i := 0; i < len(line); i++ {
    result[i] = line[i] ^ key[i%len(key)]
  }
  return result
}
