/* Cryptopals Challenge #10
 * Implement CBC mode
 * Author: j5r5myk
 */
package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "log"
)

func main() {
    const size := 16
    const IV := "\x00\x00\x00"
    // divide m into 16 byte blocks
    // c = ECB(IV xor block0)
    // ECB(c xor block1) until entire message is encrypted
    
    // Open plaintext file
    if len(os.Args) < 2 {
        fmt.Printf("Usage: 10 <path-to-plaintext-file>\n")
        os.Exit(0)
    }
    plaintext, err := ioutil.ReadFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
        os.Exit(0)
    }
    fmt.Printf("%s", plaintext)
    // Init
    ecb_encrypt(myXOR(IV, plaintext[0:16]))
    // Loop through remaining blocks
    for bs, be := 16, size+16; bs < len(plaintext); bs, be = bs+size, be+size {
        
    }
}
/*
func ecb_encrypt(message, key []byte) []byte {
    cipher, _ := aes.NewChiper([]byte(key))
    ciphertext := make([]byte, len(message))
    size := 16
    // Encrypt a block at a time
    for bs, be := 16, size+16; bs < len(data); bs, be = bs+size, be+size {
      cipher.Encrypt(message[bs:be], ciphertext[bs:be])
    }
    return ciphertext
}
*/
