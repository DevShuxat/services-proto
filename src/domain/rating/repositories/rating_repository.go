package repositories

import (
	"context"

	"github.com/DevShuxat/eater-service/src/domain/rating/models"
)

type RatingRepository interface {
	RateDelivery(ctx context.Context, orderID string, rating int, comment string) (string, error)
	UpdateDelivery(ctx context.Context, orderID string) ([]*models.DeliveryRating, error)
	ListDelivery(ctx context.Context, eaterID string) ([]*models.DeliveryRating, error)
	GetDeliveryRating(ctx context.Context, orderID string) ([]*models.DeliveryRating, error)
	RateRestaurant(ctx context.Context, eaterID, RestaurantID, comment string, rating int) (models.RestaurantRating, error)
	UpdateRestaurantRating(ctx context.Context, eaterID, RestaurantID, comment string, rating int) (models.RestaurantRating, error)
	ListRestaurantRating(ctx context.Context, eaterID string) error
}
