package services

import (
	"context" 

	"github.com/DevShuxat/eater-service/EATER-SERVICE/src/domain/eater/models"
	"github.com/DevShuxat/eater-service/EATER-SERVICE/src/domain/eater/repositories"
	"github.com/DevShuxat/eater-service/EATER-SERVICE/src/infrastructure/sms"
	"github.com/DevShuxat/eater-service/EATER-SERVICE/src/infrastructure/rand"
	"go.uber.org/zap"
)

type EaterService interface {
	SignupEater(ctx context.Context, phoneNumber string) (string, error)
	ConfirmSMSCode(ctx context.Context, eaterID, smsCode string) (*models.EaterProfile, error)
	GetEaterProfile(ctx context.Context, eaterID string) (*models.EaterProfile, error)
	UpdateEaterProfile(ctx context.Context, eaterID, name, imageUrl string) (*models.EaterProfile, error)
}


func NewEaterService(
	eaterRepo repositories.EaterRepository,
	smsClient sms.Client,
	logger *zap.Logger,
) EaterService {
	return &eaterSvcImpl{
		eaterRepo: eaterRepo,
		smsClient: smsClient,
		logger:    logger,
	}
}

type eaterSvcImpl struct {
	eaterRepo repositories.EaterRepository
	smsClient sms.Client
	logger    *zap.Logger
}

func (s *eaterSvcImpl) ConfirmSMSCode(ctx context.Context, eaterID string, smsCode string) (*models.EaterProfile, error) {
	panic("unimplemented")
}

func (s *eaterSvcImpl) GetEaterProfile(ctx context.Context, eaterID string) (*models.EaterProfile, error) {
	panic("unimplemented")
}

func (s *eaterSvcImpl) SignupEater(ctx context.Context, phoneNumber string) (string, error) {
	phoneNumberExist := true

	eater, err := s.eaterRepo.GetEaterProfile(ctx, phoneNumber)
	if err != nil {
		phoneNumberExist = false
	}
	if phoneNumberExist {
		return s.handleExistingEater(ctx, eater.ID)
	}
	return s.handleNewEater(ctx, phoneNumber)
}

func (s *eaterSvcImpl) UpdateEaterProfile(ctx context.Context, eaterID string, name string, imageUrl string) (*models.EaterProfile, error) {
	panic("unimplemented")
}


