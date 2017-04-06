package main

import (
  "bufio"
  "fmt"
  "os"
  "encoding/hex"
)

func main() {
 // hexstr := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
 rawhex := readHexLine()
 fmt.Println(rawhex)
 printHexArray(rawhex)
 
 fmt.Println(hex2b64(rawhex))
}

func printHexArray(a []byte) {
  for i := 0; i < len(a); i++ {
    fmt.Printf("%X", a[i])
  }
  fmt.Println()
}

func readHexLine() []byte {
  // Read line of text
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter hex to convert: ")
  hexstr, _ := reader.ReadString('\n')
  // test valid hex
  if !isEvenLen(hexstr) {
    fmt.Println("Error: input must be even");
    os.Exit(1)
  }
  rawhex, err := hex.DecodeString(hexstr[:len(hexstr)-1])
  if err != nil {
    panic(err)
  }
  return rawhex
}

func isEvenLen(hexstr string) bool {
  if (len(hexstr)-1)%2 != 0 {
    return false
  }
  fmt.Println("hex length: ", len(hexstr)-1)
  return true
}

func isHex(rawhex []byte) bool {
  for i := 0; i < len(rawhex); i++ {
    if (rawhex[i] < '0') || (rawhex[i] > 'f') {
      fmt.Println(rawhex[i])
      return false
    }
  }
  return true
}

func hex2b64(rawhex []byte) string {
  // read 3 octets
  // convert into 3 b64 chars
  for i := 0; i < len(rawhex); i+=3 {
    //hex[i]+hex[i+1]+hex[i+2]
  }
  return "bunk!"
}
