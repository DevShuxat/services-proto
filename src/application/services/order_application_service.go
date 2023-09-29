package services

import (
	"context"
	"errors"
	"time"
	OrderService "github.com/DevShuxat/eater-service/src/domain/rating/services"
	"github.com/DevShuxat/eater-service/src/domain/order/models"
)

type OrderRepository interface {
	SaveOrder(ctx context.Context, req ) (*dtos.SaveOrderRespons, error)
	UpdateOrder(ctx context.Context, order *models.Order) (*dtos.UpdateOrderResponse, error)
	UpdateOrderStatus(ctx context.Context, orderID, newStatus string, time time.Time) (*dtos.UpdateOrderStatusResponse, error)
	UpdateOrderPaymentStatus(ctx context.Context, orderID, newPaymentStatus string, time time.Time) (*dtos.UpdateOrderPaymentStatsu,error)
	GetOrder(ctx context.Context, orderID string) (*dtos.GetOrderResponse, error)
	ListOrderByEater(ctx context.Context, eaterID string, sort string, page, pageSize int) ([]*dtos.ListOrderByEaterResponse, error)
	DeleteOrder(ctx context.Context, orderID string) (*dtos.DeleteOrderResponse, error)
}


type orderSvcImpl struct {
	orderSvc orderSvcImpl.OrderService
}

func NewOrderApplicationService (
	orderSvc OrderService.OrderService,
	) OrderApplicationService {
	return &orderSvcImpl {
	orderSvc: orderSvc,
	}
}


func (s *orderSvcImpl) SaveOrder(ctx context.Context, eaterID string, cart *models.Cart, address *models.AddressInfo) (*dtos.SaveOrderRespons, error) {
	if eaterID == "" {
		return nil, errors.New("invalid eater ID provided")
	}
	if models.Cart == "" {
		return nil, errors.New("invalid cart provided")
	}
}