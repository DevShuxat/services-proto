package rating

import (
	"time"

	pb "github.com/DevShuxat/eater-service/src/application/protos/eater"
	"github.com/DevShuxat/eater-service/src/domain/rating/models"
)

func ToRestaurantRatingPB(rrating *models.RestaurantRating) *pb.RestaurantRating {
	return &pb.RestaurantRating{
		Id:           rrating.ID,
		EaterId:      rrating.EaterID,
		RestaurantId: rrating.RestaurantID,
		Rating:       int32(rrating.Rating),
		Comment:      rrating.Comment,
		CreatedAt:    rrating.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    rrating.UpdatedAt.Format(time.RFC3339),
	}
}

func ToRestaurantRatingsPB(rrating []*models.RestaurantRating) []*pb.RestaurantRating {
	ratingsArr := make([]*pb.RestaurantRating, len(rrating))
	for i, rating := range rrating {
		ratingsArr[i] = ToRestaurantRatingPB(rating)
	}
	return ratingsArr
}

func ToDeliveryRatingPB(drating *models.DeliveryRating) *pb.DeliveryRating {
	return &pb.DeliveryRating{
		Id:        drating.ID,
		EaterId:   drating.EaterID,
		OrderId:   drating.OrderID,
		Rating:    int32(drating.Rating),
		Comment:   drating.Comment,
		CreatedAt: drating.CreatedAt.Format(time.RFC3339),
		UpdatedAt: drating.UpdatedAt.Format(time.RFC3339),
	}
}

func ToDeliveryRatingsPB(drating []*models.DeliveryRating) []*pb.DeliveryRating {
	ratingsArr := make([]*pb.DeliveryRating, len(drating))
	for i, rating := range drating {
		ratingsArr[i] = ToDeliveryRatingPB(rating)
	}
	return ratingsArr
}
