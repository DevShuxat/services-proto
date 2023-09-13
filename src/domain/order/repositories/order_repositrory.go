package repositories

import (
	"context"
	"time"

	"github.com/DevShuxat/eater-service/src/domain/order/models"
)

type OrderRepository interface {
	SaveOrder(ctx context.Context, order *models.Order) error
	UpdateOrder(ctx context.Context, order *models.Order) error
	UpdateOrderStatus(ctx context.Context, orderID, newStatus string, time time.Time) error
	UpdateOrderPaymentStatus(ctx context.Context, orderID, newPaymentStatus string, time time.Time) error
	GetOrder(ctx context.Context, orderID string) (*models.Order, error)
	ListOrder(ctx context.Context, eaterID string, sort string, page, pageSize int) ([]*models.Order, error)
	DeleteOrder(ctx context.Context, orderID string) error
}