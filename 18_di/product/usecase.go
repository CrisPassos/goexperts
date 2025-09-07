package product

type ProductUseCase struct {
	repository ProductRepositoryInterface
}

func NewProdutUseCase(repository ProductRepositoryInterface) *ProductUseCase {
	return &ProductUseCase{repository}
}

// GetProduct return a product by its ID
// This Product was not supposed to be returned. We should return a DTO instead.
// However, for simplicity, we are returning the entity directly.
func (u *ProductUseCase) GetProduct(id int) (Product, error) {
	return u.repository.GetProduct(id)
}
