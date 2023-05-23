package main

import (
	"fmt"
	"time"
)

type Chain struct {
	mychain     []Block
	mineReward  float32
	difficulty  int
	transacPool []Transaction
}

func NewChain() *Chain {
	return &Chain{[]Block{BigBang()}, 50.0, 2, []Transaction{}}
}

func BigBang() Block {
	GenesisBlock := Block{Trans: []Transaction{{From: "god"}}, Time: time.Now()}
	GenesisBlock.Hash = GenesisBlock.computeHash()
	return GenesisBlock
}

func (this *Chain) GetLatestBlock() Block {
	return this.mychain[len(this.mychain)-1]
}

func (this *Chain) addtransaction(transaction Transaction) {
	this.transacPool = append(this.transacPool, transaction)
}

func (this *Chain) mineTransactionPool(mineRewardAdress string) {
	mineRewardTransaction := Transaction{"", mineRewardAdress, this.mineReward}
	this.transacPool = append(this.transacPool, mineRewardTransaction)
	//挖矿
	newBlock := Block{
		Trans:        this.transacPool,
		PreviousHash: this.GetLatestBlock().Hash,
		Time:         time.Now(),
	}
	newBlock.mine(this.difficulty)
	//添加到区块链，清空交易池
	this.mychain = append(this.mychain, newBlock)
	this.transacPool = []Transaction{}
}

func (this *Chain) ValidateChain() bool {
	if len(this.mychain) == 1 {
		if this.mychain[0].Hash != this.mychain[0].computeHash() {
			return false
		}
		return true
	}
	for i := 1; i < len(this.mychain); i++ {
		blocktovalidate := this.mychain[i]
		//保证当前区块没有被篡改
		if blocktovalidate.Hash != blocktovalidate.computeHash() {
			fmt.Println("数据被篡改")
			return false
		}
		//如果有人同时修改了数据和哈希值,那么区块的哈希值就和下一个区块保存的哈希对不上
		preblock := this.mychain[i-1]
		if preblock.Hash != blocktovalidate.PreviousHash {
			fmt.Println("区块断裂")
			return false
		}
	}
	return true
}

func (this *Chain) toString() {
	for i := 0; i < len(this.mychain); i++ {
		fmt.Printf("第 %v 块：\n创建时间：%v\n前块哈希: %v\n当前块哈希: %v\n交易：%v\n交易池：%v\n",
			i, this.mychain[i].Time, this.mychain[i].PreviousHash, this.mychain[i].Hash, this.mychain[i].Trans, this.transacPool)
	}
	this.ValidateChain()
}
