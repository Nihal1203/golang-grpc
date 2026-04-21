package types

import (
	"context"

	"github.com/Nihal1203/golang-grpc/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
}
