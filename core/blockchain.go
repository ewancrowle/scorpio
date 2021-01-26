package core

import (
	"scorpio/core/types"
)

var Bc Blockchain

type Blockchain struct {
	Chain  []*types.Block
	TxPool types.Transactions
}

func (bc Blockchain) Update(bs []*types.Block) {
	if len(bs) > len(bc.Chain) {
		bc.Chain = bs
	}
}

func (bc Blockchain) TxAdd(tx *types.Transaction) int {
	bc.TxPool = append(bc.TxPool, tx)
	return bc.LastBlock().Header.Index + 1
}

func (bc Blockchain) LastBlock() *types.Block {
	return bc.Chain[len(bc.Chain)-1]
}

func (bc Blockchain) NewBlock() *types.Block {
	parent := bc.LastBlock()
	b := types.NewBlock(parent.Header, bc.TxPool)
	bc.Chain = append(bc.Chain, b)
	bc.TxPool = types.Transactions{}
	return b
}
