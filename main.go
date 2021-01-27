package main

import (
	"github.com/joho/godotenv"
	"log"
	"scorpio/config"
	"scorpio/core"
	"scorpio/core/types"
	"scorpio/server"
	"time"
)

var Server server.Server

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	config.Configure()

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

	log.Fatal(Server.Run(config.Address))
}
