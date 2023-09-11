package order

import (
	"context"
	"github.com/DevShuxat/eater-service/src/domain/order/models"
	"github.com/DevShuxat/eater-service/src/domain/order/repositories"
	"gorm.io/gorm"
)

const (
	tableOrder = "orders"
)

type NewOrderRepository struct {
	db *gorm.DB
}


func (r *NewOrderRepository) CreateOrder(ctx context.Context, order *models.Order) error {
	err := r.db.WithContext(ctx).Create(order).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *NewOrderRepository) DeleteOrder(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Delete(&models.Order{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *NewOrderRepository) GetOrder(ctx context.Context, orderID string) (*models.Order, error) {
	var order models.Order
	result := r.db.WithContext(ctx).First(&order, "id = ?", orderID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

func (r *NewOrderRepository) ListOrder(ctx context.Context, eaterID string, sortByDateDesc bool) ([]models.Order, error) {
	var orders []models.Order
	query := r.db.WithContext(ctx).Table(tableOrder).Where("eater_id = ?", eaterID)

	if sortByDateDesc {
		query = query.Order("created_date desc")
	} else {
		query = query.Order("created_date asc")
	}

	if err := query.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *NewOrderRepository) UpdateOrder(ctx context.Context, order *models.Order) error {
	result := r.db.WithContext(ctx).Save(order)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *NewOrderRepository) UpdateOrderPaymentStatus(ctx context.Context, orderID string, paymentStatus string) error {
	result := r.db.WithContext(ctx).Model(&models.Order{}).Where("id = ?", orderID).Update("payment_status", paymentStatus)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *NewOrderRepository) UpdateOrderStatus(ctx context.Context, orderID string, status string) error {
	result := r.db.WithContext(ctx).Model(&models.Order{}).Where("id = ?", orderID).Update("status", status)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *NewOrderRepository) WithTx(ctx context.Context, callback func(r repositories.OrderRepository) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		r := NewOrderRepository{
			db: tx,
		}
		return callback(&r)
	})
}
