/* Cryptopals Challenge #4
 * Detect single-character XOR
 * Author: j5r5myk
 */
package main

import (
  "fmt"
  "encoding/hex"
  "strconv"
  "bufio"
  "os"
)

func main() {
  /* Init vars */
  var LINE_LENGTH int = 60
  reverseLetters := "zjqxkvbpgwyfmculdhrsnioate"
  hiScore := 0
  lineNum := 0
  hiChar := ""
  var score int
  hiHolder := make([]byte, LINE_LENGTH / 2)
  // Create hashmap
  freq := make(map[byte]int)
  for pos, char := range reverseLetters {
    freq[byte(char)] = pos
  }
  // Read line by line 
  file, err := os.Open("./4.txt")
  if err != nil {
    fmt.Println(1, err)
  }
  defer file.Close()
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    // Parse pairs into bytes
    bytes := parseHexString(scanner.Text())
    for i := 0; i < 128; i++ {
      // Decode using i and calculate frequency score
      decoded := decode(byte(i), bytes)
      score = calcScore(decoded, freq)
      // Greedily compare to previous attempts
      if score > hiScore {
        hiScore = score
        hiChar = string(i)
        copy(hiHolder, decoded)
      }
    }
    lineNum++
  }
  fmt.Printf("Winner:\nLine #%d Char: %s Score: %d\n%s", lineNum, hiChar, hiScore, hex.Dump(hiHolder))
  defer file.Close()
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
