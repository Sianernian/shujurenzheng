package util

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

// MD5 hash

func HashMd5String(data string) string{
	hashMd5:= md5.New() // 实例化一个 hash
	hashMd5.Write([]byte(data))
	bytes :=hashMd5.Sum(nil)
	return hex.EncodeToString(bytes)
}

//sha256 hash

func SHA256Hash(data []byte) ([]byte) {
	//、对数据进行sha256
	sha256Hash := sha256.New()
	sha256Hash.Write(data)
	return sha256Hash.Sum(nil)
}