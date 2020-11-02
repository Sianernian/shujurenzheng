package blockchain

import (
	"DataCertProject/util"
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

const DIFFFICULTY =2

type ProofOfWork struct {
	//目标值
	Target *big.Int
	//工作量证明值
	Block Block
}
/*
 * 实例化一个pow算法实例
 */
func NewPoW(block Block) ProofOfWork{
	target:=big.NewInt(1) //设置初始值
	target.Lsh(target,255-DIFFFICULTY) //左移
	pow :=ProofOfWork{
		Target:target,
		Block:block,
	}
	return pow
}
/*
 * pow算法：寻找符合条件的nonce值
 */
func (p ProofOfWork)Run()([]byte ,int64){
	var nonce int64
	bigBlock :=new(big.Int)
	var block256Hash []byte
	count :=0
	for {
		count++
		block :=p.Block
		heightBytes,_ :=util.IntToBytes(block.Height)
		timeBytes,_ :=util.IntToBytes(block.TimeStamp)
		nonceBytes,_ :=util.IntToBytes(nonce)
		versionBytes :=util.StringToBytes(block.Version)

		//拼接起来
		blockBytes := bytes.Join([][]byte{
			heightBytes,
			timeBytes,
			block.Data,
			block.PrevHash,
			versionBytes,
			nonceBytes,
		},[]byte{})

		sha256Hash := sha256.New()
		sha256Hash.Write(blockBytes)
		block256Hash = sha256Hash.Sum(nil)

		//sha256hash(区块 + nonce） 对应的大整数
		bigBlock =bigBlock.SetBytes(block256Hash)  //转换大整数
		if p.Target.Cmp(bigBlock) ==1 {  //满足条件时，退出循环
			fmt.Println("计算了",count,"次")
			break
		}
		nonce ++ //条件不满足 nonce 循环
	}
	return block256Hash,nonce
}