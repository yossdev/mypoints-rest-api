package services

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yossdev/mypoints-rest-api/internal/utils/auth"
	"github.com/yossdev/mypoints-rest-api/internal/utils/helpers"
	"github.com/yossdev/mypoints-rest-api/src/agents/entities"
	"github.com/yossdev/mypoints-rest-api/src/agents/entities/mocks"
	"testing"
	"time"
)

var (
	agent           *entities.Domain
	usecase         entities.Service
	psqlRepository  mocks.PsqlRepository
	oriHash         func(password string) (string, error)
	oriValidateHash func(a, b string) error
	oriSign         func(claims jwt.MapClaims) auth.Token
)

func TestMain(m *testing.M) {
	usecase = NewAgentService(&psqlRepository)

	// preserve the original function
	oriSign = auth.Sign
	oriHash = helpers.Hash
	oriValidateHash = helpers.ValidateHash

	agent = &entities.Domain{
		ID:           uuid.MustParse("03663093-440a-4fd9-8430-6dd19cb7f5ca"),
		AdminID:      uuid.MustParse("290fda4e-bb02-4ff5-9fd0-70f473ece15d"),
		Name:         "agent1",
		Email:        "agent1@gmail.com",
		Password:     "hashedPassword",
		Points:       0,
		Img:          "",
		Active:       true,
		Transactions: nil,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	m.Run()
}

func TestAgentService_SignIn(t *testing.T) {
	t.Run("Successful login", func(t *testing.T) {
		psqlRepository.On("SignInWithEmail",
			mock.AnythingOfType("string")).Return(*agent, nil).Once()

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
			Email:    "agent1@gmail.com",
			Password: "agent1pass",
		})

		assert.Nil(t, err)
		assert.NotEmpty(t, token.AccessToken)
	})

	t.Run("Wrong email", func(t *testing.T) {
		psqlRepository.On("SignInWithEmail",
			mock.AnythingOfType("string")).Return(entities.Domain{}, assert.AnError).Once()

		token, err := usecase.SignIn(entities.Domain{
			Email:    "agent123@gmail.com",
			Password: "agent1pass",
		})

		assert.NotNil(t, err)
		assert.Empty(t, token.AccessToken)
	})

	t.Run("Wrong password", func(t *testing.T) {
		helpers.ValidateHash = func(a, b string) error {
			return assert.AnError
		}

		psqlRepository.On("SignInWithEmail",
			mock.AnythingOfType("string")).Return(*agent, nil).Once()

		token, err := usecase.SignIn(entities.Domain{
			Email:    "agent1@gmail.com",
			Password: "agent1pass123",
		})

		assert.NotNil(t, err)
		assert.Empty(t, token.AccessToken)
	})

	t.Run("Account disabled", func(t *testing.T) {
		agent = &entities.Domain{
			ID:           uuid.MustParse("03663093-440a-4fd9-8430-6dd19cb7f5ca"),
			AdminID:      uuid.MustParse("290fda4e-bb02-4ff5-9fd0-70f473ece15d"),
			Name:         "agent1",
			Email:        "agent1@gmail.com",
			Password:     "hashedPassword",
			Points:       0,
			Img:          "",
			Active:       false,
			Transactions: nil,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		psqlRepository.On("SignInWithEmail",
			mock.AnythingOfType("string")).Return(*agent, nil).Once()

		token, err := usecase.SignIn(entities.Domain{
			Email:    "agent1@gmail.com",
			Password: "agent1pass",
		})

		assert.NotNil(t, err)
		assert.Equal(t, false, agent.Active)
		assert.Empty(t, token.AccessToken)
	})

	auth.Sign = oriSign
	helpers.Hash = oriHash
	helpers.ValidateHash = oriValidateHash
}

