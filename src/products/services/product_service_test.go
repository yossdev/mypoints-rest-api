package services

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yossdev/mypoints-rest-api/src/products/entities"
	"github.com/yossdev/mypoints-rest-api/src/products/entities/mocks"
	"testing"
)

var (
	//product        *entities.Domain
	usecase        entities.Service
	psqlRepository mocks.PsqlRepository
)

func TestMain(m *testing.M) {
	usecase = NewProductService(&psqlRepository)

	//product = &entities.Domain{
	//	ID:           1,
	//	AdminID:      uuid.MustParse("290fda4e-bb02-4ff5-9fd0-70f473ece15d"),
	//	Title:        "Indomie/doz",
	//	Points:       100,
	//	Img:          "https://image.com/test-img.jpg",
	//	Transactions: nil,
	//	CreatedAt:    time.Now(),
	//	UpdatedAt:    time.Now(),
	//	DeletedAt:    gorm.DeletedAt{},
	//}

	m.Run()
}

func TestProductService_CreateProduct(t *testing.T) {
	t.Run("Successful create product", func(t *testing.T) {
		var mockRes interface{} = int64(1)
		var expected interface{} = int64(1)

		psqlRepository.On("Create",
			mock.AnythingOfType("Domain")).Return(mockRes, nil).Once()

		payload := entities.Domain{
			AdminID: uuid.MustParse("290fda4e-bb02-4ff5-9fd0-70f473ece15d"),
			Title:   "Indomie/doz",
			Points:  100,
			Img:     "https://image.com/test-img.jpg",
		}

		res, err := usecase.CreateProduct(payload)

		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	})

	t.Run("Something went wrong", func(t *testing.T) {
		var mockRes interface{} = int64(0)
		var expected interface{} = int64(0)

		psqlRepository.On("Create",
			mock.AnythingOfType("Domain")).Return(mockRes, assert.AnError).Once()

		payload := entities.Domain{
			AdminID: uuid.MustParse("290fda4e-bb02-4ff5-9fd0-70f473ece15d"),
			Title:   "Indomie/doz",
			Points:  100,
			Img:     "https://image.com/test-img.jpg",
		}

		res, err := usecase.CreateProduct(payload)

		assert.NotNil(t, err)
		assert.Equal(t, expected, res)
	})
}

func TestProductService_UpdateProduct(t *testing.T) {
	t.Run("Product updated", func(t *testing.T) {
		var mockRes interface{} = int64(1)
		var expected interface{} = int64(1)

		psqlRepository.On("Update",
			mock.AnythingOfType("Domain")).Return(mockRes).Once()

		const productId uint32 = 1
		payload := entities.Domain{
			Title:  "Indomie/doz",
			Points: 120,
			Img:    "https://image.com/test-img.jpg",
		}

		res := usecase.UpdateProduct(productId, payload)

		assert.Equal(t, expected, res)
	})

	t.Run("Product not found", func(t *testing.T) {
		var mockRes interface{} = int64(0)
		var expected interface{} = int64(1)

		psqlRepository.On("Update",
			mock.AnythingOfType("Domain")).Return(mockRes).Once()

		const productId uint32 = 2
		payload := entities.Domain{
			Title:  "Indomie/doz",
			Points: 120,
			Img:    "https://image.com/test-img.jpg",
		}

		res := usecase.UpdateProduct(productId, payload)

		assert.NotEqual(t, expected, res)
	})
}

func TestProductService_DeleteProduct(t *testing.T) {
	t.Run("Product deleted", func(t *testing.T) {
		var mockRes interface{} = int64(1)
		var expected interface{} = int64(1)

		psqlRepository.On("Delete",
			mock.AnythingOfType("uint32")).Return(mockRes).Once()

		const productId uint32 = 1
		res := usecase.DeleteProduct(productId)

		assert.Equal(t, expected, res)
	})
}
