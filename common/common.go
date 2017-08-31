package main

import (
  "fmt"
)

func findKeySize(min int, max int, c []byte) []int {
  lowHam := 1000
  bestSize := 0
  secondBest := 0
  thirdBest := 0
  result := make([]int, 3)
  // Check max vs input length
  /*
  if max > len(c) / 2 {
    max = len(c) / 2
    fmt.Printf("Max key length longer than single line, shortening to %d\n", max)
  }
  */
  for i := min; i <= max; i++ {
    fmt.Println("Testing keysize: ", i)
    // Hamming distance b/w first 4 chunks
    //fmt.Printf("%v %v\n", c[0:i], c[i:2*i])
    //rawHam := calcHamming(c[0:i], c[i:2*i])
    rawHam := calcHamming(c[0:i], c[i:2*i])
    // Normalize and compare
    normHam := rawHam / i
    println("normHam: ", normHam)
    if normHam < lowHam {
      lowHam = normHam
      thirdBest = secondBest
      secondBest = bestSize
      bestSize = i
    }
  }
  result[0] = bestSize
  result[1] = secondBest
  result[2] = thirdBest
  return result
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
//  fmt.Printf("Winner:\nChar: %s (%d) Score: %d\n%s", string(hichar), hichar, hiscore, hex.Dump(hiholder))
  return hichar
}
*/

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
*/
func decodeSingleXOR(key string, line string) []byte {
  result := make([]byte, len(line))
  for i := 0; i < len(line); i++ {
    result[i] = line[i] ^ key[i%len(key)]
  }
  return result
}

