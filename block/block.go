package blockchain

import (
	"DataCertProject/util"
	"bytes"
	"encoding/gob"
	"time"
)

type Block struct {
	Height int64 //区块高度
	TimeStamp int64//时间
	Hash []byte//区块hash
	Data []byte//数据
	 PrevHash []byte//上一个区块哈希
	Version string// 版本号
	Nonce int64// 随机数 nonce
}

/*
 * 生产创世区块
 */
func CreatGenesisBlock()Block{
	block:=NewBlock(0,[]byte{},[]byte{0,0,0,0})
	return block
}

func NewBlock(height int64 , data []byte ,prevHas []byte)Block{
	// 1.构建一个新的区块
	block :=Block{
		Height:   height,
		TimeStamp: time.Now().Unix(),
		Data: data,
		PrevHash:  prevHas,
		Version:   "0x01",
	}
	// 2. 生产新的block，寻找 nonce
	// 并将block 的nonce 赋值
	suiji := NewPoW(block)
	blockHahs,nonce := suiji.Run()

	block.Nonce =nonce
	block.Hash = blockHahs
	//转化为 字节切片
	heightBytes,_ :=util.IntToBytes(block.Height)
	TimeBytes,_ :=util.IntToBytes(block.TimeStamp)
	NonceBytes,_ :=util.IntToBytes(block.Nonce)
	VersionBytes :=util.StringToBytes(block.Version)
	//拼接起来
	blockBytes :=bytes.Join([][]byte{
		heightBytes,
		TimeBytes,
		data,
		prevHas,
		NonceBytes,
		VersionBytes,
	},[]byte{})

	//设置区块的hash
	block.Hash =util.SHA256Hash(blockBytes)
	return block
}

// 序列化
func(bk Block)Serialize()([]byte ,error){
	buff :=new(bytes.Buffer)
	err :=gob.NewEncoder(buff).Encode(bk)
	if err !=nil{
		return nil , err
	}
	return buff.Bytes(),nil
}
// 反序列化
func DeSerialize (data []byte)(*Block , error){
	var block Block
	err :=gob.NewDecoder(bytes.NewReader(data)).Decode(&block)
	if err !=nil{
		return nil,err
	}
	return  &block ,nil
}