package eater

import (
	"context"
	"github.com/DevShuxat/eater-service/src/domain/eater/models"
	"github.com/DevShuxat/eater-service/src/domain/eater/repositories"
	"gorm.io/gorm"
)

const (
	tableEaters        = "eater.eaters"
	tableEaterSmsCodes = "eater.eater_sms_codes"
	tableEaterProfiles = "eater.eater_profiles"
)

type eaterRepoImpl struct {
	db *gorm.DB
}

func (r *eaterRepoImpl) DeleteEater(ctx context.Context, eaterID string) error {
	var eater models.Eater
	result := r.db.WithContext(ctx).Table(tableEaters).Delete(&eater, "id = ?", eaterID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *eaterRepoImpl) DeleteEaterProfile(ctx context.Context, eaterID string) error {
	var profile models.EaterProfile
	result := r.db.WithContext(ctx).Table(tableEaterProfiles).Delete(&profile, "id = ?", eaterID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (*eaterRepoImpl) GetEater(ctx context.Context, eaterID string) (*models.Eater, error) {
	panic("unimplemented")
}

func (*eaterRepoImpl) GetEaterByPhoneNumber(ctx context.Context, phoneNumber string) (*models.Eater, error) {
	panic("unimplemented")
}

func (*eaterRepoImpl) GetEaterProfile(ctx context.Context, eaterID string) (*models.EaterProfile, error) {
	panic("unimplemented")
}

func (*eaterRepoImpl) GetEaterSmsCode(ctx context.Context, eaterID string, code string) (*models.EaterSmsCode, error) {
	panic("unimplemented")
}

func (*eaterRepoImpl) SaveEater(ctx context.Context, eater *models.Eater) error {
	panic("unimplemented")
}

func (*eaterRepoImpl) SaveEaterProfile(ctx context.Context, profile *models.EaterProfile) error {
	panic("unimplemented")
}

func (*eaterRepoImpl) SaveEaterSmsCode(ctx context.Context, smsCode *models.EaterSmsCode) error {
	panic("unimplemented")
}

func (*eaterRepoImpl) UpdateEater(ctx context.Context, eater *models.Eater) error {
	panic("unimplemented")
}

func (*eaterRepoImpl) UpdateEaterProfile(ctx context.Context, profile *models.EaterProfile) error {
	panic("unimplemented")
}

func (*eaterRepoImpl) UpdateEaterProfilePhoneNumberConfirmed(ctx context.Context, eaterID string, confirmed bool) error {
	panic("unimplemented")
}

func NewRepository(db *gorm.DB) repositories.EaterRepository {
	return &eaterRepoImpl{
		db: db,
	}
}

func (r *eaterRepoImpl) WithTx(ctx context.Context, callback func(r repositories.EaterRepository) error) error {
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		r := eaterRepoImpl{
			db: tx,
		}
		if err := callback(&r); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
