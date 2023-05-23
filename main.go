package main

func main() {
	s := NewChain()
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
