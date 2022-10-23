package data

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gookit/config/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"log"
	"net/http"
	"strconv"
	"test-eth/helpers"
	"test-eth/internal/models"
	"time"
)

var collBlockNums *mongo.Collection
var collBlockInfo *mongo.Collection
var collTransInfo *mongo.Collection

func InitData(client *mongo.Client) {
	collBlockNums = client.Database("eth").Collection("blockNums")
	collBlockInfo = client.Database("eth").Collection("blockInfo")
	collTransInfo = client.Database("eth").Collection("transactionInfo")
	docNum, err := collBlockNums.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		log.Println(err)
	}
	if docNum == 0 {
		_ = populateDataBlockNums()
		cursor, _ := collBlockNums.Find(context.TODO(), bson.D{})

		var blockNums []models.BlockNumData
		for cursor.Next(context.Background()) {

			m := models.BlockNumData{}
			err := cursor.Decode(&m)
			if err != nil {

			}
			blockNums = append(blockNums, m)
		}
		for _, m := range blockNums {
			PopulateBlockInformation(m.Number)
		}
	} else {
		log.Println("popa")
	}

}

func initPreviousBlocks(latest int64) []interface{} {
	var latestBlocks []interface{}
	for i := 500; i > 0; i-- {
		latestBlocks = append(latestBlocks, bson.D{{"blockNum", strconv.FormatInt(latest-int64(i), 16)}})
	}
	return latestBlocks
}

func InsertBlockNum(latest string) {
	_, err := collBlockNums.InsertOne(context.TODO(), bson.D{{"blockNum", latest}})
	if err != nil {
		log.Println(err)
	}
}

func populateDataBlockNums() error {
	resp, err := http.Get(fmt.Sprintf("https://api.etherscan.io/api?module=proxy&action=eth_blockNumber&apikey=%s", config.String("API_KEY")))
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var latestBlock models.BlockNumJSON
	err = json.Unmarshal(body, &latestBlock)

	decLatestBlock := helpers.ParseHexToDec(latestBlock.Number)
	if err != nil {
		return err
	}
	fmt.Println("latest", decLatestBlock)
	_, err = collBlockNums.InsertMany(context.TODO(), initPreviousBlocks(decLatestBlock))
	if err != nil {
		return err
	}
	return nil
}

func PopulateBlockInformation(blockNum string) error {
	time.Sleep(500 * time.Millisecond)
	resp, err := http.Get(fmt.Sprintf("https://api.etherscan.io/api?module=proxy&action=eth_getBlockByNumber&tag=%s&boolean=true&apikey=%s", blockNum, config.String("API_KEY")))
	fmt.Println(blockNum)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var res models.BlockInformation
	err = json.Unmarshal(body, &res)
	_, _ = collBlockInfo.InsertOne(context.TODO(), bson.D{{"difficulty", res.Result.Difficulty},
		{"gasUsed", res.Result.GasUsed},
		{"hash", res.Result.Hash},
		{"number", res.Result.Number},
		{"timestamp", res.Result.Timestamp},
	})
	for _, n := range res.Result.Transactions {
		_ = populateTransactionInformation(n.Hash)
	}

	return nil
}

func populateTransactionInformation(hash string) error {
	time.Sleep(500 * time.Millisecond)
	resp, err := http.Get(fmt.Sprintf("https://api.etherscan.io/api?module=proxy&action=eth_getTransactionByHash&txhash=%s&apikey=%s", hash, config.String("API_KEY")))
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var res models.TransactionInfoJSON
	err = json.Unmarshal(body, &res)

	_, _ = collTransInfo.InsertOne(context.TODO(), bson.D{{"from", res.Result.From},
		{"gas", res.Result.Gas},
		{"gasPrice", res.Result.GasPrice},
		{"hash", res.Result.Hash},
		{"blockNumber", res.Result.BlockNumber},
		{"to", res.Result.To},
		{"value", res.Result.Value}})
	return nil
}

func GetLatestBlock() (models.BlockNumJSON, error) {
	var latestBlock models.BlockNumJSON
	resp, err := http.Get(fmt.Sprintf("https://api.etherscan.io/api?module=proxy&action=eth_blockNumber&apikey=%s", config.String("API_KEY")))
	time.Sleep(500 * time.Millisecond)
	if err != nil {
		return latestBlock, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return latestBlock, err
	}

	err = json.Unmarshal(body, &latestBlock)

	return latestBlock, nil
}
