package services

import (
	"context"
	"time"

	"github.com/DevShuxat/eater-service/src/domain/rating/models"
	"github.com/DevShuxat/eater-service/src/domain/rating/repositories"
)

type DeliveryRatingService interface {
	RateDelivery(ctx context.Context, orderID, eaterID, comment string, rating int, deliveryRating *models.DeliveryRating) error
	UpdateDelivery(ctx context.Context, orderID, eaterID, comment string, rating int, deliveryRating *models.DeliveryRating) error
	ListDelivery(ctx context.Context, eaterID string) ([]*models.DeliveryRating, error)
	GetDeliveryRating(ctx context.Context, orderID string) ([]*models.DeliveryRating, error)
}

type deliverySvcImpl struct {
	deliveryRepo repositories.RatingRepository
}

func NewDeliveryRatingService(
	ratingRepo repositories.RatingRepository,
) DeliveryRatingService {
	return &deliverySvcImpl{
		deliveryRepo: ratingRepo,
	}
}

func (s *deliverySvcImpl) RateDelivery(ctx context.Context, orderID, eaterID, comment string, rating int, deliveryRating *models.DeliveryRating) error {
	order := &models.DeliveryRating{
		OrderID:   orderID,
		EaterID:   eaterID,
		Comment:   comment,
		Rating:    rating,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	err := s.deliveryRepo.RateDelivery(ctx, order)
	if err != nil {
		return err
	}
	return nil
}

func (s *deliverySvcImpl) UpdateDelivery(ctx context.Context, orderID, eaterID, comment string, rating int, deliveryRating *models.DeliveryRating) error {
	order := &models.DeliveryRating{
		OrderID:   orderID,
		EaterID:   eaterID,
		Comment:   comment,
		Rating:    rating,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	err := s.deliveryRepo.UpdateDelivery(ctx, order)
	if err != nil {
		return err
	}
	return nil
}

func (s *deliverySvcImpl) GetDeliveryRating(ctx context.Context, orderID string) ([]*models.DeliveryRating, error) {
	rating, err := s.deliveryRepo.GetDeliveryRating(ctx, orderID)
	if err != nil {
		return nil, err
	}

	return rating, nil
}

func (s *deliverySvcImpl)	ListDelivery(ctx context.Context, eaterID string) ([]*models.DeliveryRating, error) {
	rating, err := s.deliveryRepo.ListDelivery(ctx, eaterID)
	if err != nil {
		return nil, err
	}
	return rating, nil
}
