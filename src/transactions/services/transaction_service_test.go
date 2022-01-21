package services

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/xendit/xendit-go"
	entities2 "github.com/yossdev/mypoints-rest-api/src/admins/entities"
	mocks3 "github.com/yossdev/mypoints-rest-api/src/admins/entities/mocks"
	entities4 "github.com/yossdev/mypoints-rest-api/src/agents/entities"
	mocks2 "github.com/yossdev/mypoints-rest-api/src/agents/entities/mocks"
	entities3 "github.com/yossdev/mypoints-rest-api/src/rewards/entities"
	mocks4 "github.com/yossdev/mypoints-rest-api/src/rewards/entities/mocks"
	"github.com/yossdev/mypoints-rest-api/src/transactions/entities"
	"github.com/yossdev/mypoints-rest-api/src/transactions/entities/mocks"
	"github.com/yossdev/mypoints-rest-api/src/transactions/repositories"
	"testing"
	"time"
)

var (
	claim                     *entities.Domain
	redeem                    *entities.Domain
	usecase                   entities.Service
	transactionPsqlRepository mocks.PsqlRepository

	agentPsqlRepository mocks2.PsqlRepository

	adminPsqlRepository mocks3.PsqlRepository
	admin               entities2.Domain

	rewardPsqlRepository mocks4.PsqlRepository
	reward               entities3.Domain

	oriXendit func(body repositories.InvoiceBodyReq) (*xendit.Invoice, *xendit.Error)
)

