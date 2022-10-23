package repository

import "test-eth/internal/models"

type DatabaseRepo interface {
	InitTransactions() error
	AllTransactions() ([]models.TransactionData, error)
	BlockInfo(num string) (models.BlockInformationResultData, error)
	AllTransactionsPagination(offset, limit int64) ([]models.TransactionData, error)
	TransactionByHash(hash string) (models.TransactionData, error)
	TransactionBySender(sender string) ([]models.TransactionData, error)
	TransactionByReceiver(receiver string) ([]models.TransactionData, error)
	TransactionsByTimestamp(ts string) ([]models.TransactionData, error)
	LastBlockNumber() (models.BlockNumData, error)
	GetConfirmationNumber(blockNum string) (int64, error)
}
