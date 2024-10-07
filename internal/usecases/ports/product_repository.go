package ports

import "golang_with_couchdb2/internal/domain/entities"

type ProductRepository interface {
	GetAllProducts() ([]entities.Product, error)
    Create(product *entities.Product) error
    GetByID(id string) (*entities.Product, error)
    Update(product *entities.Product) error
    Delete(id string) error
}
