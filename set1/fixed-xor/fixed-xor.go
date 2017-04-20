package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

func main() {
  result := readLines()
  //xor()
  fmt.Println(result)
}
func readLines() []int {
  reader := bufio.NewReader(os.Stdin)
  fmt.Println("Welcome to fixed-xor")
  str1, _ := reader.ReadString('\n')
  str2, _ := reader.ReadString('\n')
  if len(str2) != len(str1) {
    print("Error: input lengths must match")
    os.Exit(1)
  }
  // Remove newlines
  str1 = str1[:len(str1)-1]
  str2 = str2[:len(str2)-1]
  result := make([]int, len(str1))
  for pos, char:= range str1 {
    i, err := strconv.ParseInt(string(char), 16, 0)
    if err != nil {
      print("Error: ")
      fmt.Println(1, err)
    }
    j, err := strconv.ParseInt(string(str2[pos]), 16, 0)
    if err != nil {
      print("Error: ")
      fmt.Println(1, err)
    }
    result[pos] = int(i) ^ int(j)
  }
  return result
}
