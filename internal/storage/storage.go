package storage

import (
	"productStorage/internal/models"
)

type ProductStorage interface{
	Add(product *models.Product)
	GetAll()	[]*models.Product
}

type Storage struct{
	products []*models.Product
}

func NewStorage() ProductStorage {
	return &Storage{
		products: make([]*models.Product, 0)
	}
}

func (storage *Storage) Add (product *models.Product) {
	storage.products = append(storage.products, product)
}

func (storage *Storage) GetAll() []*models.Product