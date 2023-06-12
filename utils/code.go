package utils

import (
	"math/rand"
	"time"
)

func randomKey2(length, level int) (strPol string) {
	var slicePol []byte
	var str string
	switch level {
	case 1:
		str = "0123456789"
	default:
		str = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		n := rand.Intn(len(str))
		slicePol = append(slicePol, str[n])
	}
	strPol = string(slicePol)
	return
}

func GetCode() string {
	return randomKey2(6, 1)
}

func GetInvitCode() string {
	return randomKey2(6, 2)
}
