package utils

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

// sign 签名
func Sign(data string, privkey string) (signature string, err error) {
	privateKey, err := crypto.HexToECDSA(privkey[2:])
	hash := crypto.Keccak256Hash([]byte(data))
	message := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", 32, hash.Bytes())
	hash = crypto.Keccak256Hash([]byte(message))

	signatureByte, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return
	}
	signatureByte[64] += 27
	signature = hexutil.Encode(signatureByte)
	return
}

func sha3Hash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", 32, data)
	return crypto.Keccak256([]byte(msg))
}

func VerifyMessage(message string, signedMessage string) (string, error) {
	hash := crypto.Keccak256Hash([]byte(message))
	// Hash the unsigned message using EIP-191
	hashedMessage := []byte("\x19Ethereum Signed Message:\n" + strconv.Itoa(32) + string(hash.Bytes()))
	hash = crypto.Keccak256Hash(hashedMessage)

	// Get the bytes of the signed message
	decodedMessage := hexutil.MustDecode(signedMessage)
	// Handles cases where EIP-115 is not implemented (most wallets don't implement it)
	if decodedMessage[64] == 27 || decodedMessage[64] == 28 {
		decodedMessage[64] -= 27
	}

	// Recover a public key from the signed message
	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), decodedMessage)
	if sigPublicKeyECDSA == nil {
		err = errors.New("Could not get a public get from the message signature")
	}
	if err != nil {
		return "", err
	}
	return crypto.PubkeyToAddress(*sigPublicKeyECDSA).String(), nil
}

// SigRSV signatures  V,R,S returned as arrays
func SigRSV(isig interface{}) (uint8, [32]byte, [32]byte) {
	var sig []byte
	switch v := isig.(type) {
	case []byte:
		sig = v
	case string:
		sig, _ = hexutil.Decode(v)
	}

	sigstr := common.Bytes2Hex(sig)
	rS := sigstr[0:64]
	sS := sigstr[64:128]
	R := [32]byte{}
	S := [32]byte{}
	copy(R[:], common.FromHex(rS))
	copy(S[:], common.FromHex(sS))
	vStr := sigstr[128:130]
	vI, _ := strconv.Atoi(vStr)
	V := uint8(vI)

	return V, R, S
}

func SignClaim4Evm(hash common.Hash, privkey string) ([]byte, error) {
	privateKey, err := crypto.HexToECDSA(privkey[2:])
	if err != nil {
		panic(err)
	}
	rawSignature, _ := prefixMessage(hash, privateKey)
	signature := hexutil.Bytes(rawSignature)
	return signature, nil
}

func prefixMessage(message common.Hash, key *ecdsa.PrivateKey) ([]byte, []byte) {
	prefixed := solsha3.SoliditySHA3WithPrefix(message[:])
	//prefixed := SoliditySHA3WithPrefix(message[:])
	sig, err := secp256k1.Sign(prefixed, math.PaddedBigBytes(key.D, 32))
	if err != nil {
		panic(err)
	}
	return sig, prefixed
}

// SigRSV signatures  V,R,S returned as arrays
//
//	func GetSigRSV(types []string, values []interface{}, privkey string) (uint8, [32]byte, [32]byte) {
//		initdata := solsha3.SoliditySHA3(types, values)
//		signature, _ := SignClaim4Evm(common.BytesToHash(initdata), privkey)
//		return SigRSV(signature)
//	}
func GetSigRSV(types []string, values []interface{}, privkey string) []byte {
	initdata := solsha3.SoliditySHA3(types, values)
	signature, _ := SignClaim4Evm(common.BytesToHash(initdata), privkey)
	return signature
}

// V, R, S := utils.GetSig(types, values, privkeySigner)
// RR := "0x" + hex.EncodeToString(R[:])
// SS := "0x" + hex.EncodeToString(S[:])
// _battleinfo := struct {
// RR      string
// SS      string
// VV      uint8
// }{
// RR,
// SS,
// V,
// }
func GetSig(types []string, values []interface{}, privkey string) (uint8, [32]byte, [32]byte) {
	initdata := solsha3.SoliditySHA3(types, values)
	signature, _ := SignClaim4Evm(common.BytesToHash(initdata), privkey)
	return SigRSV(signature)
}
