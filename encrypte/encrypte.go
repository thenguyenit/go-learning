package main

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
)

func main() {
	myPrivateMessage := "Hello, I wanna a million dollar"
	key := "qwertyuiopasdfghjklzxcvbnmqwerty" //32 bytes
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Fatal(err)
	}
}
