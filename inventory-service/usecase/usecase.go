package usecase

import (
	"github.com/Iman-ws/ecommerce-pharmacy/inventory-service/model"
	"github.com/Iman-ws/ecommerce-pharmacy/inventory-service/repo"
)

type UseCase interface {
	AddProduct(p *model.Product) error
	GetProduct(id int) (*model.Product, error)
	UpdateProduct(p *model.Product) error
	DeleteProduct(id int) error
	ListProducts() ([]model.Product, error)
}

type ProductUseCase struct {
	repo repo.Repository
}

func NewProductUseCase(repo repo.Repository) *ProductUseCase {
	return &ProductUseCase{repo: repo}
}

func (uc *ProductUseCase) AddProduct(p *model.Product) error {
	return uc.repo.CreateProduct(p)
}

func (uc *ProductUseCase) GetProduct(id int) (*model.Product, error) {
	return uc.repo.GetProductByID(id)
}

func (uc *ProductUseCase) UpdateProduct(p *model.Product) error {
	return uc.repo.UpdateProduct(p)
}

func (uc *ProductUseCase) DeleteProduct(id int) error {
	return uc.repo.DeleteProduct(id)
}

func (uc *ProductUseCase) ListProducts() ([]model.Product, error) {
	return uc.repo.ListProducts()
}
