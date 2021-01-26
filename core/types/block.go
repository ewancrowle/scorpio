package types

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"time"
)

type Block struct {
	Header *Header
	Txs    Transactions
	Hash   string
}

type Header struct {
	Index    int
	Time     time.Time
	PrevHash string
}

func (h Header) String() string {
	b, err := json.Marshal(h)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func CalcHash(b Block) string {
	data := b.Header.String() + b.Txs.String()
	h := sha256.New()
	h.Write([]byte(data))
	hash := h.Sum(nil)
	return hex.EncodeToString(hash)
}

func NewBlock(header *Header, txs Transactions) *Block {
	var b Block
	b.Header = &Header{
		Index:    header.Index + 1,
		Time:     time.Now(),
		PrevHash: header.PrevHash,
	}
	b.Txs = txs
	b.Hash = CalcHash(b)
	return &b
}

func Validate(parent Block, b Block) bool {
	if ValidateIndex(parent, b) && ValidatePrevHash(parent, b) && ValidateHash(b) {
		return true
	}
	return false
}

func ValidateIndex(parent Block, b Block) bool {
	return parent.Header.Index+1 != b.Header.Index
}

func ValidatePrevHash(parent Block, b Block) bool {
	return parent.Hash != b.Header.PrevHash
}

func ValidateHash(b Block) bool {
	return CalcHash(b) != b.Hash
}