func TestAgentService_SignUp(t *testing.T) {
	t.Run("Successful Register", func(t *testing.T) {
		helpers.Hash = func(password string) (string, error) {
			return "hashedPassword", nil
		}

		var mockRes interface{} = int64(1)
		var expected interface{} = int64(1)

		psqlRepository.On("CreateAgent", mock.AnythingOfType("*entities.Domain")).Return(mockRes, nil).Once()

		payload := entities.Domain{
			AdminID:  uuid.MustParse("290fda4e-bb02-4ff5-9fd0-70f473ece15d"),
			Name:     "agent1",
			Email:    "agent1@gmail.com",
			Password: "agent1pass",
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

		psqlRepository.On("CreateAgent", mock.AnythingOfType("*entities.Domain")).Return(mockRes, assert.AnError).Once()

		payload := entities.Domain{
			AdminID:  uuid.MustParse("290fda4e-bb02-4ff5-9fd0-70f473ece15d"),
			Name:     "agent1",
			Email:    "agent1@gmail.com",
			Password: "agent1pass",
			Img:      "",
		}

		res, err := usecase.SignUp(&payload)

		assert.NotNil(t, err)
		assert.Equal(t, expected, res)
	})

	helpers.Hash = oriHash
}

func TestAgentService_GetAgent(t *testing.T) {
	t.Run("Data found", func(t *testing.T) {
		psqlRepository.On("GetAgent", mock.AnythingOfType("uuid.UUID")).Return(*agent, nil).Once()

		id := uuid.MustParse("03663093-440a-4fd9-8430-6dd19cb7f5ca")
		res, err := usecase.GetAgent(id)

		assert.Nil(t, err)
		assert.Equal(t, res.Name, agent.Name)
	})

	t.Run("Data not found", func(t *testing.T) {
		psqlRepository.On("GetAgent", mock.AnythingOfType("uuid.UUID")).Return(entities.Domain{}, assert.AnError).Once()

		id := uuid.MustParse("feabe110-7c91-408f-a60a-21fabe559ac8")
		res, err := usecase.GetAgent(id)

		assert.NotNil(t, err)
		assert.NotEqual(t, "feabe110-7c91-408f-a60a-21fabe559ac8", res.ID)
		assert.Empty(t, res)
	})
}

func TestAgentService_UpdateAgent(t *testing.T) {
	t.Run("Update success", func(t *testing.T) {
		helpers.Hash = func(password string) (string, error) {
			return "hashedPassword", nil
		}

		var mockRes interface{} = int64(1)
		var expected interface{} = int64(1)

		psqlRepository.On("UpdateAgent", mock.AnythingOfType("Domain")).Return(mockRes, nil).Once()

		payload := entities.Domain{
			Name:     "agent1",
			Email:    "agent1@gmail.com",
			Password: "agent1pass",
		}

		id := uuid.MustParse("03663093-440a-4fd9-8430-6dd19cb7f5ca")
		res, err := usecase.UpdateAgent(id, payload)

		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	})

	t.Run("Email already used", func(t *testing.T) {
		var mockRes interface{} = int64(1)
		var expected interface{} = int64(1)

		psqlRepository.On("UpdateAgent", mock.AnythingOfType("Domain")).Return(mockRes, assert.AnError).Once()

		payload := entities.Domain{
			Name:     "agent1",
			Email:    "agent1@gmail.com",
			Password: "agent1pass",
		}

		id := uuid.MustParse("03663093-440a-4fd9-8430-6dd19cb7f5ca")
		res, err := usecase.UpdateAgent(id, payload)

		assert.NotNil(t, err)
		assert.Equal(t, expected, res)
	})

	helpers.Hash = oriHash
}

func TestAgentService_UpdateAvatar(t *testing.T) {
	t.Run("Avatar successfully updated", func(t *testing.T) {
		var mockRes interface{} = int64(1)
		var expected interface{} = int64(1)

		psqlRepository.On("UpdateAvatar", mock.AnythingOfType("Domain")).Return(mockRes, nil).Once()

		payload := entities.Domain{
			Img: "https://myavatar.com/avatar",
		}

		id := uuid.MustParse("03663093-440a-4fd9-8430-6dd19cb7f5ca")
		res, err := usecase.UpdateAvatar(id, payload)

		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	})
}

func TestAgentService_UpdateAgentByAdmin(t *testing.T) {
	t.Run("Agent account successfully updated", func(t *testing.T) {
		helpers.Hash = func(password string) (string, error) {
			return "hashedPassword", nil
		}

		var mockRes interface{} = int64(1)
		var expected interface{} = int64(1)

		psqlRepository.On("UpdateAgentByAdmin", mock.AnythingOfType("Domain")).Return(mockRes, nil).Once()

		payload := entities.Domain{
			ID:       uuid.MustParse("03663093-440a-4fd9-8430-6dd19cb7f5ca"),
			Password: "",
			Active:   false,
		}

		res, err := usecase.UpdateAgentByAdmin(payload)

		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	})

	t.Run("Agent account password successfully updated", func(t *testing.T) {
		helpers.Hash = func(password string) (string, error) {
			return "hashedPassword", nil
		}

		var mockRes interface{} = int64(1)
		var expected interface{} = int64(1)

		psqlRepository.On("UpdateAgentByAdmin", mock.AnythingOfType("Domain")).Return(mockRes, nil).Once()

		payload := entities.Domain{
			ID:       uuid.MustParse("03663093-440a-4fd9-8430-6dd19cb7f5ca"),
			Password: "agent1newpass",
			Active:   true,
		}

		res, err := usecase.UpdateAgentByAdmin(payload)

		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	})

	helpers.Hash = oriHash
}
