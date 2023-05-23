package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
)

func generateKeyPair() (*rsa.PublicKey, *rsa.PrivateKey) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publicKey := &privateKey.PublicKey
	return publicKey, privateKey
}

func sign(message string, privateKey *rsa.PrivateKey) string {
	hash := sha256.New()
	hash.Write([]byte(message))
	digest := hash.Sum(nil)
	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, digest, nil)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(signature)
}

func verify(message, signature string, publicKey *rsa.PublicKey) bool {
	decodedSignature, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false
	}
	hash := sha256.New()
	hash.Write([]byte(message))
	digest := hash.Sum(nil)
	err = rsa.VerifyPSS(publicKey, crypto.SHA256, digest, decodedSignature, nil)
	return err == nil
}

func main() {
	publicKey, privateKey := generateKeyPair()
	message := "Hello World!"
	signature := sign(message, privateKey)
	isValid := verify(message, signature, publicKey)
	println(isValid) // Prints true
}
