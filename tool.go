package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func Sha256(src string) string {
	m := sha256.New()
	m.Write([]byte(src))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

func toJSONString(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}
	return string(b)
}
