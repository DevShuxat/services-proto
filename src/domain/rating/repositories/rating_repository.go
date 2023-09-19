package repositories

import (
	"context"

	"github.com/DevShuxat/eater-service/src/domain/rating/models"
)

type RatingRepository interface {
	RateDelivery(ctx context.Context, order *models.DeliveryRating) error
	UpdateDelivery(ctx context.Context, orderID *models.DeliveryRating) error
	ListDelivery(ctx context.Context, eaterID string) ([]*models.DeliveryRating, error)
	GetDeliveryRating(ctx context.Context, orderID string) ([]*models.DeliveryRating, error)
}
