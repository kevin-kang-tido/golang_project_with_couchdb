package interactors

import (
    "golang_with_couchdb2/internal/domain/entities"
    "golang_with_couchdb2/internal/usecases/ports"
)

type ProductInteractor struct {
    ProductRepo ports.ProductRepository
}

func NewProductInteractor(repo ports.ProductRepository) *ProductInteractor {
    return &ProductInteractor{
        ProductRepo: repo,
    }
}

// Get all products
func (interactor *ProductInteractor) GetAllProducts() ([]entities.Product, error) {
    return interactor.ProductRepo.GetAllProducts()
}

func (i *ProductInteractor) CreateProduct(product *entities.Product) error {
    return i.ProductRepo.Create(product)
}

func (i *ProductInteractor) GetProductByID(id string) (*entities.Product, error) {
    return i.ProductRepo.GetByID(id)
}

func (i *ProductInteractor) UpdateProduct(product *entities.Product) error {
    return i.ProductRepo.Update(product)
}

func (i *ProductInteractor) DeleteProduct(id string) error {
    return i.ProductRepo.Delete(id)
}
