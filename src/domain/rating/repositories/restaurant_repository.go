package repositories

import (
	"context"

	"github.com/DevShuxat/eater-service/src/domain/rating/models"
)

type RestaurantRatingRepository interface {
	RateRestaurant(ctx context.Context, rating *models.DeliveryRating) error
	UpdateRestaurantRating(ctx context.Context, eaterID, RestaurantID, comment string, rating int) (models.RestaurantRating, error)
	ListRestaurantRating(ctx context.Context, eaterID string) error
}
