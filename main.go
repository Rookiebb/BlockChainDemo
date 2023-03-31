package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

func Sha256(src string) string {
	m := sha256.New()
	m.Write([]byte(src))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

type Transaction struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float32 `json:"amount"`
}

type Block struct {
	Trans        []Transaction
	Time         time.Time
	PreviousHash string
	Hash         string
	nounce       int
}

func toJSONString(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}
	return string(b)
}

func (block *Block) ComputeHash() string {
	s := block.PreviousHash + toJSONString(block.Trans) + toJSONString(block.Time) + toJSONString(block.nounce)
	return Sha256(s)
}

// 挖矿方法
func getAnswer(difficulty int) string {
	answer := ""
	for i := 0; i < difficulty; i++ {
		answer += "0"
	}
	return answer
}

func (this *Block) mine(difficulty int) {
	this.Hash = this.ComputeHash()
	for {
		if this.Hash[0:difficulty] != getAnswer(difficulty) {
			this.nounce++
			this.Hash = this.ComputeHash()
		} else {
			break
		}
	}
	//fmt.Println("挖矿结束", this.Hash)
}

type Chain struct {
	mychain     []Block
	mineReward  float32
	difficulty  int
	transacPool []Transaction
}

func newChain() *Chain {
	return &Chain{[]Block{BigBang()}, 50.0, 2, []Transaction{}}
}

func BigBang() Block {
	GenesisBlock := Block{Trans: []Transaction{{From: "god"}}, Time: time.Now()}
	GenesisBlock.Hash = GenesisBlock.ComputeHash()
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
		if this.mychain[0].Hash != this.mychain[0].ComputeHash() {
			return false
		}
		return true
	}
	for i := 1; i < len(this.mychain); i++ {
		blocktovalidate := this.mychain[i]
		//保证当前区块没有被篡改
		if blocktovalidate.Hash != blocktovalidate.ComputeHash() {
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

func main() {
	s := newChain()
	t1 := Transaction{"addr1", "addr2", 10.0}
	t2 := Transaction{"addr2", "addr3", 6.0}
	s.addtransaction(t1)
	s.addtransaction(t2)
	s.mineTransactionPool("addr4")
	t3 := Transaction{"addr4", "addr5", 16.0}
	t4 := Transaction{"addr2", "addr3", 8.0}
	s.addtransaction(t3)
	s.addtransaction(t4)
	s.mineTransactionPool("addr5")

	//s.mychain[1].Trans[0].Amount = 100
	s.toString()
}
