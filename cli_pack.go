package main

import (
	"fmt"
	"log"
)

func (cli *CLI) pack(addr string) {
	if !ValidateAddress(addr) {
		log.Panic("ERROR: 发送地址非法")
	}

	bc := NewBlockchain() //打开数据库，读取区块链并构建区块链实例
	defer bc.Db.Close()   //转账完毕，关闭数据库
	//txs := conn_recv(addr)
	txs := recv_tx(addr)
	bc.MineBlock(txs)
	fmt.Println("成功打包交易")
}
