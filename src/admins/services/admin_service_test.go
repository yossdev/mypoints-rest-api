package services

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yossdev/mypoints-rest-api/internal/utils/auth"
	"github.com/yossdev/mypoints-rest-api/internal/utils/helpers"
	"github.com/yossdev/mypoints-rest-api/src/admins/entities"
	"github.com/yossdev/mypoints-rest-api/src/admins/entities/mocks"
	"testing"
	"time"
)

var (
	admin           *entities.Domain
	usecase         entities.Service
	psqlRepository  mocks.PsqlRepository
	oriHash         func(password string) (string, error)
	oriValidateHash func(a, b string) error
	oriSign         func(claims jwt.MapClaims) auth.Token
)

func TestMain(m *testing.M) {
	usecase = NewAdminService(&psqlRepository)

	// preserve the original function
	oriSign = auth.Sign
	oriHash = helpers.Hash
	oriValidateHash = helpers.ValidateHash

	admin = &entities.Domain{
		ID:        uuid.MustParse("290fda4e-bb02-4ff5-9fd0-70f473ece15d"),
		Name:      "admin1",
		Email:     "admin1@gmail.com",
		Password:  "hashedPassword",
		Img:       "",
		Agents:    nil,
		Products:  nil,
		Rewards:   nil,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	m.Run()
}

func TestAdminService_SignIn(t *testing.T) {
	t.Run("Successful login", func(t *testing.T) {
		psqlRepository.On("SignInWithEmail",
			mock.AnythingOfType("string")).Return([2]string{admin.Password, admin.ID.String()}, nil).Once()

		helpers.ValidateHash = func(a, b string) error {
			return nil
		}

		auth.Sign = func(claims jwt.MapClaims) auth.Token {
			return auth.Token{
				AccessToken:  "access token",
				RefreshToken: "refresh token",
			}
		}

		token, err := usecase.SignIn(entities.Domain{
			Email:    "admin1@gmail.com",
			Password: "admin1pass",
		})

		assert.Nil(t, err)
		assert.NotEmpty(t, token.AccessToken)
	})

	t.Run("Wrong email", func(t *testing.T) {
		psqlRepository.On("SignInWithEmail",
			mock.AnythingOfType("string")).Return([2]string{"", ""}, assert.AnError).Once()

		token, err := usecase.SignIn(entities.Domain{
			Email:    "admin123@gmail.com",
			Password: "admin1pass",
		})

		assert.NotNil(t, err)
		assert.Empty(t, token.AccessToken)
	})

	t.Run("Wrong password", func(t *testing.T) {
		helpers.ValidateHash = func(a, b string) error {
			return assert.AnError
		}

		psqlRepository.On("SignInWithEmail",
			mock.AnythingOfType("string")).Return([2]string{admin.Password, admin.ID.String()}, nil).Once()

		token, err := usecase.SignIn(entities.Domain{
			Email:    "admin1@gmail.com",
			Password: "pass123",
		})

		assert.NotNil(t, err)
		assert.Empty(t, token.AccessToken)
	})

	auth.Sign = oriSign
	helpers.Hash = oriHash
	helpers.ValidateHash = oriValidateHash
}

func TestAdminService_SignUp(t *testing.T) {
	t.Run("Successful Register", func(t *testing.T) {
		helpers.Hash = func(password string) (string, error) {
			return "hashedPassword", nil
		}

		var mockRes interface{} = int64(1)
		var expected interface{} = int64(1)

		psqlRepository.On("CreateAdmin", mock.AnythingOfType("*entities.Domain")).Return(mockRes, nil).Once()

		payload := entities.Domain{
			Name:     "admin1",
			Email:    "admin1@gmail.com",
			Password: "admin1pass",
			Img:      "",
		}

		res, err := usecase.SignUp(&payload)

		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	})

	t.Run("Account already exist", func(t *testing.T) {
		helpers.Hash = func(password string) (string, error) {
			return "hashedPassword", nil
		}

		var mockRes interface{} = int64(0)
		var expected interface{} = int64(0)

		psqlRepository.On("CreateAdmin", mock.AnythingOfType("*entities.Domain")).Return(mockRes, assert.AnError).Once()

		payload := entities.Domain{
			Name:     "admin1",
			Email:    "admin1@gmail.com",
			Password: "admin1pass",
			Img:      "",
		}

		res, err := usecase.SignUp(&payload)

		assert.NotNil(t, err)
		assert.Equal(t, expected, res)
	})

	helpers.Hash = oriHash
}

func TestAdminService_UpdateAdmin(t *testing.T) {
	t.Run("Update success", func(t *testing.T) {
		helpers.Hash = func(password string) (string, error) {
			return "hashedPassword", nil
		}

		var mockRes interface{} = int64(1)
		var expected interface{} = int64(1)

		psqlRepository.On("UpdateAdmin", mock.AnythingOfType("Domain")).Return(mockRes, nil).Once()

		payload := entities.Domain{
			Name:     "admin1",
			Email:    "admin1@gmail.com",
			Password: "admin1pass",
		}

		id := uuid.MustParse("290fda4e-bb02-4ff5-9fd0-70f473ece15d")
		res, err := usecase.UpdateAdmin(id, payload)

		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	})

	t.Run("Email already used", func(t *testing.T) {
		var mockRes interface{} = int64(0)
		var expected interface{} = int64(0)

		psqlRepository.On("UpdateAdmin", mock.AnythingOfType("Domain")).Return(mockRes, assert.AnError).Once()

		payload := entities.Domain{
			Name:  "admin1",
			Email: "admin2@gmail.com",
		}

		id := uuid.MustParse("290fda4e-bb02-4ff5-9fd0-70f473ece15d")
		res, err := usecase.UpdateAdmin(id, payload)

		assert.NotNil(t, err)
		assert.Equal(t, expected, res)
	})

	helpers.Hash = oriHash
}

func TestAdminService_UpdateAvatar(t *testing.T) {
	t.Run("Avatar successfully updated", func(t *testing.T) {
		var mockRes interface{} = int64(1)
		var expected interface{} = int64(1)

		psqlRepository.On("UpdateAvatar", mock.AnythingOfType("Domain")).Return(mockRes, nil).Once()

		payload := entities.Domain{
			Img: "https://myavatar.com/avatar",
		}

		id := uuid.MustParse("290fda4e-bb02-4ff5-9fd0-70f473ece15d")
		res, err := usecase.UpdateAvatar(id, payload)

		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	})
}
