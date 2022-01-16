package services

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yossdev/mypoints-rest-api/src/rewards/entities"
	"github.com/yossdev/mypoints-rest-api/src/rewards/entities/mocks"
	"gorm.io/gorm"
	"testing"
	"time"
)

var (
	reward         *entities.Domain
	usecase        entities.Service
	psqlRepository mocks.PsqlRepository
)

func TestMain(m *testing.M) {
	usecase = NewRewardService(&psqlRepository)

	reward = &entities.Domain{
		ID:           1,
		AdminID:      uuid.MustParse("290fda4e-bb02-4ff5-9fd0-70f473ece15d"),
		Title:        "E-Wallet BCA 100k",
		Value:        100000,
		Points:       10000,
		Img:          "https://image.com/reward-test-img.jpg",
		Transactions: nil,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    gorm.DeletedAt{},
	}

	m.Run()
}

func TestRewardService_CreateReward(t *testing.T) {
	t.Run("Successful create reward", func(t *testing.T) {
		var mockRes interface{} = int64(1)
		var expected interface{} = int64(1)

		psqlRepository.On("Create",
			mock.AnythingOfType("entities.Domain")).Return(mockRes, nil).Once()

		payload := entities.Domain{
			AdminID: uuid.MustParse("290fda4e-bb02-4ff5-9fd0-70f473ece15d"),
			Title:   "E-Wallet BCA 100k",
			Value:   100000,
			Points:  10000,
			Img:     "https://image.com/reward-test-img.jpg",
		}

		res, err := usecase.CreateReward(payload)

		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	})

	t.Run("Something went wrong", func(t *testing.T) {
		var mockRes interface{} = int64(0)
		var expected interface{} = int64(0)

		psqlRepository.On("Create",
			mock.AnythingOfType("entities.Domain")).Return(mockRes, assert.AnError).Once()

		payload := entities.Domain{
			AdminID: uuid.MustParse("290fda4e-bb02-4ff5-9fd0-70f473ece15d"),
			Title:   "Indomie/doz",
			Value:   100000,
			Points:  10000,
			Img:     "https://image.com/reward-test-img.jpg",
		}

		res, err := usecase.CreateReward(payload)

		assert.NotNil(t, err)
		assert.Equal(t, expected, res)
	})
}

func TestRewardService_UpdateReward(t *testing.T) {
	t.Run("Reward updated", func(t *testing.T) {
		var mockRes interface{} = int64(1)
		var expected interface{} = int64(1)

		psqlRepository.On("Update",
			mock.AnythingOfType("entities.Domain")).Return(mockRes, nil).Once()

		const rewardId uint32 = 1
		payload := entities.Domain{
			Title:  "E-Wallet BCA 100k",
			Value:  100000,
			Points: 10000,
			Img:    "https://image.com/reward-test-img.jpg",
		}

		res, err := usecase.UpdateReward(rewardId, payload)

		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	})

	t.Run("Reward not found", func(t *testing.T) {
		var mockRes interface{} = int64(0)
		var expected interface{} = int64(1)

		psqlRepository.On("Update",
			mock.AnythingOfType("entities.Domain")).Return(mockRes, assert.AnError).Once()

		const rewardId uint32 = 1
		payload := entities.Domain{
			Title:  "E-Wallet BCA 100k",
			Value:  100000,
			Points: 10000,
			Img:    "https://image.com/reward-test-img.jpg",
		}

		res, err := usecase.UpdateReward(rewardId, payload)

		assert.NotNil(t, err)
		assert.NotEqual(t, expected, res)
	})
}

func TestRewardService_DeleteReward(t *testing.T) {
	t.Run("Product deleted", func(t *testing.T) {
		var mockRes interface{} = int64(1)
		var expected interface{} = int64(1)

		psqlRepository.On("Delete",
			mock.AnythingOfType("uint32")).Return(mockRes, nil).Once()

		const productId uint32 = 1
		res, err := usecase.DeleteReward(productId)

		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	})
}
