package repositories

import (
	"context"

	"github.com/DevShuxat/eater-service/src/domain/rating/models"
)

type RatingRepository interface {
	RateDelivery(ctx context.Context, rating *models.DeliveryRating) error
	UpdateDelivery(ctx context.Context, rating *models.DeliveryRating) error
	ListDelivery(ctx context.Context, eaterID string) ([]*models.DeliveryRating, error)
	GetDeliveryRating(ctx context.Context, OrderID string) ([]*models.DeliveryRating, error)

	GetRestaurantRating(ctx context.Context, eaterID string) ([]*models.RestaurantRating, error)
	RateRestaurant(ctx context.Context, rating *models.RestaurantRating) error
	UpdateRestaurantRating(ctx context.Context, rating *models.RestaurantRating) error
	ListRestaurantRating(ctx context.Context, restaurantID string) ([]*models.RestaurantRating, error)
}
