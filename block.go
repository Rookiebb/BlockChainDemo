package main

import "time"

type Block struct {
	Trans        []Transaction
	Time         time.Time
	PreviousHash string
	Hash         string
	nounce       int
}

func (block *Block) computeHash() string {
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
	this.Hash = this.computeHash()
	for {
		if this.Hash[0:difficulty] != getAnswer(difficulty) {
			this.nounce++
			this.Hash = this.computeHash()
		} else {
			break
		}
	}
	//fmt.Println("挖矿结束", this.Hash)
}
