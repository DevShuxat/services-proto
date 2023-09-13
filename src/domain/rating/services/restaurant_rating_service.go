package services

import (
	"context"
	"errors"
	"time"

	"github.com/DevShuxat/eater-service/src/domain/rating/models"
	"github.com/DevShuxat/eater-service/src/domain/rating/repositories"
)

type ResturantRatingService interface {
	RateRestaurant(ctx context.Context, eaterID, RestaurantID, comment string, rating int) (models.RestaurantRating, error)
	UpdateRestaurantRating(ctx context.Context, eaterID, RestaurantID, comment string, rating int) (models.RestaurantRating, error)
	ListRestaurantRating(ctx context.Context, eaterID string) error
}
type rateRestaurantSvcImpl struct {
	ratingRepo repositories.RatingRepository
}

func NewRateRestaurantService(
	ratingRepo repositories.RatingRepository,
) ResturantRatingService {
	return &rateRestaurantSvcImpl{
		ratingRepo: ratingRepo,
	}
}

func (s *rateRestaurantSvcImpl) RateRestaurant(ctx context.Context, EaterID, RestaurantID, comment string, rating int) (models.RestaurantRating, error) {
	ratingData := &models.RestaurantRating{
		ID:        RestaurantID,
		EaterID:   EaterID,
		Comment:   comment,
		Rating:    rating,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	err := s.ratingRepo.WithTx(ctx, func(r repositories.RatingRepository) error {
		if err := r.RateRestaurant(ctx, ratingData); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return models.RestaurantRating{}, err
	}
	return &ratingData, nil
}

func (s *rateRestaurantSvcImpl) UpdateRestaurantRating(ctx context.Context, eaterID, RestaurantID, comment string, rating int) (models.RestaurantRating, error) {
	if RestaurantID == "" {
		return models.RestaurantRating{}, errors.New("Restaurant ID is required for updating")
	}

	existingRating, err := s.ratingRepo.RateRestaurant(ctx, eaterID string, RestaurantID string, comment string, rating int) (models.RestaurantRating, error)
	if err != nil {
		return models.RestaurantRating{}, err
	}

	existingRating.Rating = rating
	existingRating.Comment = comment
	existingRating.UpdatedAt = time.Now().UTC()

	err = s.ratingRepo.WithTx(ctx, func(r repositories.RatingRepository) error {
		if err := r.UpdateRestaurantRating(ctx, existingRating); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return models.RestaurantRating{}, err
	}

	return &existingRating, nil
}

func (s *rateRestaurantSvcImpl) ListRestaurantRating(ctx context.Context, eaterID string) ([]models.RestaurantRating, error) {
	ratings, err := s.ratingRepo.ListRestaurantRating(ctx, eaterID)
	if err != nil {
		return nil, err
	}
	return ratings, nil
}

