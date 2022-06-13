package product

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/DuduNeves/Estoque_v1/database/entity"
	"github.com/DuduNeves/Estoque_v1/util"
	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type ServiceSpy struct {
	IProductsService
	GetAllErr   error
	GetAllResp  []entity.Products
	GetByIdErr  error
	GetByIdResp *entity.Products
	CreateErr   error
	UpdateErr   error
	DeleteErr   error
}

func (s *ServiceSpy) GetAllProducts() ([]entity.Products, error) {
	return s.GetAllResp, s.GetAllErr
}

func (s *ServiceSpy) GetProduct(productId uint64) (*entity.Products, error) {
	return s.GetByIdResp, s.GetByIdErr
}

func (s *ServiceSpy) CreateProducts(product *entity.Products) error {
	return s.CreateErr
}

func (s *ServiceSpy) UpdateProducts(product *entity.Products) error {
	return s.UpdateErr
}

func (s *ServiceSpy) DeleteProducts(id *int, product *entity.Products) error {
	return s.DeleteErr
}

func TestGetAllController(t *testing.T) {
	t.Run("Success - Should return no error", func(t *testing.T) {
		w, ctx := util.MockGin()

		_, err := http.NewRequest("GET", "/", nil)
		assert.Nil(t, err)

		serviceSpy := &ServiceSpy{}
		controller := NewProductsController(serviceSpy)

		controller.GetAllProducts(ctx)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})
	t.Run("Error - Should return error", func(t *testing.T) {
		w, ctx := util.MockGin()

		_, err := http.NewRequest("GET", "/", nil)
		assert.Nil(t, err)

		serviceSpy := &ServiceSpy{
			GetAllErr: gorm.ErrRecordNotFound,
		}
		controller := NewProductsController(serviceSpy)

		controller.GetAllProducts(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})
}

func TestGetByIdController(t *testing.T) {
	t.Run("Success - Should return no error", func(t *testing.T) {
		controller := NewProductsController(&ServiceSpy{
			GetByIdErr: nil,
		})

		w, ctx := util.MockGin()

		_, err := http.NewRequest("GET", "/", nil)
		assert.Nil(t, err)

		ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: "1"})

		controller.GetProduct(ctx)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})
	t.Run("Error - Should return BadRequest", func(t *testing.T) {
		w, ctx := util.MockGin()

		_, err := http.NewRequest("GET", "/", nil)
		assert.Nil(t, err)

		controller := NewProductsController(&ServiceSpy{})

		ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: "test"})

		controller.GetProduct(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("Error - Should return Internal Server Error", func(t *testing.T) {
		controller := NewProductsController(&ServiceSpy{
			GetByIdErr: errors.New("Server Internal Server Error"),
		})

		w, ctx := util.MockGin()

		ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: "1"})

		_, err := http.NewRequest("GET", "/", nil)
		assert.Nil(t, err)

		controller.GetProduct(ctx)

		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}

func TestCreateProducts(t *testing.T) {
	t.Run("Success - Should return no error", func(t *testing.T) {
		controller := NewProductsController(&ServiceSpy{
			CreateErr: nil,
		})

		w, ctx := util.MockGin()

		body := entity.Products{}
		faker.FakeData(&body)
		bodyByte, err := json.Marshal(body)
		assert.Nil(t, err)

		bodyByteIoReader := io.NopCloser(bytes.NewBuffer(bodyByte))
		resp, err := http.NewRequest("CREATE", "/", bodyByteIoReader)
		assert.Nil(t, err)

		ctx.Request = resp

		controller.CreateProducts(ctx)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})
	t.Run("Error - Should return StatusBadRequest", func(t *testing.T) {
		controller := NewProductsController(&ServiceSpy{
			CreateErr: errors.New("Error Status Bad Request"),
		})

		w, ctx := util.MockGin()

		resp, err := http.NewRequest("CREATE", "/", nil)
		assert.Nil(t, err)

		ctx.Request = resp

		controller.CreateProducts(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})
	t.Run("Error - Should return StatusBadRequest", func(t *testing.T) {
		controller := NewProductsController(&ServiceSpy{
			CreateErr: errors.New("Erro Internal Server Error"),
		})

		w, ctx := util.MockGin()

		body := entity.Products{}
		faker.FakeData(&body)
		bodyByte, err := json.Marshal(body)
		assert.Nil(t, err)

		bodyByteIoReader := io.NopCloser(bytes.NewBuffer(bodyByte))
		resp, err := http.NewRequest("CREATE", "/", bodyByteIoReader)
		assert.Nil(t, err)

		ctx.Request = resp

		controller.CreateProducts(ctx)

		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}

func TestUpdateProducts(t *testing.T) {
	t.Run("Success - Should return no error", func(t *testing.T) {
		controller := NewProductsController(&ServiceSpy{
			UpdateErr: nil,
		})

		w, ctx := util.MockGin()

		body := entity.Products{}
		faker.FakeData(&body)
		bodyByte, err := json.Marshal(body)
		assert.Nil(t, err)

		bodyByteIoReader := io.NopCloser(bytes.NewBuffer(bodyByte))
		resp, err := http.NewRequest("UPDATE", "/", bodyByteIoReader)
		assert.Nil(t, err)

		ctx.Request = resp

		controller.UpdateProducts(ctx)

		assert.Equal(t, http.StatusCreated, w.Result().StatusCode)
	})
	t.Run("Error - Should return error", func(t *testing.T) {
		controller := NewProductsController(&ServiceSpy{
			UpdateErr: nil,
		})

		w, ctx := util.MockGin()

		resp, err := http.NewRequest("UPDATE", "/", nil)
		assert.Nil(t, err)

		ctx.Request = resp

		controller.UpdateProducts(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})
	t.Run("Success - Should return no error", func(t *testing.T) {
		controller := NewProductsController(&ServiceSpy{
			UpdateErr: errors.New("Erro Status Internal Server Error"),
		})

		w, ctx := util.MockGin()

		body := entity.Products{}
		faker.FakeData(&body)
		bodyByte, err := json.Marshal(body)
		assert.Nil(t, err)

		bodyByteIoReader := io.NopCloser(bytes.NewBuffer(bodyByte))
		resp, err := http.NewRequest("UPDATE", "/", bodyByteIoReader)
		assert.Nil(t, err)

		ctx.Request = resp

		controller.UpdateProducts(ctx)

		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})

}

func TestDeleteProducts(t *testing.T) {
	t.Run("Success - Should return no error", func(t *testing.T) {
		controller := NewProductsController(&ServiceSpy{
			DeleteErr: nil,
		})

		w, ctx := util.MockGin()

		ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: "1"})

		resp, err := http.NewRequest("DELETE", "/", nil)
		assert.Nil(t, err)

		ctx.Request = resp

		controller.DeleteProducts(ctx)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})
	t.Run("Error - Should return StatusBadRequest error", func(t *testing.T) {
		controller := NewProductsController(&ServiceSpy{
			DeleteErr: nil,
		})

		w, ctx := util.MockGin()

		ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: "test"})

		resp, err := http.NewRequest("DELETE", "/", nil)
		assert.Nil(t, err)

		ctx.Request = resp

		controller.DeleteProducts(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})
	t.Run("Error - Should return Internal Server error", func(t *testing.T) {
		controller := NewProductsController(&ServiceSpy{
			DeleteErr: errors.New("Status Internal Server Error"),
		})

		w, ctx := util.MockGin()

		ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: "1"})

		resp, err := http.NewRequest("DELETE", "/", nil)
		assert.Nil(t, err)

		ctx.Request = resp

		controller.DeleteProducts(ctx)

		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})

}
