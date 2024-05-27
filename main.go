package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"fmt"
	"os"
)

const Key = "1beach2sun3waves4ocean!"

type Command interface {
	Parse(args []string)
	Run()
}

func main() {
	if len(os.Args) < 2 {
		DieUsage()
	}

	var cmd Command

	switch os.Args[1] {
	case "encode":
		cmd = &Encode{}
	case "compress":
		cmd = &Compress{}
	case "decompress":
		cmd = &Decompress{}
	default:
		DieUsage()
	}
	cmd.Parse(os.Args[2:])
	cmd.Run()
}

func DieUsage() {
	fmt.Fprintf(os.Stderr, "Usage: %s (encode | compress | decompress) [flags]\n", os.Args[0])
	os.Exit(1)
}

func Must(e error) {
	if e != nil {
		panic(e)
	}
}

func AESKey() []byte {
	return Hash([]byte(Key))[:16]
}

func Hash(data []byte) []byte {
	hash := sha1.New()
	hash.Write(data)
	return hash.Sum(nil)
}

func Encrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(AESKey())
	if err != nil {
		return nil, err
	}

	cipher, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Use a deterministic nonce (I'm not actually sure this is safe!)
	// The goal here is to make encoding deterministic.
	nonce := Hash(data)[:cipher.NonceSize()]

	cipherText := cipher.Seal(nil, nonce, data, nil)

	return append(nonce, cipherText...), nil
}

func Decrypt(encryptedData []byte) ([]byte, error) {
	block, err := aes.NewCipher(AESKey())
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	nonce, cipherText := encryptedData[:nonceSize], encryptedData[nonceSize:]

	return aesGCM.Open(nil, nonce, cipherText, nil)
}
