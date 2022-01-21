package services

import (
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/internal/web"
	_admin "github.com/yossdev/mypoints-rest-api/src/admins/entities"
	_agent "github.com/yossdev/mypoints-rest-api/src/agents/entities"
	_reward "github.com/yossdev/mypoints-rest-api/src/rewards/entities"
	"github.com/yossdev/mypoints-rest-api/src/transactions/entities"
	_xendit "github.com/yossdev/mypoints-rest-api/src/transactions/repositories"
)

type transactionService struct {
	transactionPsqlRepository entities.PsqlRepository
	agentPsqlRepository       _agent.PsqlRepository
	adminPsqlRepository       _admin.PsqlRepository
	rewardPsqlRepository      _reward.PsqlRepository
}

func NewTransactionService(
	t entities.PsqlRepository, ag _agent.PsqlRepository, ad _admin.PsqlRepository, r _reward.PsqlRepository,
) entities.Service {
	return &transactionService{
		transactionPsqlRepository: t,
		agentPsqlRepository:       ag,
		adminPsqlRepository:       ad,
		rewardPsqlRepository:      r,
	}
}

func (s *transactionService) Claims(payload entities.Domain) (int64, error) {
	res, err := s.transactionPsqlRepository.CreateClaims(payload)
	return res, err
}

func (s *transactionService) ClaimsStatus(id uuid.UUID, status string) (int64, error) {
	// get by agent id
	t, e := s.transactionPsqlRepository.GetTransaction(id.String())
	if e != nil {
		return 0, e
	}

	if t.Status == "Approved" {
		return 0, web.AlreadyApproved
	} else if status == "Approved" {
		_, err := s.agentPsqlRepository.UpdatePoints(t.AgentID, int32(t.Points))
		if err != nil {
			return 0, err
		}
	}

	res, err := s.transactionPsqlRepository.UpdateClaimsStatus(id, status)
	return res, err
}

func (s *transactionService) Redeem(payload entities.Domain) (int64, error) {
	admin, _ := s.adminPsqlRepository.GetAdminByAgentID(payload.AgentID)
	reward, _ := s.rewardPsqlRepository.GetReward(payload.RewardID)

	agentPoints := admin.Agents[0].Points
	rewardPoints := reward.Points
	if agentPoints < rewardPoints {
		return 0, web.NotEnoughPoints
	}

	body := _xendit.InvoiceBodyReq{
		Name:  admin.Name,
		Email: admin.Email,
		Value: float64(reward.Value),
		Title: payload.Title,
		Desc:  payload.RedeemDesc,
	}

	// invoice will be created by xendit
	invoice, e := _xendit.CreateInvoice(body)
	if e != nil {
		return 0, e
	}
	//fmt.Printf("created invoice: %+v\n", invoice)

	payload.RedeemInvoiceID = invoice.ID
	payload.RedeemInvoiceURL = invoice.InvoiceURL

	// placeholder for both repo below it
	var res int64
	var err error

	res, err = s.transactionPsqlRepository.CreateRedeem(payload)
	if err != nil {
		return 0, err
	}

	res, err = s.agentPsqlRepository.UpdatePoints(payload.AgentID, -int32(rewardPoints))
	if err != nil {
		return 0, err
	}

	return res, err
}

func (s *transactionService) CallbackXendit(payload entities.InvoiceCallback) error {
	// get by reward id
	t, e := s.transactionPsqlRepository.GetTransaction(payload.ID)
	if e != nil {
		return e
	}

	if t.Status == "Settled" {
		return web.AlreadySettled
	} else if t.Status == "Expired" {
		return web.TransactionExpired
	}

	var status string
	if payload.Status == "PAID" {
		status = "Settled"
	} else {
		status = "Expired"
	}

	if payload.Status == "EXPIRED" {
		if _, err := s.agentPsqlRepository.UpdatePoints(t.AgentID, int32(t.Points)); err != nil {
			return err
		}
	}

	return s.transactionPsqlRepository.UpdateRedeemStatus(payload.ID, status)
}
