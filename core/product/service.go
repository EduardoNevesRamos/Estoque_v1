package product

import (
	"github.com/DuduNeves/Estoque_v1/database/entity"
)

type IProductsService interface {
	GetAllProducts() ([]entity.Products, error)
	GetProduct(productId uint64) (*entity.Products, error)
	CreateProducts(product *entity.Products) error
	UpdateProducts(product *entity.Products) error
}

type ProductsService struct {
	repository IProductsRepository
}

func NewProductsService(
	repository IProductsRepository,
) IProductsService {
	return &ProductsService{
		repository: repository,
	}
}

func (s *ProductsService) GetProduct(productId uint64) (*entity.Products, error) {
	product, err := s.repository.GetProduct(productId)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductsService) GetAllProducts() ([]entity.Products, error) {
	product, err := s.repository.GetAllProducts()

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductsService) CreateProducts(product *entity.Products) error {
	return s.repository.CreateProducts(product)
}

func (s *ProductsService) UpdateProducts(product *entity.Products) error {
	return s.repository.UpdateProducts(product)
}
