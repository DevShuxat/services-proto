package services

import (
	"context"
	"time"

	"github.com/DevShuxat/eater-service/src/domain/rating/models"
	"github.com/DevShuxat/eater-service/src/domain/rating/repositories"
	"github.com/DevShuxat/eater-service/src/infrastructure/rand"
)

type RatingService interface {
	RateDelivery(ctx context.Context, orderID, eaterID, comment string, rating int) (*models.DeliveryRating, error)
	UpdateDelivery(ctx context.Context, ratingID, comment string, rating int) (*models.DeliveryRating, error)
	ListDelivery(ctx context.Context, eaterID string) ([]*models.DeliveryRating, error)
	GetDeliveryRating(ctx context.Context, EaterID string) ([]*models.DeliveryRating, error)

	RateRestaurant(ctx context.Context, EaterID, RestaurantID, comment string, rating int) (*models.RestaurantRating, error)
	UpdateRestaurantRating(ctx context.Context, ratingID, comment string, rating int) (*models.RestaurantRating, error)
	ListRestaurantRating(ctx context.Context, eaterID string) ([]*models.RestaurantRating, error)
}

type ratingSvcImp struct {
	ratingRepo repositories.RatingRepository
}

func NewDeliveryRatingService(
	ratingRepo repositories.RatingRepository,
) RatingService {
	return &ratingSvcImp{
		ratingRepo: ratingRepo,
	}
}

func (s *ratingSvcImp) RateDelivery(ctx context.Context, orderID, eaterID, comment string, rating int) (*models.DeliveryRating, error) {
	ratingD := models.DeliveryRating{
		ID:        rand.UUID(),
		OrderID:   orderID,
		EaterID:   eaterID,
		Comment:   comment,
		Rating:    rating,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	err := s.ratingRepo.RateDelivery(ctx, &ratingD)
	if err != nil {
		return nil, err
	}
	return &ratingD, nil
}

func (s *ratingSvcImp) UpdateDelivery(ctx context.Context, ratingID, comment string, rating int) (*models.DeliveryRating, error) {
	ratingD := models.DeliveryRating{
		ID:        rand.UUID(),
		Comment:   comment,
		Rating:    rating,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	err := s.ratingRepo.UpdateDelivery(ctx, &ratingD)
	if err != nil {
		return nil, err
	}
	return &ratingD,nil
}

func (s *ratingSvcImp) GetDeliveryRating(ctx context.Context, EaterID string) ([]*models.DeliveryRating, error) {
	rating, err := s.ratingRepo.GetDeliveryRating(ctx, EaterID)
	if err != nil {
		return nil, err
	}

	return rating, nil
}

func (s *ratingSvcImp) ListDelivery(ctx context.Context, eaterID string) ([]*models.DeliveryRating, error) {
	rating, err := s.ratingRepo.ListDelivery(ctx, eaterID)
	if err != nil {
		return nil, err
	}
	return rating, nil
}

//########restaurant########################

func (s *ratingSvcImp) GetRestaurantRating(ctx context.Context, eaterID string) ([]*models.RestaurantRating, error) {
	rating, err := s.ratingRepo.GetRestaurantRating(ctx, eaterID)
	if err != nil {
		return nil, err
	}

	return rating, nil
}
func (s *ratingSvcImp) RateRestaurant(ctx context.Context, EaterID, RestaurantID, comment string, rating int) (*models.RestaurantRating, error) {
	ratingR := models.RestaurantRating{
		RestaurantID: RestaurantID,
		EaterID:      EaterID,
		Comment:      comment,
		Rating:       rating,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}

	err := s.ratingRepo.RateRestaurant(ctx, &ratingR)
	if err != nil {
		return nil, err
	}
	return &ratingR, nil
}

func (s *ratingSvcImp) UpdateRestaurantRating(ctx context.Context, ratingID, comment string, rating int) (*models.RestaurantRating, error) {
	ratingR := models.RestaurantRating{
		ID:        rand.UUID(),
		Comment:   comment,
		Rating:    rating,
		UpdatedAt: time.Now().UTC(),
	}

	err := s.ratingRepo.UpdateRestaurantRating(ctx, &ratingR)
	if err != nil {
		return nil, err
	}
	return &ratingR, err
}

func (s *ratingSvcImp) ListRestaurantRating(ctx context.Context, eaterID string) ([]*models.RestaurantRating, error) {
	ratings, err := s.ratingRepo.ListRestaurantRating(ctx, eaterID)
	if err != nil {
		return nil, err
	}
	return ratings, nil
}
