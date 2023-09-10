package repositories

import (
	"context"

	"github.com/DevShuxat/eater-service/src/domain/order/models"
)

type OrderRepository interface {
	WithTx(ctx context.Context, f func(r OrderRepository) error) error
	DeleteOrder(ctx context.Context, orderID string) error
	GetOrder(ctx context.Context, orderID string) error
	ListOrder(ctx context.Context, eaterID string) (models.Order, error)
	UpdateOrder(ctx context.Context, order *models.Order) error
	UpdateOrderStatus(ctx context.Context, order *models.Order) error
	UpdateOrderPaymentStatus(ctx context.Context, order *models.Order) error
	CreateOrder(ctx context.Context, order *models.Order) error
}