package repositories

import (
	"context"

	"github.com/DevShuxat/eater-service/src/domain/eater/models"
)

type EaterRepository interface {
	WithTx(ctx context.Context, f func(r EaterRepository) error) error
	SaveEater(ctx context.Context, eater *models.Eater) error
	UpdateEater(ctx context.Context, eater *models.Eater) error
	DeleteEater(ctx context.Context, eaterID string) 
	GetEater(ctx context.Context, eaterID string) (*models.Eater, error)
	GetEaterByPhoneNumber(ctx context.Context, phoneNumber string) (*models.Eater, error)
	SaveEaterSmsCode(ctx context.Context, smsCode *models.EaterSmsCode) error
	GetEaterSmsCode(ctx context.Context, eaterID, code string) (*models.EaterSmsCode, error)
	SaveEaterProfile(ctx context.Context, profile *models.EaterProfile) error
	UpdateEaterProfile(ctx context.Context, profile *models.EaterProfile) error
	UpdateEaterProfilePhoneNumberConfirmed(ctx context.Context, eaterID string, confirmed bool) error
	DeleteEaterProfile(ctx context.Context, eaterID string) error
	GetEaterProfile(ctx context.Context, eaterID string) (*models.EaterProfile, error)
}


