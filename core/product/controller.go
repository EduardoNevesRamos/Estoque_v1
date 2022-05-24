package product

import (
	"net/http"
	"strconv"

	"github.com/DuduNeves/Estoque_v1/database"
	"github.com/DuduNeves/Estoque_v1/database/entity"
	"github.com/gin-gonic/gin"
)

type ProductsController struct {
	service IProductsService
}

func NewProductsController(service IProductsService) ProductsController {
	return ProductsController{
		service: service,
	}
}

func (c *ProductsController) GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	parseId, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(400, gin.H{
			"Error:": "ID has to be integer",
		})
		return
	}

	customer, err := c.service.GetProduct(uint64(parseId))

	if err != nil {
		ctx.JSON(400, gin.H{
			"Error:": "Can't find a product" + err.Error(),
		})
		return
	}

	ctx.JSON(200, customer)
}

func (c *ProductsController) GetAllProducts(ctx *gin.Context) {

	productsEntity, err := c.service.GetAllProducts()

	if err != nil {
		ctx.JSON(400, gin.H{
			"Error:": "Can't find products: " + err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, productsEntity)
}

func (c *ProductsController) CreateProducts(ctx *gin.Context) {
	newProducts := entity.Products{}
	if err := ctx.ShouldBindJSON(&newProducts); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := c.service.CreateProducts(&newProducts)
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error:": "Can't create product" + err.Error(),
		})

		return
	}

	ctx.JSON(200, newProducts)
}

func (c *ProductsController) UpdateProducts(ctx *gin.Context) {
	req := &entity.Products{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error:": "Can't bind JSON" + err.Error(),
		})
		return
	}

	err = c.service.UpdateProducts(req)
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error:": "Can't update product" + err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, req)
}

func DeleteProducts(ctx *gin.Context) {
	id := ctx.Param("id")
	parseId, err := strconv.Atoi(id)

	product := &entity.Products{}

	if err != nil {
		ctx.JSON(400, gin.H{
			"Error:": "Can't bind JSON" + err.Error(),
		})
	}

	db := database.GetDatabase()

	err = db.Delete(product, parseId).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error:": "Can't update product" + err.Error(),
		})

		return
	}

	ctx.JSON(200, product)
}
