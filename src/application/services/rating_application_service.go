package services

import (
	"context"
	"errors"

	dtos "github.com/DevShuxat/eater-service/src/application/dtos/rating"
	pb "github.com/DevShuxat/eater-service/src/application/protos/eater"
	RatingService "github.com/DevShuxat/eater-service/src/domain/wallet/services"
)

type RatingApplicationService interface {
	RateDelivery(ctx context.Context, req *pb.RateDeliveryRequest) (*pb.RateDeliveryResponse, error)
	UpdateDelivery(ctx context.Context, req *pb.UpdateDeliveryRatingRequest) (*pb.UpdateDeliveryRatingResponse, error)
	ListDelivery(ctx context.Context, req *pb.ListDeliveryRatingByEaterRequest) (*pb.ListDeliveryRatingByEaterResponse, error)
	GetDeliveryRating(ctx context.Context, req *pb.GetDeliveryRatingByOrderRequest) (*pb.GetDeliveryRatingByOrderResponse, error)

	RateRestaurant(ctx context.Context, req *pb.RateRestaurantRequest) (*pb.RateRestaurantResponse, error)
	UpdateRestaurantRating(ctx context.Context, req *pb.UpdateRestaurantRatingRequest) (*pb.UpdateRestaurantRatingResponse, error)
	ListRestaurantRatingByEater(ctx context.Context, req *pb.ListRestaurantRatingByEaterRequest) (*pb.ListRestaurantRatingByEaterResponse, error)
}

type ratingsSvc struct {
	ratingSvc RatingService.RatingService
}

func NewRatingApplicationService(
	ratingSvc RatingService.RatingService,
) RatingApplicationService {
	return &ratingsSvc{
		ratingSvc: ratingSvc,
	}
}

// DeliveryRating start #####################################
// func (s *ratingsSvc) RateDelivery(ctx context.Context, req *pb.RateDeliveryRequest) (*pb.RateDeliveryResponse, error) {
//     if req.GetRating() == 0 {
//         return nil, errors.New("invalid rating request")
//     }
//     if req.GetComment() == "" {
//         return nil, errors.New("comment required")
//     }
//     if req.GetOrderId() == "" {
//         return nil, errors.New("order ID required")
//     }
//     if req.GetEaterId() == "" {
//         return nil, errors.New("eater ID required")
//     }

//     rating, err := s.ratingSvc.RateDelivery(ctx, req.GetOrderId(), req.GetEaterId(), req.GetComment(), int(req.GetRating()))
//     if err != nil {
//         return nil, err
//     }

//     // Use ToDeliveryRatingsPB to convert the rating to a slice of DeliveryRating
//     ratingPB := dtos.ToDeliveryRatingsPB(rating)

//     return &pb.RateDeliveryResponse{
//         Rating: ratingPB,
//     }, nil
// }

func (s *ratingsSvc) RateDelivery(ctx context.Context, req *pb.RateDeliveryRequest) (*pb.RateDeliveryResponse, error) {
	if req.GetRating() == 0 {
		return nil, errors.New("invalid rating request")
	}
	if req.GetComment() == "" {
		return nil, errors.New("comment required")
	}
	if req.GetOrderId() == "" {
		return nil, errors.New("order ID required")
	}
	if req.GetEaterId() == "" {
		return nil, errors.New("eater ID required")
	}

	rating, err := s.ratingSvc.RateDelivery(ctx, req.GetOrderId(), req.GetEaterId(), req.GetComment(), int(req.GetRating()))
	if err != nil {
		return nil, err
	}
	ratingPB := dtos.ToDeliveryRatingPB(rating) // Assuming this function handles a single rating

	return &pb.RateDeliveryResponse{
		Rating: ratingPB,
	}, nil
}


