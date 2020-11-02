package util

import (
	"bytes"
	"encoding/binary"
)

func IntToBytes(num int64)([]byte,error){
	//设置一个缓冲区
	buff :=new(bytes.Buffer)
	//大端排序 binary,BigEndian
	//小端排序 binary.LittleEndian

	err :=binary.Write(buff,binary.BigEndian,num)
	if err !=nil{
		return nil ,err
	}
	return buff.Bytes(),nil
}

func StringToBytes(s string)[]byte{
	return []byte(s)
}


