package util

import (
	"crypto/md5"
	"encoding/hex"
)

func HashMd5String(){
	hashMd5:= md5.New() // 实例化一个 hash
	hashMd5.Write([]byte(u.Pwd))
	bytes :=hashMd5.Sum(nil)
	u.Pwd = hex.EncodeToString(bytes)
}