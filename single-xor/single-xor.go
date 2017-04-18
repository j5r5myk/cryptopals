package main

import (
  "fmt"
  //"os"
  "strconv"
)

func main() {
  input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
  letters := "etaoinsrhdlucmfywgpbvkxqjz"
  // Create hashmap
  freq := make(map[string]int)
  for pos, char := range letters {
    freq[char] = pos
  }
  bytes := parseHexString(input)
  for pos, char := range input {
    decoded := decode(char, bytes)
    score := score(decoded, freq)
    fmt.Printf("Decoded: %v, Score: %d\n", decoded, score)
    // compare to previous attempts greedily
  }
  // output winner
}
func parseHexString(s string) []byte {
  outpout := make([]byte, len(s) / 2)
  j := 0
  for i := 0; i < len(s); i++ {
    hibits, err := strconv.ParseInt(string(s[i]), 16, 0)
    lowbits, err := strconv.ParseInt(string(s[i+1), 16, 0)
    if err != nil {
      fmt.Println(1, err)
    }
    b := (byte(hibits) << 4 | byte(lowbits))
    output[j] = b
  }
  return output
}
func decode(r rune, bytes []byte) []byte {
  result := make([]byte, len(bytes))
  for pos, char := range bytes {
     result[pos] = char ^ r
  }
  return result
}
func score(s []byte, freq map[string]int) int {
  score := 0
  for char, pos := range s {
    score += freq[char]
  }
  return score
}
