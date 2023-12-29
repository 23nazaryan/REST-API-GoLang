package services

import (
	"gin/dto"
	"gin/entities"
	"gin/repositories"
	"github.com/mashingan/smapping"
	"log"
)

type StockService interface {
	Create(sock dto.StockDTO) entities.Stock
	Update(stock dto.StockDTO) entities.Stock
	Delete(stockID string) error
	FindAll(stockType string) []entities.Stock
}

type stockService struct {
	stockRepository repositories.StockRepository
}

func NewStockService(stockRepo repositories.StockRepository) StockService {
	return &stockService{
		stockRepository: stockRepo,
	}
}

func (service *stockService) Create(stock dto.StockDTO) entities.Stock {
	stockToCreate := entities.Stock{}
	err := smapping.FillStruct(&stockToCreate, smapping.MapFields(&stock))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	return service.stockRepository.Insert(stockToCreate)
}

func (service *stockService) Update(stock dto.StockDTO) entities.Stock {
	stockToUpdate := entities.Stock{}
	err := smapping.FillStruct(&stockToUpdate, smapping.MapFields(&stock))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}

	return service.stockRepository.Update(stockToUpdate)
}

func (service *stockService) Delete(stockID string) error {
	return service.stockRepository.Delete(stockID)
}

func (service *stockService) FindAll(stockType string) []entities.Stock {
	return service.stockRepository.FindAll(stockType)
}
