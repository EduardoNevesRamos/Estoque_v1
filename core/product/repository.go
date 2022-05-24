package product

import (
	"fmt"

	"github.com/DuduNeves/Estoque_v1/database/entity"
	"gorm.io/gorm"
)

type IProductsRepository interface {
	GetAllProducts() ([]entity.Products, error)
	GetProduct(productId uint64) (*entity.Products, error)
	CreateProducts(product *entity.Products) error
	UpdateProducts(product *entity.Products) error
}

type ProductsRepository struct {
	db *gorm.DB
}

func NewProductsRepository(db *gorm.DB) IProductsRepository {
	return &ProductsRepository{
		db: db,
	}
}

func (r *ProductsRepository) GetAllProducts() ([]entity.Products, error) {
	var Products []entity.Products

	err := r.db.Find(&Products).Error
	if err != nil {
		return nil, fmt.Errorf("cannot find products: %v", err)
	}
	return Products, err
}

func (r *ProductsRepository) GetProduct(productId uint64) (*entity.Products, error) {
	var product entity.Products

	err := r.db.First(&product, productId).Error
	if err != nil {
		return &entity.Products{}, fmt.Errorf("cannot find products by id: %v", err)
	}

	return &product, err
}

func (r *ProductsRepository) CreateProducts(product *entity.Products) error {
	return r.db.Create(&product).
		Error
}

func (r *ProductsRepository) UpdateProducts(product *entity.Products) error {
	return r.db.Where(&product.ID).Save(&product).Error
}
