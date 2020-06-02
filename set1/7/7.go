/* Cryptopals Challenge #7
 * AES in ECB mode
 * Author: j5r5myk
 */
package main

import (
  "crypto/aes"
  "encoding/base64"
  "fmt"
  "io/ioutil"
  "encoding/hex"
)

func main() {
    KEY := []byte("YELLOW SUBMARINE")
    // Open file
    ciphertext, err := ioutil.ReadFile("/Users/xje/go/src/cryptopals/input/7.txt")
    if err != nil {
      fmt.Print(err)
    }
    // Decode b64
    cipherDecoded, err := base64.StdEncoding.DecodeString(string(ciphertext))
    if err != nil {
      fmt.Print(err)
    }
    fmt.Printf("%s\n", hex.Dump(cipherDecoded))
    plaintext := decrypt(cipherDecoded, KEY)
    if err != nil {
      fmt.Println(err)
    }
    fmt.Printf("%s\n", plaintext)
}
func decrypt(data, key []byte) []byte {
    cipher, _ := aes.NewCipher([]byte(key))
    plaintext := make([]byte, len(data))
    println(len(data))
    size := 16
    // Decrypt a block at a time
    for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
      cipher.Decrypt(plaintext[bs:be], data[bs:be])
    }

    return plaintext
}
