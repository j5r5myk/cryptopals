package main

import (
  "fmt"
  "encoding/hex"
  "strconv"
)

func main() {
  input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
  letters := "etaoinsrhdlucmfywgpbvkxqjz"
  chars := "abcdefghijklmnopqrstuvqxyz`1234567890-=~!@#$%^&*()_+[]\\{}|;':\",./<>?"
  // Arbitrary large value
  lowscore := 100000
  var score int
  lowholder := make([]byte, len(input) / 2)
  // Create hashmap
  freq := make(map[byte]int)
  for pos, char := range chars {
    freq[byte(char)] = pos
  }
  bytes := parseHexString(input)
  for _, char := range letters {
    decoded := decode(byte(char), bytes)
    score = calcScore(decoded, freq)
    //fmt.Printf("Char: %r, Score: %d\n%s", char, score, hex.Dump(decoded))
    // Greedily compare to previous attempts
    if score < lowscore {
      lowscore = score
      copy(lowholder, decoded)
    }
  }
  fmt.Printf("Winner:\n%s", hex.Dump(lowholder))
}
func parseHexString(s string) []byte {
  output := make([]byte, len(s) / 2)
  j := 0
  for i := 0; i < len(s); i+=2 {
    hibits, err := strconv.ParseInt(string(s[i]), 16, 0)
    lowbits, err := strconv.ParseInt(string(s[i+1]), 16, 0)
    if err != nil {
      fmt.Println(1, err)
    }
    b := (byte(hibits) << 4 | byte(lowbits))
    output[j] = b
    j++
  }
  return output
}
func decode(b byte, bytes []byte) []byte {
  result := make([]byte, len(bytes))
  for pos, char := range bytes {
     result[pos] = char ^ b
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
