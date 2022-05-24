package dependency

import (
	"github.com/DuduNeves/Estoque_v1/core/product"
	"github.com/DuduNeves/Estoque_v1/database"
)

func Products() product.ProductsController {
	productsRepository := product.NewProductsRepository(database.GetDatabase())
	productsService := product.NewProductsService(productsRepository)
	productsController := product.NewProductsController(productsService)

	return productsController
}
