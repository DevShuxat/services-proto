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

type rateRepoImpl struct {
	db *gorm.DB
}
func NewRatingRepository(db *gorm.DB) repositories.RatingRepository {
	return &rateRepoImpl{
		db: db,
	}
}

///start delivery################################################

func (r *rateRepoImpl)  RateDelivery(ctx context.Context, rating *models.DeliveryRating) error {
	result := r.db.WithContext(ctx).Table(tableDelivery).Create(&rating)
	if result.Error != nil {
		return nil
	}
	return nil

}

func (r *rateRepoImpl) UpdateDelivery(ctx context.Context, rating *models.DeliveryRating) error {
  result := r.db.WithContext(ctx).Table(tableDelivery).Where("id = ?", rating).Save(&rating)
	if result.Error != nil {
		return nil
	}
	return  nil
}

func (r *rateRepoImpl) GetDeliveryRating(ctx context.Context, orderID string) ([]*models.DeliveryRating, error) {
 var rating []*models.DeliveryRating
 result := r.db.WithContext(ctx).Table(tableDelivery).First(&rating, "id = ?", orderID)
if result.Error != nil {
		return nil, result.Error
	}
	return rating, nil
}

func (r *rateRepoImpl) ListDelivery(ctx context.Context, eaterID string) ([]*models.DeliveryRating, error) {
	var rating []*models.DeliveryRating
	result := r.db.WithContext(ctx).Table(tableDelivery).Where(&rating, "id = ?", eaterID)
	if result.Error != nil {
		return rating, result.Error
	}
	return rating, nil
}
///end delivery rate################################


///restaurant################################
func (r *rateRepoImpl) GetRestaurantRating(ctx context.Context, eaterID string) ([]*models.RestaurantRating, error) {
 var rating []*models.RestaurantRating
 result := r.db.WithContext(ctx).Table(tableRestaurantRating).First(&rating, "id = ?", eaterID)
if result.Error != nil {
		return nil, result.Error
	}
	return rating, nil
}


func (r *rateRepoImpl) RateRestaurant(ctx context.Context, rating *models.RestaurantRating) error {
	result := r.db.WithContext(ctx).Table(tableRestaurantRating).Create(&rating)
	if result.Error != nil {
		return nil
	}
	return nil

}

func (r *rateRepoImpl) UpdateRestaurantRating(ctx context.Context, rating *models.RestaurantRating) error {
  result := r.db.WithContext(ctx).Table(tableRestaurantRating).Where("id = ?", rating).Save(&rating)
	if result.Error != nil {
		return nil
	}
	return  nil
}

func (r *rateRepoImpl) ListRestaurantRating(ctx context.Context, restaurantID string) ([]*models.RestaurantRating, error) {
	var ratings []*models.RestaurantRating
	result := r.db.WithContext(ctx).Table(tableRestaurantRating).Where("id = ?", restaurantID).Find(ratings)
	if result.Error != nil {
		return nil, result.Error
	}
	return ratings, nil
}
