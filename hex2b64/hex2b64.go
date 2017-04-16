package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

func main() {
 // hexstr := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
 ints := readHexLine()
 fmt.Println(ints)
 fmt.Println(hex2b64(ints))
}

func readHexLine() []byte {
  // Read line of text
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter hex to convert: ")
  hexstr, _ := reader.ReadString('\n')
  fmt.Println("Input: ", hexstr)
  // Remove newline
  hexstr = hexstr[:len(hexstr)-1]
  // TODO: test valid hex
    ints := make([]byte, len(hexstr) / 2)
    for i := 0; i < len(hexstr) / 2; i+=2 {
      hibits, err := strconv.ParseInt(string(hexstr[i]), 16, 0)
      lowbits, err := strconv.ParseInt(string(hexstr[i+1]), 16, 0)
      if err != nil {
        fmt.Println(1, err)
      }
      // Create the byte
      // TODO prevent out of bounds
      b := (byte(hibits) << 4) | byte(lowbits)
      fmt.Printf("%b\n", b)
      if i == 0 {
        ints[i] = b
      } else {
        ints[i-1] = b
      }
    }
  return ints
}

/*
A lil map
Bytes:  11111111 22222222 3333333
b64s:   11111122 22223333 33444444 
*/
func hex2b64(ints []byte) string {
  dict := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="
  var b64 string
  var b byte
  // process 3 bytes at a time
  for i := 0; i < len(ints); i+=3 {
    // first 6 bits
    b = ints[i] & 0xFC >> 2
    fmt.Printf("Byte %d: %b\n", i, ints[i])
    b64 = b64 + string(dict[b])
    // next 6
    b = ints[i] & 0x03 << 4
    if i + 1 < len(ints) {
      b = b | (ints[i+1] & 0xF0 >> 4)
      fmt.Printf("Byte %d: %b\n", i+1, ints[i+1])
      b64 = b64 + string(dict[b])
      // 6 more
      b = ints[i+1] & 0x0F << 2
      if i + 2 < len(ints) {
        b = b | ints[i+2] & 0xC0 >> 4
        // last 6
        b = ints[i+2] & 0x3F
      } else {
        b64 = b64 + "="
      }
    } else {
      b64 = b64 + "=="
    }
  }
  println()
  return b64
}
