package services

import (
	"gin/dto"
	"gin/entities"
	"gin/repositories"
	"github.com/mashingan/smapping"
	"log"
)

type ProductService interface {
	Create(product dto.ProductDTO) entities.Product
	Update(product dto.ProductDTO) entities.Product
	Delete(productID string) error
	FindAll(qtyFilter string) []entities.Product
	IsDuplicateSkuID(skuID string) bool
}

type productService struct {
	productRepository repositories.ProductRepository
}

func NewProductService(productRepo repositories.ProductRepository) ProductService {
	return &productService{
		productRepository: productRepo,
	}
}

func (service *productService) Create(product dto.ProductDTO) entities.Product {
	productToCreate := entities.Product{}
	err := smapping.FillStruct(&productToCreate, smapping.MapFields(&product))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	return service.productRepository.Insert(productToCreate)
}

func (service *productService) Update(product dto.ProductDTO) entities.Product {
	productToUpdate := entities.Product{}
	err := smapping.FillStruct(&productToUpdate, smapping.MapFields(&product))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}

	return service.productRepository.Update(productToUpdate)
}

func (service *productService) Delete(productID string) error {
	return service.productRepository.Delete(productID)
}

func (service *productService) FindAll(qtyFilter string) []entities.Product {
	return service.productRepository.FindAll(qtyFilter)
}

func (service *productService) IsDuplicateSkuID(skuID string) bool {
	res := service.productRepository.IsDuplicateSkuID(skuID)
	return !(res.Error == nil)
}
