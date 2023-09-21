package services

import (
	"context"
	"time"

	"github.com/DevShuxat/eater-service/src/domain/order/models"
	"github.com/DevShuxat/eater-service/src/domain/order/repositories"
	"github.com/DevShuxat/eater-service/src/infrastructure/rand"
	"go.uber.org/zap"
)

type OrderService interface {
	SaveOrder(ctx context.Context, eaterID string, cart *models.Cart, address *models.AddressInfo) error
	UpdateOrder(ctx context.Context, order *models.Order) error
	UpdateOrderStatus(ctx context.Context, orderID, newStatus string, time time.Time) error
	UpdateOrderPaymentStatus(ctx context.Context, orderID, newPaymentStatus string, time time.Time) error
	GetOrder(ctx context.Context, orderID string) (*models.Order, error)
	ListOrderByEater(ctx context.Context, eaterID string, sort string, page, pageSize int) ([]*models.Order, error)
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

func (s *orderSvcImpl) SaveOrder(ctx context.Context, eaterID string, cart *models.Cart, address *models.AddressInfo) error {
	var (
		now = time.Now()
	)

	restaurant := &models.RestaurantInfo{
		Name:     cart.Restaurant.Name,
		ImageUrl: cart.Restaurant.ImageUrl,
	}

	delivery := &models.DeliveryInfo{
		Time: cart.Delivery.Time,
		Notes: cart.Delivery.Notes,
	}

	payment := &models.PaymentInfo{
		Method: cart.Payment.Method,
		CardID: cart.Payment.CardID,
		DeliveryPrice: cart.Payment.DeliveryPrice,
		TotalPrice: cart.Payment.TotalPrice,
	}
	orderItems := []*models.OrderItem{}

	for key, o := range cart.Items{
		orderItems[key] = &models.OrderItem{
			ID: rand.UUID(),
			ProductID: o.ProductID,
			Price: o.ProductPrice,
			Quantity: o.Quantity,
			TotalPrice: o.ProductPrice * o.Quantity,
			CreatedAt: now,
	}
}
order := &models.Order{
	ID: rand.UUID(),
	EaterID: eaterID,
	Instruction: cart.Instruction,
	RestaurantID: cart.RestaurantID,
	Restaurant: restaurant,
	Delivery: delivery,
	Payment: payment,
	Items: orderItems,
	Status: models.OrderStatusPending,
	PaymentStatus: models.PaymentStatusPending,
	CreatedAt: now,
	UpdatedAt: now,
}

err := s.orderRepo.WithTx(ctx, func(r repositories.OrderRepository) error {

	if err := r.SaveOrder(ctx, order); err != nil {
		return err
	}
	if err := r.SaveOrderItems(ctx, order.Items); err != nil {
		return err
	}
	return nil
})
if err != nil {
	return err
}
return nil
}

func (s *orderSvcImpl) UpdateOrder(ctx context.Context, order *models.Order) error {
		return s.orderRepo.UpdateOrder(ctx,order)
}

func (s *orderSvcImpl) UpdateOrderStatus(ctx context.Context, orderID, newStatus string, time time.Time) error {
		return s.orderRepo.UpdateOrderStatus(ctx,orderID, newStatus,time)
}

func (s *orderSvcImpl) UpdateOrderPaymentStatus(ctx context.Context, orderID, newPaymentStatus string, time time.Time) error {
		return s.orderRepo.UpdateOrderPaymentStatus(ctx, orderID, newPaymentStatus,time)
}

func (s *orderSvcImpl) DeleteOrder(ctx context.Context, orderID string) error {
 return s.orderRepo.DeleteOrder(ctx, orderID)
}

func (s *orderSvcImpl) GetOrder(ctx context.Context, orderID string) (*models.Order, error) {
 return s.orderRepo.GetOrder(ctx, orderID)
}

func (s *orderSvcImpl) ListOrderByEater(ctx context.Context, eaterID string, sort string, page, pageSize int) ([]*models.Order, error) {
  return s.orderRepo.ListOrderByEater(ctx, eaterID,sort, page,pageSize)
}