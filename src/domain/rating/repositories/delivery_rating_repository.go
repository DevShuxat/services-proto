package repositories

import (
	"context"

	"github.com/DevShuxat/eater-service/src/domain/rating/models"
)

type DeliveryRatingRepository interface {
	WithTx(ctx context.Context, f func(r *DeliveryRatingRepository) error) error
	RateDelivery(ctx context.Context, orderID string, rating int, comment string) (models.DeliveryRating, error)
	Update(ctx context.Context, deliveryRating *models.DeliveryRating) error
	GetDeliveryRating(ctx context.Context, orderID string) (models.DeliveryRating, error)
}