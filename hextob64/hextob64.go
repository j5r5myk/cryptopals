package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {

  a := []int{1, 2, 3}
  fmt.Printf("%v\n", a)

  hex := readHexLine()
  //b64 = hexToB64(hex)
  //fmt.Print(b64)
  fmt.Printf("%v %T\n", hex, hex);
}

func readHexLine() []int {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter hex: ")
  text, _ := reader.ReadString('\n')
  // test valid hex
  /*
  if !isHex(text) {
    fmt.Println("Error input must be hexadecimal");
    os.Exit(1);
  }
  */
  // convert to int and store
  hex := make([]int, len(text))
  for i := range hex {
    hex[i] = text[i] - 48
  }
  fmt.Printf("type: %T\n", text)
  return hex
}

func isHex(hex string) bool {
  fmt.Println("string length: ", len(hex))
  for i := 0; i < len(hex) - 1; i++ {
    if (hex[i] < '0' || hex[i] > '9') {
      return false
    }
    if (hex[i] < 'a' || hex[i] > 'f') {
      return false
    }
  }
  return true
}

func hexToB64(hex string) {
  // read 3 octets
  // convert into 3 b64 chars
  for i := 0; i < len(hex); i+=3 {
    //hex[i]+hex[i+1]+hex[i+2]
  }
}
