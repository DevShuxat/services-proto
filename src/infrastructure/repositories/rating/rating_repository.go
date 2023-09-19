package rating

import (
	"context"

	"github.com/DevShuxat/eater-service/src/domain/rating/models"
	"github.com/DevShuxat/eater-service/src/domain/rating/repositories"
	"gorm.io/gorm"
)


const (
	tableDelivery = "eater.delivery_ratings"
	tableRestaurantRating = "eater.restaurant_ratings"
)

type deliveryRateRepoImpl struct {
	db *gorm.DB
}
func NewEaterRepository(db *gorm.DB) repositories.RatingRepository {
	return &deliveryRateRepoImpl{
		db: db,
	}
}

func (r *deliveryRateRepoImpl)  RateDelivery(ctx context.Context, orderID, eaterID,comment string, rating int, *models.DeliveryRating)  error {
	result := r.db.WithContext(ctx).Table(tableDelivery).Create(&orderID)
	if result.Error != nil {
		return err
	}
	return nil

}


func (r *deliveryRateRepoImpl) UpdateDelivery(ctx context.Context, orderID string) ([]*models.DeliveryRating, error) {
  err := r.db.WithContext(ctx).Model(&models.DeliveryRating{}).Where("id = ?", orderID).Updates(&orderID)
	if err != nil {
		return nil, err
	}
	return orderID, nil
}

func (r *deliveryRateRepoImpl) GetDeliveryRating(ctx context.Context, orderID string) ([]*models.DeliveryRating, error) {
 var rating []*models.DeliveryRating
 result := r.db.WithContext(ctx).Table(tableDelivery).First(&rating, "id = ?", orderID)
if result.Error != nil {
		return nil, result.Error
	}
	return rating, nil
}

func (r *deliveryRateRepoImpl) ListDelivery(ctx context.Context, eaterID string) ([]*models.DeliveryRating, error) {
	var rating []*models.DeliveryRating
	result := r.db.WithContext(ctx).Table(tableDelivery).Where(&rating, "id = ?", eaterID)
	if result.Error != nil {
		return rating, result.Error
	}
	return rating, nil
}


func (r *deliveryRateRepoImpl) ListRestaurantRating(ctx context.Context, eaterID string) error {
	var rating *models.DeliveryRating
	result := r.db.WithContext(ctx).Table(tableDelivery).Where(&rating, "id = ?", eaterID)
	if result.Error != nil {
		return rating, result.Error
	}
	return rating, nil
}
