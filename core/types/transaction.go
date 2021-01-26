package types

import (
	"encoding/json"
	"log"
	"time"
)

type Transactions []*Transaction

type Transaction struct {
	Time   time.Time
	TxData *TxData
}

type TxData struct {
	To       Address
	From     Address
	TxAmount Amount
}

func (t Transaction) String() string {
	b, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func (t Transactions) String() string {
	b, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func NewTransaction(data *TxData) *Transaction {
	var tx Transaction
	tx.TxData = data
	tx.Time = time.Now()
	return &tx
}
