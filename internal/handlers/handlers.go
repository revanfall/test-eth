package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"strconv"
	"test-eth/helpers"
	"test-eth/internal/models"
	"test-eth/internal/repository"
	"test-eth/internal/repository/dbrepo"
)

var Repo *Repository

type Repository struct {
	DB     repository.DatabaseRepo
	Client *mongo.Client
}

func NewRepository(client *mongo.Client) *Repository {
	return &Repository{DB: dbrepo.NewMongoRepo(client)}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Transactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := m.DB.AllTransactions()
	if err != nil {
		log.Println(err)
		http.Error(w, "Error getting transactions", http.StatusInternalServerError)
	}

	var result []models.ResultTransactionJSON
	var blockData models.BlockInformationResultData
	for _, n := range data {
		blockData, _ = m.DB.BlockInfo(n.BlockNumber)
		cNum, _ := m.DB.GetConfirmationNumber(n.BlockNumber)
		result = append(result, models.ResultTransactionJSON{
			ID:                 n.Hash,
			From:               n.From,
			To:                 n.To,
			BlockNum:           n.BlockNumber,
			ConfirmationNumber: cNum,
			Timestamp:          blockData.Timestamp,
			Value:              helpers.CountValue(n.Value),
			Commission:         helpers.CountCommission(n.Gas, n.GasPrice),
		})
	}
	u, err := json.Marshal(result)
	w.Write(u)

}

func (m *Repository) TransactionsPaginationWithLimit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	page, err := strconv.Atoi(chi.URLParam(r, "pnum"))
	if err != nil {
		http.Error(w, "Error getting page", http.StatusBadRequest)
	}
	limit, err := strconv.Atoi(chi.URLParam(r, "lnum"))
	if err != nil {
		http.Error(w, "Error getting page", http.StatusBadRequest)
	}

	data, err := m.DB.AllTransactionsPagination(int64(page), int64(limit))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error getting transactions", http.StatusInternalServerError)
	}

	var result []models.ResultTransactionJSON
	var blockData models.BlockInformationResultData
	for _, n := range data {
		blockData, _ = m.DB.BlockInfo(n.BlockNumber)
		cNum, _ := m.DB.GetConfirmationNumber(n.BlockNumber)

		result = append(result, models.ResultTransactionJSON{
			ID:                 n.Hash,
			From:               n.From,
			To:                 n.To,
			BlockNum:           n.BlockNumber,
			ConfirmationNumber: cNum,
			Timestamp:          blockData.Timestamp,
			Value:              helpers.CountValue(n.Value),
			Commission:         helpers.CountCommission(n.Gas, n.GasPrice),
		})
	}
	u, err := json.Marshal(result)
	w.Write(u)
}

func (m *Repository) TransactionsPagination(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	page, err := strconv.Atoi(chi.URLParam(r, "pnum"))
	if err != nil {
		http.Error(w, "Error getting page", http.StatusBadRequest)
	}

	data, err := m.DB.AllTransactionsPagination(int64(page), 50)
	if err != nil {
		http.Error(w, "Error getting transactions", http.StatusInternalServerError)
	}

	var result []models.ResultTransactionJSON
	var blockData models.BlockInformationResultData
	for _, n := range data {
		blockData, _ = m.DB.BlockInfo(n.BlockNumber)
		cNum, _ := m.DB.GetConfirmationNumber(n.BlockNumber)

		result = append(result, models.ResultTransactionJSON{
			ID:                 n.Hash,
			From:               n.From,
			To:                 n.To,
			BlockNum:           n.BlockNumber,
			ConfirmationNumber: cNum,
			Timestamp:          blockData.Timestamp,
			Value:              helpers.CountValue(n.Value),
			Commission:         helpers.CountCommission(n.Gas, n.GasPrice),
		})
	}

	u, err := json.Marshal(result)
	w.Write(u)
}

func (m *Repository) TransactionByHash(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	hash := chi.URLParam(r, "hash")
	data, err := m.DB.TransactionByHash(hash)
	if err != nil {
		http.Error(w, "Error getting transaction by hash", http.StatusInternalServerError)
	}

	blockData, err := m.DB.BlockInfo(data.BlockNumber)
	if err != nil {
		http.Error(w, "Error getting transaction by hash", http.StatusInternalServerError)
	}
	cNum, _ := m.DB.GetConfirmationNumber(data.BlockNumber)

	res := models.ResultTransactionJSON{
		ID:                 data.Hash,
		From:               data.From,
		To:                 data.To,
		BlockNum:           data.BlockNumber,
		ConfirmationNumber: cNum,
		Timestamp:          blockData.Timestamp,
		Value:              helpers.CountValue(data.Value),
		Commission:         helpers.CountCommission(data.Gas, data.GasPrice),
	}

	u, err := json.Marshal(res)
	w.Write(u)

}

func (m *Repository) TransactionBySender(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	s := chi.URLParam(r, "s")
	data, err := m.DB.TransactionBySender(s)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error getting transactions", http.StatusInternalServerError)
	}

	var result []models.ResultTransactionJSON
	var blockData models.BlockInformationResultData
	for _, n := range data {
		blockData, _ = m.DB.BlockInfo(n.BlockNumber)
		cNum, _ := m.DB.GetConfirmationNumber(n.BlockNumber)
		result = append(result, models.ResultTransactionJSON{
			ID:                 n.Hash,
			From:               n.From,
			To:                 n.To,
			BlockNum:           n.BlockNumber,
			ConfirmationNumber: cNum,
			Timestamp:          blockData.Timestamp,
			Value:              helpers.CountValue(n.Value),
			Commission:         helpers.CountCommission(n.Gas, n.GasPrice),
		})
	}
	u, err := json.Marshal(result)
	w.Write(u)

}

func (m *Repository) TransactionByReceiver(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rec := chi.URLParam(r, "r")
	data, err := m.DB.TransactionByReceiver(rec)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error getting transactions", http.StatusInternalServerError)
	}

	var result []models.ResultTransactionJSON
	var blockData models.BlockInformationResultData
	for _, n := range data {
		blockData, _ = m.DB.BlockInfo(n.BlockNumber)
		cNum, _ := m.DB.GetConfirmationNumber(n.BlockNumber)
		result = append(result, models.ResultTransactionJSON{
			ID:                 n.Hash,
			From:               n.From,
			To:                 n.To,
			BlockNum:           n.BlockNumber,
			ConfirmationNumber: cNum,
			Timestamp:          blockData.Timestamp,
			Value:              helpers.CountValue(n.Value),
			Commission:         helpers.CountCommission(n.Gas, n.GasPrice),
		})
	}
	u, err := json.Marshal(result)
	w.Write(u)
}

func (m *Repository) TransactionsByTimeStamp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ts := chi.URLParam(r, "ts")
	data, err := m.DB.TransactionsByTimestamp(ts)
	if err != nil {
		http.Error(w, "Error getting transactions by timestamp", http.StatusInternalServerError)
	}

	var result []models.ResultTransactionJSON
	for _, n := range data {
		cNum, _ := m.DB.GetConfirmationNumber(n.BlockNumber)

		result = append(result, models.ResultTransactionJSON{
			ID:                 n.Hash,
			From:               n.From,
			To:                 n.To,
			BlockNum:           n.BlockNumber,
			ConfirmationNumber: cNum,
			Timestamp:          ts,
			Value:              helpers.CountValue(n.Value),
			Commission:         helpers.CountCommission(n.Gas, n.GasPrice),
		})
	}
	u, err := json.Marshal(result)
	w.Write(u)
}
