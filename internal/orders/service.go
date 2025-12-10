package orders

import (
	"errors"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Service interface{
	CreateNewOrder(*NewOrder) (OrderDetails, error)
}

type svc struct {
	DB *pgxpool.Pool
}

func NewService() Service {
	return &svc{}
}

func (s *svc) CreateNewOrder(order *NewOrder) (OrderDetails, error) {
	if order == nil {
		return OrderDetails{}, errors.New("order is nil")
	}

	// call repo to store and update with new order
	

	price, _ := strconv.ParseFloat(order.Price, 64)

	return OrderDetails{
		Name: order.ItemName,
		Price: price,
		Size: order.Size,
		Status: "created",
	}, nil
}