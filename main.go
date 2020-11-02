package main

import (
	blockchain "DataCertProject/block"
	"DataCertProject/db_mysql"
	_ "DataCertProject/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {

	bc :=blockchain.NewBlockChain()
	fmt.Printf("区块Hash值：%x\n",bc.LastHash)
	block,err :=bc.SaveBlock([]byte("存储上链信息"))
	if err !=nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("区块的高度：%d\n",block.Height)
	fmt.Printf("区块的PrevHash:%x\n", block.PrevHash)


	return
	// 1.链接数据库
	db_mysql.ConnectDB()

	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/img","./static/img")
	beego.SetStaticPath("/css","./static/css")

	beego.Run()
}

