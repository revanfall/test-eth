package dbrepo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	options2 "go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"test-eth/internal/models"
)

func (m *mongoDBRepo) InitTransactions() error {
	return nil
}

func (m *mongoDBRepo) AllTransactions() ([]models.TransactionData, error) {
	cursor, err := m.DB.Database("eth").Collection("transactionInfo").Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	var data []models.TransactionData
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (m *mongoDBRepo) BlockInfo(blockNum string) (models.BlockInformationResultData, error) {
	var data models.BlockInformationResultData

	err := m.DB.Database("eth").Collection("blockInfo").FindOne(context.TODO(),
		bson.D{{"number", blockNum}}).Decode(&data)
	if err != nil {
		log.Println(err)
		return data, err
	}

	return data, nil
}

func (m *mongoDBRepo) AllTransactionsPagination(offset, limit int64) ([]models.TransactionData, error) {
	options := options2.Find()
	options.SetLimit(limit)
	options.SetSkip(offset * limit)
	cursor, err := m.DB.Database("eth").Collection("transactionInfo").Find(context.TODO(), bson.D{}, options)
	if err != nil {
		return nil, err
	}

	var data []models.TransactionData
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (m *mongoDBRepo) TransactionByHash(hash string) (models.TransactionData, error) {
	var data models.TransactionData

	err := m.DB.Database("eth").Collection("transactionInfo").FindOne(context.TODO(),
		bson.D{{"hash", hash}}).Decode(&data)
	if err != nil {
		log.Println(err)
		return data, err
	}
	return data, nil

}

func (m *mongoDBRepo) TransactionBySender(sender string) ([]models.TransactionData, error) {
	cursor, err := m.DB.Database("eth").Collection("transactionInfo").Find(context.TODO(), bson.D{{"from", sender}})
	if err != nil {
		return nil, err
	}
	var data []models.TransactionData
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (m *mongoDBRepo) TransactionByReceiver(receiver string) ([]models.TransactionData, error) {
	cursor, err := m.DB.Database("eth").Collection("transactionInfo").Find(context.TODO(), bson.D{{"to", receiver}})
	if err != nil {
		return nil, err
	}
	var data []models.TransactionData
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (m *mongoDBRepo) TransactionsByTimestamp(ts string) ([]models.TransactionData, error) {
	var blockData models.BlockInformationResultData

	err := m.DB.Database("eth").Collection("blockInfo").FindOne(context.TODO(),
		bson.D{{"timestamp", ts}}).Decode(&blockData)
	if err != nil {
		return nil, err
	}

	cursor, err := m.DB.Database("eth").Collection("transactionInfo").Find(context.TODO(), bson.D{{"blockNumber", blockData.Number}})
	if err != nil {
		return nil, err
	}
	var data []models.TransactionData
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (m *mongoDBRepo) LastBlockNumber() (models.BlockNumData, error) {
	var blockNum models.BlockNumData
	opts := options2.FindOne().SetSort(bson.M{"$natural": -1})
	err := m.DB.Database("eth").Collection("blockNums").FindOne(context.TODO(), bson.D{}, opts).Decode(&blockNum)
	if err != nil {
		return blockNum, err
	}
	return blockNum, nil
}

func (m *mongoDBRepo) GetConfirmationNumber(blockNum string) (int64, error) {
	block, err := m.BlockInfo(blockNum)
	if err != nil {
		return 0, err
	}
	filter := bson.D{{"_id", bson.D{{"$gt", block.ID}}}}
	blockCount, err := m.DB.Database("eth").Collection("blockNums").CountDocuments(context.TODO(), filter)
	return blockCount, nil
}
