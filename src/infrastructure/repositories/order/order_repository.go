package order

import (
	"context"
	"time"

	"github.com/DevShuxat/eater-service/src/domain/order/models"
	"github.com/DevShuxat/eater-service/src/domain/order/repositories"
	"github.com/DevShuxat/eater-service/src/infrastructure/utils"
	"gorm.io/gorm"
)

const (
	tableOrder = "eater.orders"
)

type orderSvcImpl struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) repositories.OrderRepository {
	return &orderSvcImpl{
		db: db,
	}
}


func (r *orderSvcImpl) SaveOrder(ctx context.Context, order *models.Order) error {
	err := r.db.WithContext(ctx).Create(order).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *orderSvcImpl) GetOrder(ctx context.Context, orderID string) (*models.Order, error) {
	var order *models.Order
	result := r.db.WithContext(ctx).Table(tableOrder).First(&order, "id = ?", orderID)
	if result.Error != nil {
		return nil, result.Error
	}
	return order, nil
}

// func (r *orderSvcImpl) ListOrder(ctx context.Context, eaterID string, sortByDateDesc bool) ([]models.Order, error) {
// 	var orders []models.Order
// 	query := r.db.WithContext(ctx).Table(tableOrder).Where("eater_id = ?", eaterID)
// 	if sortByDateDesc {
// 		query = query.Order("created_date desc")
// 	} else {
// 		query = query.Order("created_date asc")
// 	}
// 	if err := query.Find(&orders).Error; err != nil {
// 		return nil, err
// 	}
// 	return orders, nil
// }

func (r *orderSvcImpl) UpdateOrder(ctx context.Context, order *models.Order) error {
	result := r.db.WithContext(ctx).Table(tableOrder).Save(order)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *orderSvcImpl) UpdateOrderPaymentStatus(ctx context.Context, orderID, newPaymentStatus string, time time.Time) error {
	result := r.db.WithContext(ctx).Table(tableOrder).Where("id = ?", orderID).Save("payment_status")
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *orderSvcImpl) UpdateOrderStatus(ctx context.Context, orderID, newStatus string, time time.Time) error {
	result := r.db.WithContext(ctx).Table(tableOrder).Where("id = ?", orderID).Save("status")
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *orderSvcImpl) ListOrderByEater(ctx context.Context, eaterID string, sort string, page, pageSize int) ([]*models.Order, error) {
	var orders []*models.Order
	result := r.db.WithContext(ctx).Table(tableOrder).Where("eater_id = ?", eaterID).
	Scopes(utils.Paginate(page, pageSize), utils.Sort(sort)).
	Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}

func (r *orderSvcImpl) DeleteOrder(ctx context.Context, orderID string) error {
	var order models.Order
	result := r.db.WithContext(ctx).Table(tableOrder).Delete(&order, "id = ?", orderID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}


