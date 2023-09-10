package services

import (
	"context"

	"github.com/DevShuxat/eater-service/src/domain/order/models"
	"github.com/DevShuxat/eater-service/src/domain/order/repositories"
	"go.uber.org/zap"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order models.Order) error
	UpdateOrder(ctx context.Context, order models.Order) error
	UpdateOrderStatus(ctx context.Context, order models.Order) error
	UpdateOrderPaymentStatus(ctx context.Context, order models.Order)
	GetOrder(ctx context.Context, orderID string) (*models.Order, error)
	ListOrders(ctx context.Context, eaterID string) ([]*models.Order, error)
	DeleteOrder(ctx context.Context, orderID string) error
}

type orderSvcImpl struct {
	orderRepo repositories.OrderRepository
	logger    *zap.Logger 
}

func NewOrderService(
	orderRepo repositories.OrderRepository,
	logger *zap.Logger,
) OrderService {
	return &orderSvcImpl{
		orderRepo: orderRepo,
		logger:    logger,
	}
}

func (s *orderSvcImpl) CreateOrder(ctx context.Context, cart string, Order string,Restaurant string ) (*models.Order, error) {

}

func (s *orderSvcImpl) UpdateOrder(ctx context.Context, order models.Order) error {
	return s.orderRepo.UpdateOrder(ctx, &order)
}

func (s *orderSvcImpl) UpdateOrderStatus(ctx context.Context, order models.Order) error {
	return s.orderRepo.UpdateOrderStatus(ctx, &order)
}

func (s *orderSvcImpl) UpdateOrderPaymentStatus(ctx context.Context, order models.Order) {
	s.orderRepo.UpdateOrderPaymentStatus(ctx, &order)
}
