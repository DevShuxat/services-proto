package rating

import (
	"context"
	"time"

	"github.com/DevShuxat/eater-service/src/domain/rating/models"
	"github.com/DevShuxat/eater-service/src/domain/rating/repositories"
	"gorm.io/gorm"
)

const (
	tableRating = "restaurant_rating.restaurant_rating"
)

type ratingRepoImpl struct {
	db *gorm.DB
}

// ListRestaurantRating implements repositories.RestaurantRatingRepository.
func (r *ratingRepoImpl) ListRestaurantRating(ctx context.Context, eaterID string) error {
	var restaurant models.RestaurantRating
	result := r.db.WithContext(ctx).Table(tableRating).First(&restaurant, "eater_id = ?", eaterID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &restaurant,
}

// RateRestaurant implements repositories.RestaurantRatingRepository.
func (r *ratingRepoImpl) RateRestaurant(ctx context.Context, RestaurantID string, rating int, comment string) (*models.RestaurantRating, error) {
	newRating := &models.RestaurantRating{
		RestaurantID: RestaurantID,
		Rating:       rating,
		Comment:      comment,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}

	result := r.db.WithContext(ctx).Table(tableRating).Create(newRating)
	if result.Error != nil {
		return nil, result.Error
	}

	return newRating, nil
}

// UpdateRestaurantRating implements repositories.RestaurantRatingRepository.
func (r *ratingRepoImpl) UpdateRestaurantRating(ctx context.Context, restaurantRating *models.RestaurantRating) error {
	panic("unimplemented")
}

func (r *ratingRepoImpl) WithTx(ctx context.Context, callback func(r repositories.RestaurantRatingRepository) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		r := ratingRepoImpl{
			db: tx,
		}
		if err := callback(&r); err != nil {
			return err
		}
		return nil
	})
}