func (s *ratingsSvc) UpdateDelivery(ctx context.Context, req *pb.UpdateDeliveryRatingRequest) (*pb.UpdateDeliveryRatingResponse, error) {
	if req.RatingId == "" {
		return nil, errors.New("id required")
	}
	if req.Rating == 0 {
		return nil, errors.New("invalid rating request")
	}
	if req.Comment == "" {
		return nil, errors.New("comment required")
	}

	rating, err := s.ratingSvc.UpdateDelivery(ctx, req.RatingId, req.Comment, int(req.Rating))
	if err != nil {
		return nil, err
	}

	return &pb.UpdateDeliveryRatingResponse{
		Rating: dtos.ToDeliveryRatingPB(rating),
	}, nil
}
func (s *ratingsSvc) ListDelivery(ctx context.Context, req *pb.ListDeliveryRatingByEaterRequest) (*pb.ListDeliveryRatingByEaterResponse, error) {
	if req.EaterId == "" {
		return nil, errors.New("no more delivery")
	}

	rating, err := s.ratingSvc.ListDelivery(ctx, req.EaterId)
	if err != nil {
		return nil, err
	}

	return &pb.ListDeliveryRatingByEaterResponse{
		Ratings: dtos.ToDeliveryRatingsPB(rating),
	}, nil
}
func (s *ratingsSvc) GetDeliveryRating(ctx context.Context, req *pb.GetDeliveryRatingByOrderRequest) (*pb.GetDeliveryRatingByOrderResponse, error) {
	if req.GetOrderId() == "" {
		return nil, errors.New("order id is invalid")
	}

	rating, err := s.ratingSvc.GetDeliveryRating(ctx, req.GetOrderId())
	if err != nil {
		return nil, err
	}
	return &pb.GetDeliveryRatingByOrderResponse{
		Rating: dtos.ToDeliveryRatingPB(rating),
	}, nil
}

//end DeliveryRatin #########################################

func (s *ratingsSvc) RateRestaurant(ctx context.Context, req *pb.RateRestaurantRequest) (*pb.RateRestaurantResponse, error) {
	if req.Rating < 1 || req.Rating > 5 {
		return nil, errors.New("rating beetwen 1 and 5")
	}
	if req.EaterId == "" {
		return nil, errors.New("more comment")
	}
	if req.RestaurantId == "" {
		return nil, errors.New("restaurant id is invalid")
	}

	rating, err := s.ratingSvc.RateRestaurant(ctx, req.EaterId, req.RestaurantId, req.Comment, int(req.Rating))
	if err != nil {
		return nil, err
	}

	return &pb.RateRestaurantResponse{
		Rating: dtos.ToRestaurantRatingPB(rating),
	}, nil
}

func (s *ratingsSvc) UpdateRestaurantRating(ctx context.Context, req *pb.UpdateRestaurantRatingRequest) (*pb.UpdateRestaurantRatingResponse, error) {
	if req.RatingId == "" {
		return nil, errors.New("ratin id is required")
	}
	if req.Rating == 0 {
		return nil, errors.New("rating is required")
	}
	rating, err := s.ratingSvc.UpdateRestaurantRating(ctx, req.RatingId, req.Comment, int(req.Rating))
	if err != nil {
		return nil, err
	}
	return &pb.UpdateRestaurantRatingResponse{
		Rating: dtos.ToRestaurantRatingPB(rating),
	}, nil
}

func (s *ratingsSvc) ListRestaurantRatingByEater(ctx context.Context, req *pb.ListRestaurantRatingByEaterRequest) (*pb.ListRestaurantRatingByEaterResponse, error) {
	if req.EaterId == "" {
		return nil, errors.New("eater id is required")
	}
	ratings, err := s.ratingSvc.ListRestaurantRating(ctx, req.EaterId)
	if err != nil {
		return nil, err
	}

	ratingsPB := dtos.ToRestaurantRatingsPB(ratings) // Use ToRestaurantRatingsPB to convert the slice

	return &pb.ListRestaurantRatingByEaterResponse{
		Ratings: ratingsPB,
	}, nil
}

