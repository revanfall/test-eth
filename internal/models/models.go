package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlockNumJSON struct {
	Number string `json:"result"`
}

type BlockNumData struct {
	ID     *primitive.ObjectID `bson:"_id"`
	Number string              `bson:"blockNum"`
}

type BlockInformation struct {
	Result BlockInformationResultJSON `json:"result"`
}

type BlockInformationResultJSON struct {
	Difficulty   string            `json:"difficulty"`
	GasUsed      string            `json:"gasUsed"`
	Hash         string            `json:"hash"`
	Number       string            `json:"number"`
	Timestamp    string            `json:"timestamp"`
	Transactions []TransactionJSON `json:"transactions"`
}

type TransactionJSON struct {
	BlockNumber string `json:"blockNumber"`
	From        string `json:"from"`
	Gas         string `json:"gas"`
	GasPrice    string `json:"gasPrice"`
	Hash        string `json:"hash"`
	To          string `json:"to"`
}

type BlockInformationResultData struct {
	ID         *primitive.ObjectID `bson:"_id"`
	Difficulty string              `bson:"difficulty"`
	GasUsed    string              `bson:"gasUsed"`
	Hash       string              `bson:"hash"`
	Number     string              `bson:"number"`
	Timestamp  string              `bson:"timestamp"`
}

type TransactionData struct {
	BlockNumber string `bson:"blockNumber"`
	From        string `bson:"from"`
	Gas         string `bson:"gas"`
	GasPrice    string `bson:"gasPrice"`
	Hash        string `bson:"hash"`
	To          string `bson:"to"`
	Value       string `bson:"value"`
}

type TransactionInfoJSON struct {
	Result TransactionResultJSON `json:"result"`
}

type TransactionResultJSON struct {
	BlockNumber string `json:"blockNumber"`
	From        string `json:"from"`
	Gas         string `json:"gas"`
	GasPrice    string `json:"gasPrice"`
	Hash        string `json:"hash"`
	To          string `json:"to"`
	Value       string `json:"value"`
	Timestamp   string `json:"timestamp"`
}

type ResultTransactionJSON struct {
	ID                 string  `json:"id"`
	From               string  `json:"from"`
	To                 string  `json:"to"`
	BlockNum           string  `json:"blockNum"`
	ConfirmationNumber int64   `json:"confirmationNumber"`
	Timestamp          string  `json:"timestamp"`
	Value              float64 `json:"value"'`
	Commission         float64 `json:"commission"`
}
