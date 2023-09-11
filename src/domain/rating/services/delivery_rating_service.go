package services

import (
	"context"
	"errors"
	"time"

	"github.com/DevShuxat/eater-service/src/domain/rating/models"
	"github.com/DevShuxat/eater-service/src/domain/rating/repositories"
)

type DeliveryRatingService interface {
	RateDelivery(ctx context.Context, orderID string, rating int, comment string) (string, error)
	UpdateDelivery(ctx context.Context, orderID string) ([]*models.DeliveryRating, error)
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
		ratingRepo: ratingRepo,
	}
}
func (s *deliverySvcImpl) RateDelivery(ctx context.Context, orderID string, rating int, comment string) error {
	existingRating, err := s.deliveryRepo.GetDeliveryRating(ctx, orderID)
	if err != nil {
		return err
	}

	if existingRating != nil {
		return errors.New("eater can rate a delivered order only once")
	}

	deliveryRating := models.DeliveryRating{
		OrderID:   orderID,
		Rating:    rating,
		Comment:   comment,
		CreatedAt: time.Now().UTC(),
	}

	if err := s.deliveryRepo.Save(ctx, &deliveryRating); err != nil {
		return err
	}

	return nil
}

func (s *deliverySvcImpl) UpdateDeliveryRating(ctx context.Context, orderID string, rating int, comment string) error {
	existingRating, err := s.deliveryRepo.GetDeliveryRating(ctx, orderID)
	if err != nil {
		return err
	}

	if existingRating == nil {
		return errors.New("rating not found")
	}

	existingRating.Rating = rating
	existingRating.Comment = comment
	existingRating.UpdatedAt = time.Now().UTC()

	if err := s.deliveryRepo.Update(ctx, existingRating); err != nil {
		return err
	}

	return nil
}

func (s *deliverySvcImpl) GetDeliveryRating(ctx context.Context, orderID string) ([]*DeliveryRating, error) {
	rating, err := s.deliveryRepo.GetDeliveryRating(ctx, orderID)
	if err != nil {
		return nil, err
	}

	return rating, nil
}
