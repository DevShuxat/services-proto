package services

import (
	"context"


	// dtos "github.com/DevShuxat/eater-service/src/application/dtos/order"
	pb "github.com/DevShuxat/eater-service/src/application/protos/eater"
	ordersvc "github.com/DevShuxat/eater-service/src/domain/order/services"
	addresssvc "github.com/DevShuxat/eater-service/src/domain/address/services"
	// "github.com/DevShuxat/eater-service/src/domain/order/models"
)

type OrderApplicationService interface {
	SaveOrder(ctx context.Context, req *pb.PlaceOrderRequest ) (*pb.PlaceOrderResponse, error)
	UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error)
	GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error)
	ListOrderByEater(ctx context.Context, req *pb.ListOrderByEaterRequest) (*pb.ListOrderByEaterResponse, error)
	DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error)
}


type orderSvcImpl struct {
	orderSvc ordersvc.OrderService
	addressSvc addresssvc.AddressService

}

func NewOrderApplicationService (
	orderSvc ordersvc.OrderService,
	addressSvc addresssvc.AddressService,
	) OrderApplicationService {
	return &orderSvcImpl {
	orderSvc: orderSvc,
	addressSvc: addressSvc,
	}
}

	func(s *orderSvcImpl) SaveOrder(ctx context.Context, req *pb.PlaceOrderRequest ) (*pb.PlaceOrderResponse, error) {
		if 
	}
	func(s *orderSvcImpl) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error)
	func(s *orderSvcImpl) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error)
	func(s *orderSvcImpl) ListOrderByEater(ctx context.Context, req *pb.ListOrderByEaterRequest) (*pb.ListOrderByEaterResponse, error)
	func(s *orderSvcImpl) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error)