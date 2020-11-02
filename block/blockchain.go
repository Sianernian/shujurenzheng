package blockchain

import (
	"fmt"
	"github.com/boltdb/bolt"
)

//bucket 的名字
var BUCKET_NAME ="class3"

// 最新区块的 key
var LAST_KEY = "lasthash"
//存储区块链的文件名
var BLOCKCHAIN_FILE_NAME ="BlockChainn.db"

/**
 * 区块链结构体实例定义:用于表示代表一条区块链
 * 该区块链包含以下功能:
 		① 将新产生的区块与已有的区块链接起来，并保存
		② 可以查询某个区块的信息
		③ 可以将所有区块进行遍历，输出区块信息
*/

type BlockChain struct{
	LastHash []byte
	BoltDb *bolt.DB
}

/**
 * 用于创建一条区块链，并返回该区块链实例
	解释：由于区块链就是由一个一个的区块组成的，因此,如果要创建一条区块链，那么必须要先
		创建一个区块，该区块作为该条区块链的创世区块。
*/

func NewBlockChain() BlockChain{
	// 1.打开存储区块数据的chain.db文件
	db,err :=bolt.Open(BLOCKCHAIN_FILE_NAME,0600,nil)
	if err !=nil{
		panic(err.Error())
	}
	var bl BlockChain
	//先从区块链中都看是否创世区块已经存在
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			bucket ,err =tx.CreateBucket([]byte(BUCKET_NAME))
			if err != nil {
				panic(err.Error())
			}
		}
		lasthash := bucket.Get([]byte(LAST_KEY))

		if len(lasthash) == 0 { // 如果没有创世区块
			//2. 创建一个创世区块
			gen := CreatGenesisBlock()
			fmt.Printf("创世区块hash:%x\n",gen.Hash)
			//3. 创建一个存储区块链的文件
			bl = BlockChain{
				LastHash: gen.Hash,
				BoltDb:   db,
			}
			genbytes, _ := gen.Serialize()
			 bucket.Put(gen.Hash, genbytes)
			 bucket.Put([]byte(LAST_KEY), gen.Hash)
		} else {
			//有创世区块
			lastHash := bucket.Get([]byte(LAST_KEY))
			lastBlockBytes := bucket.Get(lastHash)
			lastBlock, err := DeSerialize(lastBlockBytes)
			if err != nil {
				panic("数据有误")
			}
			bl = BlockChain{
				LastHash: lastBlock.Hash,
				BoltDb:   db,
			}

		}

		return nil
	})
	return bl
}

/**
 * 调用BlockChain的该SaveBlock方法，该方法可以将一个生成的新区块保存到chain.db文件中
 */

func (bc BlockChain)SaveBlock(data []byte)(Block ,error){

	var lastBlock *Block
	db :=bc.BoltDb

	db.View(func(tx *bolt.Tx) error {
	bucket:=tx.Bucket([]byte(BUCKET_NAME))

	 if bucket ==nil{
	 	panic("没有bucket")
	 }
	  LastblockByte:=bucket.Get(bc.LastHash)
	  lastBlock ,_=DeSerialize(LastblockByte)

		return nil
	})

	// 生产新的区块
	newBlock :=NewBlock(lastBlock.Height+1,data,lastBlock.Hash)
	//更新 .db文件 把 新的区块存入 .db中
	db.Update(func(tx *bolt.Tx) error {
		bucket :=tx.Bucket([]byte(BUCKET_NAME))

		//区块序列化
		newBlockBytes , _:=newBlock.Serialize()
		//信息保存到db中
		bucket.Put(newBlock.Hash,newBlockBytes)
		//跟新追后一个区块hash值的记录
		bucket.Put([]byte(LAST_KEY),newBlock.Hash)
		bc.LastHash = newBlock.Hash

		return nil
	})
	return newBlock ,nil
}




