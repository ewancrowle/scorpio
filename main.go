package main

import (
	"log"
	"scorpio/core"
	"scorpio/core/types"
	"scorpio/server"
	"time"
)

var (
	Server server.Server
	Config = core.DefaultConfig
)

func main() {
	go func() {
		head := &types.Header{
			Index:    -1,
			Time:     time.Now(),
			PrevHash: "0",
		}
		b := types.NewBlock(head, nil)
		core.Bc.Chain = append(core.Bc.Chain, b)
		core.Bc.TxPool = []*types.Transaction{}
	}()

	log.Fatal(Server.Run(Config.Port))
}