func TestMain(m *testing.M) {
	usecase = NewTransactionService(
		&transactionPsqlRepository,
		&agentPsqlRepository,
		&adminPsqlRepository,
		&rewardPsqlRepository,
	)

	// preserve the original function
	oriXendit = repositories.CreateInvoice

	claim = &entities.Domain{
		ID:        uuid.New(),
		AgentID:   uuid.MustParse("03663093-440a-4fd9-8430-6dd19cb7f5ca"),
		ProductID: 1,
		Title:     "Test Product",
		Points:    100,
		NotaImg:   "https://image.com/test-nota-img.jpg",
		Type:      "Debit",
		Status:    "Pending",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	redeem = &entities.Domain{
		ID:               uuid.New(),
		AgentID:          uuid.MustParse("03663093-440a-4fd9-8430-6dd19cb7f5ca"),
		RewardID:         1,
		Title:            "Test Reward",
		Points:           100,
		RedeemInvoiceID:  "",
		RedeemInvoiceURL: "",
		RedeemDesc:       "Testing",
		Type:             "Credit",
		Status:           "Pending",
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	m.Run()
}

func TestTransactionService_Claims(t *testing.T) {
	t.Run("Create Claims", func(t *testing.T) {
		var mockRes interface{} = int64(1)
		var expected interface{} = int64(1)

		transactionPsqlRepository.On("CreateClaims",
			mock.AnythingOfType("Domain")).Return(mockRes, nil).Once()

		payload := entities.Domain{
			AgentID:   uuid.MustParse("03663093-440a-4fd9-8430-6dd19cb7f5ca"),
			ProductID: 1,
			Title:     "Test Product",
			Points:    100,
			NotaImg:   "https://image.com/test-nota-img.jpg",
		}

		res, err := usecase.Claims(payload)

		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	})

	t.Run("Something went wrong", func(t *testing.T) {
		var mockRes interface{} = int64(0)
		var expected interface{} = int64(0)

		transactionPsqlRepository.On("CreateClaims",
			mock.AnythingOfType("Domain")).Return(mockRes, assert.AnError).Once()

		payload := entities.Domain{
			AgentID:   uuid.MustParse("03663093-440a-4fd9-8430-6dd19cb7f5ca"),
			ProductID: 1,
			Title:     "Test Product",
			Points:    100,
			NotaImg:   "https://image.com/test-nota-img.jpg",
		}

		res, err := usecase.Claims(payload)

		assert.NotNil(t, err)
		assert.Equal(t, expected, res)
	})
}

func TestTransactionService_ClaimsStatus(t *testing.T) {
	t.Run("Status Approved", func(t *testing.T) {
		var mockRes interface{} = int64(1)
		var expected interface{} = int64(1)

		transactionPsqlRepository.On("GetTransaction",
			mock.AnythingOfType("string")).Return(*claim, nil).Once()

		agentPsqlRepository.On("UpdatePoints",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("int32")).Return(mockRes, nil).Once()

		transactionPsqlRepository.On("UpdateClaimsStatus",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("string")).Return(mockRes, nil).Once()

		var agentId = uuid.MustParse("03663093-440a-4fd9-8430-6dd19cb7f5ca")
		var status = "Approved"
		res, err := usecase.ClaimsStatus(agentId, status)

		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	})

	t.Run("Something went wrong when updating points", func(t *testing.T) {
		var mockRes interface{} = int64(0)
		var expected interface{} = int64(0)

		transactionPsqlRepository.On("GetTransaction",
			mock.AnythingOfType("string")).Return(*claim, nil).Once()

		agentPsqlRepository.On("UpdatePoints",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("int32")).Return(mockRes, assert.AnError).Once()

		var agentId = uuid.MustParse("03663093-440a-4fd9-8430-6dd19cb7f5ca")
		var status = "Approved"
		res, err := usecase.ClaimsStatus(agentId, status)

		assert.NotNil(t, err)
		assert.Equal(t, expected, res)
	})

	t.Run("Transaction not found", func(t *testing.T) {
		var expected interface{} = int64(0)

		transactionPsqlRepository.On("GetTransaction",
			mock.AnythingOfType("string")).Return(*claim, assert.AnError).Once()

		var agentId = uuid.MustParse("03663093-440a-4fd9-8430-6dd19cb7f5ca")
		var status = "Approved"
		res, err := usecase.ClaimsStatus(agentId, status)

		assert.NotNil(t, err)
		assert.Equal(t, expected, res)
	})

	t.Run("Already Approved", func(t *testing.T) {
		var expected interface{} = int64(0)

		claim.Status = "Approved"

		transactionPsqlRepository.On("GetTransaction",
			mock.AnythingOfType("string")).Return(*claim, nil).Once()

		var agentId = uuid.MustParse("03663093-440a-4fd9-8430-6dd19cb7f5ca")
		var status = "Rejected"
		res, err := usecase.ClaimsStatus(agentId, status)

		assert.NotNil(t, err)
		assert.Equal(t, expected, res)
	})
}

func TestTransactionService_Redeem(t *testing.T) {
	t.Run("Redeem Successful", func(t *testing.T) {
		repositories.CreateInvoice = func(body repositories.InvoiceBodyReq) (*xendit.Invoice, *xendit.Error) {
			invoice := &xendit.Invoice{
				ID:         "1f1f0c8e-da72-4230-a8e7-af4be45b1bba",
				InvoiceURL: "https://xendit.com/invoice-redeem",
			}
			return invoice, nil
		}

		admin = entities2.Domain{
			Name:  "admin1",
			Email: "admin1@mail.com",
			Agents: []entities4.Domain{
				{
					Points: 200,
				},
			},
		}

		reward = entities3.Domain{
			Points: 100,
			Value:  10000,
		}

		var mockRes interface{} = int64(1)
		var expected interface{} = int64(1)

		adminPsqlRepository.On("GetAdminByAgentID",
			mock.AnythingOfType("uuid.UUID")).Return(admin, nil).Once()

		rewardPsqlRepository.On("GetReward",
			mock.AnythingOfType("uint32")).Return(reward, nil).Once()

		transactionPsqlRepository.On("CreateRedeem",
			mock.AnythingOfType("Domain")).Return(mockRes, nil).Once()

		agentPsqlRepository.On("UpdatePoints",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("int32")).Return(mockRes, nil).Once()

		payload := entities.Domain{
			AgentID:    uuid.MustParse("03663093-440a-4fd9-8430-6dd19cb7f5ca"),
			RewardID:   1,
			Title:      "Test Reward",
			Points:     100,
			RedeemDesc: "Testing",
		}

		res, err := usecase.Redeem(payload)

		assert.Nil(t, err)
		assert.Equal(t, expected, res)
	})

	t.Run("Not enough points", func(t *testing.T) {
		repositories.CreateInvoice = func(body repositories.InvoiceBodyReq) (*xendit.Invoice, *xendit.Error) {
			invoice := &xendit.Invoice{
				ID:         "1f1f0c8e-da72-4230-a8e7-af4be45b1bba",
				InvoiceURL: "https://xendit.com/invoice-redeem",
			}
			return invoice, nil
		}

		admin = entities2.Domain{
			Name:  "admin1",
			Email: "admin1@mail.com",
			Agents: []entities4.Domain{
				{
					Points: 50,
				},
			},
		}

		reward = entities3.Domain{
			Points: 100,
			Value:  10000,
		}

		var expected interface{} = int64(0)

		adminPsqlRepository.On("GetAdminByAgentID",
			mock.AnythingOfType("uuid.UUID")).Return(admin, nil).Once()

		rewardPsqlRepository.On("GetReward",
			mock.AnythingOfType("uint32")).Return(reward, nil).Once()

		payload := entities.Domain{
			AgentID:    uuid.MustParse("03663093-440a-4fd9-8430-6dd19cb7f5ca"),
			RewardID:   1,
			Title:      "Test Reward",
			Points:     100,
			RedeemDesc: "Testing",
		}

		res, err := usecase.Redeem(payload)

		assert.NotNil(t, err)
		assert.Equal(t, expected, res)
	})

	t.Run("Error xendit", func(t *testing.T) {
		repositories.CreateInvoice = func(body repositories.InvoiceBodyReq) (*xendit.Invoice, *xendit.Error) {
			invoice := &xendit.Invoice{
				ID:         "1f1f0c8e-da72-4230-a8e7-af4be45b1bba",
				InvoiceURL: "https://xendit.com/invoice-redeem",
			}

			err := &xendit.Error{
				Status:    400,
				ErrorCode: "Bad Request",
				Message:   "Bad Request",
			}

			return invoice, err
		}

		admin = entities2.Domain{
			Name:  "admin1",
			Email: "admin1@mail.com",
			Agents: []entities4.Domain{
				{
					Points: 200,
				},
			},
		}

		reward = entities3.Domain{
			Points: 100,
			Value:  10000,
		}

		var expected interface{} = int64(0)

		adminPsqlRepository.On("GetAdminByAgentID",
			mock.AnythingOfType("uuid.UUID")).Return(admin, nil).Once()

		rewardPsqlRepository.On("GetReward",
			mock.AnythingOfType("uint32")).Return(reward, nil).Once()

		payload := entities.Domain{
			AgentID:    uuid.MustParse("03663093-440a-4fd9-8430-6dd19cb7f5ca"),
			RewardID:   1,
			Title:      "Test Reward",
			Points:     100,
			RedeemDesc: "Testing",
		}

		res, err := usecase.Redeem(payload)

		assert.NotNil(t, err)
		assert.Equal(t, expected, res)
	})

	t.Run("Error create redeem", func(t *testing.T) {
		repositories.CreateInvoice = func(body repositories.InvoiceBodyReq) (*xendit.Invoice, *xendit.Error) {
			invoice := &xendit.Invoice{
				ID:         "1f1f0c8e-da72-4230-a8e7-af4be45b1bba",
				InvoiceURL: "https://xendit.com/invoice-redeem",
			}
			return invoice, nil
		}

		admin = entities2.Domain{
			Name:  "admin1",
			Email: "admin1@mail.com",
			Agents: []entities4.Domain{
				{
					Points: 200,
				},
			},
		}

		reward = entities3.Domain{
			Points: 100,
			Value:  10000,
		}

		var mockRes interface{} = int64(0)
		var expected interface{} = int64(0)

		adminPsqlRepository.On("GetAdminByAgentID",
			mock.AnythingOfType("uuid.UUID")).Return(admin, nil).Once()

		rewardPsqlRepository.On("GetReward",
			mock.AnythingOfType("uint32")).Return(reward, nil).Once()

		transactionPsqlRepository.On("CreateRedeem",
			mock.AnythingOfType("Domain")).Return(mockRes, assert.AnError).Once()

		payload := entities.Domain{
			AgentID:    uuid.MustParse("03663093-440a-4fd9-8430-6dd19cb7f5ca"),
			RewardID:   1,
			Title:      "Test Reward",
			Points:     100,
			RedeemDesc: "Testing",
		}

		res, err := usecase.Redeem(payload)

		assert.NotNil(t, err)
		assert.Equal(t, expected, res)
	})

	t.Run("Error update points", func(t *testing.T) {
		repositories.CreateInvoice = func(body repositories.InvoiceBodyReq) (*xendit.Invoice, *xendit.Error) {
			invoice := &xendit.Invoice{
				ID:         "1f1f0c8e-da72-4230-a8e7-af4be45b1bba",
				InvoiceURL: "https://xendit.com/invoice-redeem",
			}
			return invoice, nil
		}

		admin = entities2.Domain{
			Name:  "admin1",
			Email: "admin1@mail.com",
			Agents: []entities4.Domain{
				{
					Points: 200,
				},
			},
		}

		reward = entities3.Domain{
			Points: 100,
			Value:  10000,
		}

		var mockRes interface{} = int64(0)
		var expected interface{} = int64(0)

		adminPsqlRepository.On("GetAdminByAgentID",
			mock.AnythingOfType("uuid.UUID")).Return(admin, nil).Once()

		rewardPsqlRepository.On("GetReward",
			mock.AnythingOfType("uint32")).Return(reward, nil).Once()

		transactionPsqlRepository.On("CreateRedeem",
			mock.AnythingOfType("Domain")).Return(mockRes, nil).Once()

		agentPsqlRepository.On("UpdatePoints",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("int32")).Return(mockRes, assert.AnError).Once()

		payload := entities.Domain{
			AgentID:    uuid.MustParse("03663093-440a-4fd9-8430-6dd19cb7f5ca"),
			RewardID:   1,
			Title:      "Test Reward",
			Points:     100,
			RedeemDesc: "Testing",
		}

		res, err := usecase.Redeem(payload)

		assert.NotNil(t, err)
		assert.Equal(t, expected, res)
	})

	repositories.CreateInvoice = oriXendit
}

func TestTransactionService_CallbackXendit(t *testing.T) {
	t.Run("Transaction PAID", func(t *testing.T) {
		transactionPsqlRepository.On("GetTransaction",
			mock.AnythingOfType("string")).Return(*redeem, nil).Once()

		transactionPsqlRepository.On("UpdateRedeemStatus",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(nil).Once()

		payload := entities.InvoiceCallback{
			ID:           "1f1f0c8e-da72-4230-a8e7-af4be45b1bba",
			Status:       "PAID",
			MerchantName: "MyPoints",
			Amount:       10000,
		}

		err := usecase.CallbackXendit(payload)

		assert.Nil(t, err)
	})

	t.Run("Error while get transaction", func(t *testing.T) {
		transactionPsqlRepository.On("GetTransaction",
			mock.AnythingOfType("string")).Return(*redeem, assert.AnError).Once()

		payload := entities.InvoiceCallback{
			ID:           "1f1f0c8e-da72-4230-a8e7-af4be45b1bba",
			Status:       "PAID",
			MerchantName: "MyPoints",
			Amount:       10000,
		}

		err := usecase.CallbackXendit(payload)

		assert.NotNil(t, err)
	})

	t.Run("Transaction EXPIRED", func(t *testing.T) {
		transactionPsqlRepository.On("GetTransaction",
			mock.AnythingOfType("string")).Return(*redeem, nil).Once()

		var mockRes interface{} = int64(1)
		agentPsqlRepository.On("UpdatePoints",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("int32")).Return(mockRes, nil).Once()

		transactionPsqlRepository.On("UpdateRedeemStatus",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(nil).Once()

		payload := entities.InvoiceCallback{
			ID:           "1f1f0c8e-da72-4230-a8e7-af4be45b1bba",
			Status:       "EXPIRED",
			MerchantName: "MyPoints",
			Amount:       10000,
		}

		err := usecase.CallbackXendit(payload)

		assert.Nil(t, err)
	})

	t.Run("Error while update points", func(t *testing.T) {
		transactionPsqlRepository.On("GetTransaction",
			mock.AnythingOfType("string")).Return(*redeem, nil).Once()

		var mockRes interface{} = int64(1)
		agentPsqlRepository.On("UpdatePoints",
			mock.AnythingOfType("uuid.UUID"),
			mock.AnythingOfType("int32")).Return(mockRes, assert.AnError).Once()

		payload := entities.InvoiceCallback{
			ID:           "1f1f0c8e-da72-4230-a8e7-af4be45b1bba",
			Status:       "EXPIRED",
			MerchantName: "MyPoints",
			Amount:       10000,
		}

		err := usecase.CallbackXendit(payload)

		assert.NotNil(t, err)
	})

	t.Run("Transaction already settled", func(t *testing.T) {
		redeem.Status = "Settled"
		transactionPsqlRepository.On("GetTransaction",
			mock.AnythingOfType("string")).Return(*redeem, nil).Once()

		payload := entities.InvoiceCallback{
			ID:           "1f1f0c8e-da72-4230-a8e7-af4be45b1bba",
			Status:       "EXPIRED",
			MerchantName: "MyPoints",
			Amount:       10000,
		}

		err := usecase.CallbackXendit(payload)

		assert.NotNil(t, err)
	})

	t.Run("Transaction already expired", func(t *testing.T) {
		redeem.Status = "Expired"
		transactionPsqlRepository.On("GetTransaction",
			mock.AnythingOfType("string")).Return(*redeem, nil).Once()

		payload := entities.InvoiceCallback{
			ID:           "1f1f0c8e-da72-4230-a8e7-af4be45b1bba",
			Status:       "EXPIRED",
			MerchantName: "MyPoints",
			Amount:       10000,
		}

		err := usecase.CallbackXendit(payload)

		assert.NotNil(t, err)
	})

}
