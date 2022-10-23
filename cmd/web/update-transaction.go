package main

import (
	"log"
	"test-eth/data"
	"test-eth/internal/handlers"
	"time"
)

func listenForBlocks(r *handlers.Repository) {
	go func() {
		for {
			last, err := r.DB.LastBlockNumber()
			if err != nil {
				log.Println(err)
			}
			latest, err := data.GetLatestBlock()
			if err != nil {
				log.Println(err)
			}

			if last.Number == latest.Number && err != nil {
				continue
			} else {
				data.InsertBlockNum(latest.Number)
				err := data.PopulateBlockInformation(latest.Number)
				if err != nil {
					log.Println(err)
				}
			}
			time.Sleep(20 * time.Second)
		}
	}()

}
