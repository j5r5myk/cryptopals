/* Cryptopals Challenge #10
 * Implement CBC mode
 * Author: j5r5myk
 */
package main

import (
  "fmt"
  "os"
)

func main() {
  // IV = /x00/x00/x00
  // encrypt m with ecb
  // xor those?
}
func ecb_encrypt(message, key []byte) []byte {
  cipher, _ := aes.NewChiper([]byte(key))
  ciphertext := make([]byte, len(message))
  size := 16
  // I copied this from Stackoverflow and don't understand it
  for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
    cipher.Encrypt(message[bs:be], ciphertext[bs:be])
  }
  return ciphertext
}
