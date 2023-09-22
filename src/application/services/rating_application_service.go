package services

import (
	"context"
	RatingService "github.com/DevShuxat/eater-service/src/domain/rating/services"
)

type RatingApplicationService interface {
	RateDelivery(ctx context.Context, orderID, eaterID, comment string, rating int) (*dtos.DeliveryRatingResponse, error)
	UpdateDelivery(ctx context.Context, ratingID, comment string, rating int) (*dtos.DeliveryRatingResponse, error)
	ListDelivery(ctx context.Context, eaterID string) ([]*dtos.DeliveryRatingResponse, error)
	GetDeliveryRating(ctx context.Context, orderID string) ([]*dtos.DeliveryRatingResponse, error)

	RateRestaurant(ctx context.Context, EaterID, RestaurantID, comment string, rating int) (*dtos.RestaurantRatingResponse, error)
	UpdateRestaurantRating(ctx context.Context, ratingID, comment string, rating int) (*dtos.RestaurantRatingResponse, error)
	ListRestaurantRating(ctx context.Context, restaurantID string) ([]*dtos.RestaurantRatingResponse, error)
}

type ratingSvcImpl struct {
	ratingSvc RatingService.RatingService
} 




func NewRatingApplicationService {
	ratingSvcImpl := &ratingSvcImpl{
RatingApplicationService {
	return &ratingSvcImpl{
		ratingSvc: ratingSvcImpl
}