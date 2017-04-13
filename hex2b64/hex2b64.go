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

/*
A lil map
Bytes:  11111111 22222222 3333333
b64s:   11111122 22223333 33444444 
*/
func hex2b64(ints []int) string {
  dict := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="
  var b64 string
  var b int
  // process 3 bytes at a time
  for i := 0; i < len(ints); i+=3 {
    // first 6 bits
    b = ints[i] & 0xFC >> 2
    b64 = b64 + dict[b]
    // next 6
    b = ints[i] & 0x03 << 4
    b = b | (ints[i+1] & 0xF0 >> 4)
    b64 = b64 + dict[b]
    // 6 more
    b = ints[i+1] & 0x0F << 2
    b = b | ints[i+2] & 0xC0 >> 4
    b64 = b64 + dict[b]
    // last 6
    b = ints[i+2] & 0x3F
    b64 = b64 + dict[b]
  }
  return b64
}
