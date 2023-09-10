package repositories

import (
	"context"

	"github.com/DevShuxat/eater-service/src/domain/rating/models"
)

type RestaurantRatingRepository interface {
	WithTx(ctx context.Context, f func(r RestaurantRatingRepository) error) error
	RateRestaurant(ctx context.Context, RestaurantID string, rating int, comment string) (models.RestaurantRating, error)
	UpdateRestaurantRating(ctx context.Context, restaurantRating *models.RestaurantRating) error
	ListRestaurantRating(ctx context.Context, eaterID string) error
}