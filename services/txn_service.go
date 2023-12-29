package services

import (
	"gin/dto"
	"gin/entities"
	"gin/repositories"
	"github.com/mashingan/smapping"
	"log"
)

type TxnService interface {
	Create(txn dto.TransactionDTO) entities.Transaction
	Update(txn dto.TransactionDTO) entities.Transaction
	Delete(txnID string) error
	List(section string, id string) []entities.Transaction
}

type txnService struct {
	txnRepository repositories.TxnRepository
}

func NewTxnService(txnRepo repositories.TxnRepository) TxnService {
	return &txnService{
		txnRepository: txnRepo,
	}
}

func (service *txnService) Create(txn dto.TransactionDTO) entities.Transaction {
	txnToCreate := entities.Transaction{}
	err := smapping.FillStruct(&txnToCreate, smapping.MapFields(&txn))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	return service.txnRepository.Insert(txnToCreate)
}

func (service *txnService) Update(txn dto.TransactionDTO) entities.Transaction {
	txnToUpdate := entities.Transaction{}
	err := smapping.FillStruct(&txnToUpdate, smapping.MapFields(&txn))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}

	return service.txnRepository.Update(txnToUpdate)
}

func (service *txnService) Delete(txnID string) error {
	return service.txnRepository.Delete(txnID)
}

func (service *txnService) List(section string, id string) []entities.Transaction {
	return service.txnRepository.List(section, id)
}
