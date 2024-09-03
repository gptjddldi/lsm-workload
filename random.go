package lsm_workload

import (
	"math/rand"
	"strconv"
)

const (
	// CharEnglishAlphabetLowercase = "abcdefghijklmnopqrstuvwxyz"
	CharEnglishAlphabetNumber = "abcdefghijklmnopqrstuvwxyz0123456789"
	CharNumber                = "0123456789"
	CharBase62                = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

type RandomString struct {
	Charset string
}

// mode: "english", "number", ""
func NewRandomString(mode string) *RandomString {
	charset := ""
	switch mode {
	case "english":
		charset = CharEnglishAlphabetNumber
	case "number":
		charset = CharNumber
	default:
		charset = CharBase62
	}

	return &RandomString{
		Charset: charset,
	}
}

func (r *RandomString) RandomKey() string {
	length := 1

	switch r.Charset {
	case CharEnglishAlphabetNumber:
		length += rand.Intn(12)
	case CharBase62:
		length += rand.Intn(100)
	case CharNumber:
		return r.generateRandomInt()
	}

	return r.generateRandomStringWithLength(r.Charset, length)
}

func (r *RandomString) RandomValue() string {
	length := rand.Intn(40) + 1
	return r.generateRandomStringWithLength(CharBase62, length)
}

func (r *RandomString) generateRandomStringWithLength(charset string, length int) string {
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = charset[rand.Intn(len(charset))]
	}
	return string(bytes)
}

func (r *RandomString) generateRandomInt() string {
	num := rand.Int63n(1e12) // 0부터 999,999,999,999까지의 숫자 생성
	numStr := strconv.FormatInt(num, 10)

	for len(numStr) < 12 {
		numStr += "0"
	}
	return numStr
}
