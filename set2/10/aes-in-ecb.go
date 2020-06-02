/* Cryptopals Challenge #
 * Implement AES in ECB mode
 * Author: j5r5myk
 */

package main

import (
    "crypto/aes"
    "encoding/base64"
    //"os"
    "fmt"
    "io/ioutil"
    "encoding/hex"
)  

func main() {
    // Hard coded key
    KEY := []byte("YELLOW SUBMARINE")
    // Open file
    ciphertext, err := ioutil.ReadFile("7.txt")
    if err != nil {
      fmt.Print(err)
    }
    // Decode b64
    cipherDecoded, err := base64.StdEncoding.DecodeString(string(ciphertext))
    if err != nil {
      fmt.Print(err)
    }
    // Hex dump cipher text
    fmt.Printf("%s\n", hex.Dump(cipherDecoded))
    // Decrypt
    plaintext := decrypt(cipherDecoded, KEY)
    if err != nil {
      fmt.Println(err)
    }
    fmt.Printf("%s\n", plaintext)
}
func decrypt(data, key []byte) []byte {
    cipher, _ := aes.NewCipher([]byte(key))
    plaintext := make([]byte, len(data))
    // Debug print length
    println(len(data))
    size := 16
    // Decrypt a block at a time
    for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
      cipher.Decrypt(plaintext[bs:be], data[bs:be])
    }

    return plaintext
}
