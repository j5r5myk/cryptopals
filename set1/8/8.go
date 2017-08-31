/* Cryptopals Challenge #8
 * Detect AES in ECB mode
 * Author: j5r5myk
 */
package main

import (
  "fmt"
  "os"
  "bufio"
)

func main() {
  /*  Detect AES ECB.
   *  same 16 byte plaintext ALWAYS gives same 16 byte ciphertext
   *
   *  Do we have the same plaintext twice?
   *  If it was the same 16 byte key every time, you could xor 2 ciphers
   *  To get the key, I think that's a previous exercise though...
   */
  // For each cipher
  // Divide into blocks
  // Compare blocks each block with eachother for n^2 goodness

  file, err := os.Open("8.txt")
  if err != nil {
    os.Exit(1)
  }
  defer file.Close()
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    cipher := scanner.Text()
    for i := 0; i < len(cipher) / 16; i++ {
      b := cipher[i*16:(i+1)*16]
      for j := i+1; j < len(cipher) / 16; j++ {
        if b == cipher[j*16:(j+1)*16] {
          fmt.Printf("%s\n", "Winner!:")
          fmt.Printf("%s\n", cipher)
        }
      }
    }
  }
}
