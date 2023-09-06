package eater

import (
	"context"
	"github.com/DevShuxat/eater-service/EATER-SERVICE/src/domain/eater/models"
	"github.com/DevShuxat/eater-service/EATER-SERVICE/src/domain/eater/repositories"
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

func NewRepository(db *gorm.DB) repositories.EaterRepository {
	return &eaterRepoImpl{
		db: db,
	}
}
func (r *eaterRepoImpl) WithTx(ctx context.Context, f func(r repositories.EaterRepository) error) error {
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		r := eaterRepoImpl{
			db: tx,
		}
		if err := f(&r); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
func (r *eaterRepoImpl) SaveEater(ctx context.Context, eater *models.Eater) error {
	return nil
}
func (r *eaterRepoImpl) UpdateEater(ctx context.Context, eater *models.Eater) error {
	return nil
}
func (r *eaterRepoImpl) DeleteEater(ctx context.Context, eaterID string) error {
	return nil
}
func (r *eaterRepoImpl) GetEater(ctx context.Context, eaterID string) (*models.Eater, error) {
	return nil, nil
}
func (r *eaterRepoImpl) GetEaterByPhoneNumber(ctx context.Context, phoneNumber string) (*models.Eater, error) {
	return nil, nil
}
func (r *eaterRepoImpl) SaveEaterSmsCode(ctx context.Context, smsCode *models.EaterSmsCode) error {
	return nil
}
func (r *eaterRepoImpl) GetEaterSmsCode(ctx context.Context, eaterID, code string) (*models.EaterSmsCode, error) {
	return nil, nil
}
func (r *eaterRepoImpl) SaveEaterProfile(ctx context.Context, profile *models.EaterProfile) error {
	result := r.db.WithContext(ctx).Table(tableEaterProfiles).Create(profile)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *eaterRepoImpl) UpdateEaterProfile(ctx context.Context, profile *models.EaterProfile) error {
	return nil
}
func (r *eaterRepoImpl) UpdateEaterProfilePhoneNumberConfirmed(ctx context.Context, eaterID string, confirmed bool) error {
	return nil
}
func (r *eaterRepoImpl) DeleteEaterProfile(ctx context.Context, eaterID string) error {
	return nil
}
func (r *eaterRepoImpl) GetEaterProfile(ctx context.Context, eaterID string) (*models.EaterProfile, error) {
	return nil, nil
}
