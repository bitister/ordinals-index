package utils

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"

	"github.com/astaxie/beego"
	"github.com/mervick/aes-everywhere/go/aes256"
)

func GET_CPUID() (string, error) {
	cmd := exec.Command("/bin/sh", "-c", `sudo dmidecode -t 4 | grep ID `)
	// stdout, err := cmd.StdoutPipe()
	// if err != nil {
	// 	beego.Error("StdoutPipe: " + err.Error())
	// 	return "", err
	// }

	stderr, err := cmd.StderrPipe()
	if err != nil {
		beego.Error("StderrPipe: ", err.Error())
		return "", err
	}

	if err := cmd.Start(); err != nil {
		beego.Error("Start: ", err.Error())
		return "", err
	}

	bytesErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		beego.Error("ReadAll stderr: ", err.Error())
		return "", err
	}

	if len(bytesErr) != 0 {
		beego.Error("stderr is not nil: %s", bytesErr)
		return "", err
	}

	// bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		beego.Error("ReadAll stdout: ", err.Error())
		return "", err
	}

	if err := cmd.Wait(); err != nil {
		beego.Error("Wait: ", err.Error())
		return "", err
	}

	// return strings.TrimSpace(strings.ReplaceAll(string(bytes), " ", "")), err
	return "123456", nil
}

func EncryptWithAES(key, message string) string {
	return aes256.Encrypt(message, key)
}

func DecryptWithAES(key, message string) string {
	return aes256.Decrypt(message, key)
}

func Zip(src string) []byte {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	r := []byte(src)
	if _, err := gz.Write(r); err != nil {
		panic(err)
	}

	if err := gz.Flush(); err != nil {
		panic(err)
	}

	if err := gz.Close(); err != nil {
		panic(err)
	}

	return b.Bytes()
}

func Unzip(src []byte) string {
	rdata := bytes.NewReader(src)
	r, _ := gzip.NewReader(rdata)
	s, _ := ioutil.ReadAll(r)
	return string(s)
}

func GetPrivkeyAdmin() string {
	/**
	 * 获取电脑CPUId
	 */

	str, _ := GET_CPUID()

	//解押文件
	//读取文件全部内容
	b, err := ioutil.ReadFile("/root/.key/keyadmin")
	// b, err := ioutil.ReadFile(keypath)
	if err != nil {
		beego.Error("ioutil ReadFile error: ", err)
	}
	//解密
	h := sha256.New()
	h.Write([]byte(str))
	hash := fmt.Sprintf("%x", h.Sum(nil))
	//content := DecryptWithAES("123456", zip.Unzip([]byte(b)))
	content := DecryptWithAES("gpJ1SXT]f~nU^c793Al[L-eD%WWAY2", Unzip([]byte(b)))
	res := DecryptWithAES("gpJ1SXT]f~nU^c793Al[L-eD%WWAY2"+hash, content)
	return res
}

func GetPrivkeySign() []string {
	/**
	 * 获取电脑CPUId
	 */
	str, _ := GET_CPUID()

	//解押文件
	//读取文件全部内容
	b, err := ioutil.ReadFile("/root/.key/keysign")
	// b, err := ioutil.ReadFile(keypath)
	if err != nil {
		beego.Error("ioutil ReadFile error: ", err)
	}
	//解密
	h := sha256.New()
	h.Write([]byte(str))
	hash := fmt.Sprintf("%x", h.Sum(nil))
	//content := DecryptWithAES("123456", zip.Unzip([]byte(b)))
	content := DecryptWithAES("gpJ1SXT]f~nU^c793Al[L-eD%WWAY2", Unzip([]byte(b)))
	res := DecryptWithAES("gpJ1SXT]f~nU^c793Al[L-eD%WWAY2"+hash, content)
	return strings.Split(res, ",")
}
