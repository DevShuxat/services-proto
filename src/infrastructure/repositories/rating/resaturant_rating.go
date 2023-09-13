package rating

import (
	"context"
	"time"

	"github.com/DevShuxat/eater-service/src/domain/rating/models"
	"github.com/DevShuxat/eater-service/src/domain/rating/repositories"
	"gorm.io/gorm"
)

const (
	tableRating = "restaurant_rating.restaurant_rating"
)

type ratingRepoImpl struct {
	db *gorm.DB
}

func (r *ratingRepoImpl) ListRestaurantRating(ctx context.Context, eaterID string) error {
var restaurant models.RestaurantRating
	result := r.db.WithContext(ctx).Table(tableRating).First(&restaurant, "eater_id = ?", eaterID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &restaurant,
}

