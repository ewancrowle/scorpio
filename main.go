package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"scorpio/core"
	"scorpio/core/types"
	"scorpio/server"
	"time"
)

var (
	Server server.Server
	Config core.Config
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	Config = core.Config{
		Address: os.Getenv("ADDR"),
	}

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

	log.Fatal(Server.Run(Config.Address))
}
