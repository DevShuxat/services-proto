package rating

import (
	"context"
	"time"

	"github.com/DevShuxat/eater-service/src/domain/rating/models"
	"github.com/DevShuxat/eater-service/src/domain/rating/repositories"
	"gorm.io/gorm"
)


const (
	tableDelivery = "delivery_rating.delivery_rating"
)

type deliveryRateRepoImpl struct {
	db *gorm.DB
}
func NewEaterRepository(db *gorm.DB) repositories.DeliveryRateRepository {
	return &deliveryRateRepoImpl{
		db: db,
	}
}

func (r *deliveryRateRepoImpl) RateDelivery(ctx context.Context, orderID string, rating int, comment string) (models.DeliveryRating, error) {
  newDelRate := &models.DeliveryRating{
		OrderID: orderID,
		Rating:       rating,
		Comment:      comment,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}

	result := r.db.WithContext(ctx).Table(tableDelivery).Create(newDelRate)
	if result.Error != nil {
		return nil, result.Error
	}

return newDelRate, nil
}


func (r *deliveryRateRepoImpl) UpdateDelivery(ctx context.Context, rating *models.DeliveryRating) error {
  err := r.db.WithContext(ctx).Model(&models.DeliveryRating{}).Where("id = ?", rating.ID).Updates(&rating).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *deliveryRateRepoImpl) GetDeliveryRating(ctx context.Context, orderID string) (models.DeliveryRating, error) {
 var rating models.DeliveryRating
 result := r.db.WithContext(ctx).Table(tableDelivery).First(&rating, "id = ?", orderID)
if result.Error != nil {
		return nil, result.Error
	}
	return &rating, nil
}
