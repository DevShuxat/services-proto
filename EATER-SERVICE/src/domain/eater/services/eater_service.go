package services

import (
	"context"
	"crypto"
	"fmt"
	"time"

	"github.com/DevShuxat/eater-service/EATER-SERVICE/src/domain/eater/models"
	"github.com/DevShuxat/eater-service/EATER-SERVICE/src/domain/eater/repositories"
	"github.com/DevShuxat/eater-service/EATER-SERVICE/src/infrastructure/rand"
	"github.com/DevShuxat/eater-service/EATER-SERVICE/src/infrastructure/sms"
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

func (s *eaterSvcImpl) handleNewEater(ctx context.Context, phoneNumber string) (string, error) {
	var (
		eaterID    = rand.UUID()
		eaterName  = fmt.Sprintf("eater-%s", rand.NumericString(5))
		salt       = crypto.GetEaterSalt()
		saltedPass = crypto.Combine(salt, phoneNumber)
		passHash   = crypto.HashPassword(saltedPass)
		now        = time.Now().UTC()
	)
	eater := models.Eater{
		ID:           eaterID,
		PhoneNumber:  phoneNumber,
		PasswordHash: passHash,
		PasswordSalt: salt,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	eaterProfile := models.EaterProfile{
		EaterID:     eaterID,
		PhoneNumber: phoneNumber,
		Name:        eaterName,
		ImageUrl:    "",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	smsCode := models.EaterSmsCode{
		EaterID:   eaterID,
		Code:      rand.NumericString(5),
		ExpiresIn: 300, //5minutes
		CreatedAt: now,
	}
	err := s.eaterRepo.WithTx(ctx, func(r repositories.EaterRepository) error {
		if err := r.SaveEater(ctx, &eater); err != nil {
			return err
		}
		if err := r.SaveEaterProfile(ctx, &eaterProfile); err != nil {
			return err
		}
		if err := r.SaveEaterSmsCode(ctx, &smsCode); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return "", err

	}

	smsMsg := fmt.Sprintf("food.uz Code: %s", smsCode.Code)
	if err := s.smsClient.SendMessage(ctx, phoneNumber, smsMsg); err != nil {
		return "", err
	}
	return eaterID, nil
}

func (s *eaterSvcImpl) handleExistingEater(ctx context.Context, eaterID string) {
	eater, err := s.eaterRepo.GetEater(ctx, eaterID)
	if err != nil {
		return "", err
	}

	smsCode := models.EaterSmsCode{
		EaterID:   eaterID,
		Code:      rand.NumericString(5),
		ExpiresIn: 300, //5minut
		CreatedAt: time.Now(),
	}

	if err := s.eaterRepo.SaveEaterSmsCode(ctx, &smsCode); err != nil {
		return "", err
	}

	smsMsg := fmt.Sprintf("food.uz Code: %s", smsCode.Code)
	if err := s.smsClient.SendMessage(ctx, eater.PhoneNumber, smsMsg); err != nil {
		return "", err
	}

	return eaterID, nil;
}
}
func (s *eaterSvcImpl) ConfirmSMSCode(ctx context.Context, eaterID, code string) (*models.Eater, error) {
	smsCode, err := s.eaterRepo.GetEaterSmsCode(ctx, eaterID, code)
	if err != nil {
		return nil, err
	}
	if smsCode.IsExpired() {
		return nil, errors.New("code is expired")
	}
}
