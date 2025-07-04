package usecase

import (
	"product_stock/model"
	"product_stock/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}
	product.ID = productId

	return product, nil
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) GetProductById(id_product int) (*model.Product, error) {

	product, err := pu.repository.GetProductById(id_product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (pu *ProductUsecase) UpdateProduct(product model.Product) (*model.Product, error) {

	updateProduct, err := pu.repository.UpdateProduct(product)
	if err != nil {
		return nil, err
	}
	return updateProduct, nil
}

func (pu *ProductUsecase) DeleteProduct(id int) error {
	return pu.repository.DeleteProduct(id)
}
