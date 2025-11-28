package orders

import (
	"errors"
	"strconv"
)

type Service interface{
	CreateNewOrder(*NewOrder) (OrderDetails, error)
}

type svc struct {
	// declare dependencies for service module
}

func NewService() Service {
	return &svc{}
}

func (s *svc) CreateNewOrder(order *NewOrder) (OrderDetails, error) {
	if order == nil {
		return OrderDetails{}, errors.New("order is nil")
	}

	price, _ := strconv.ParseFloat(order.Price, 64)

	return OrderDetails{
		Name: order.ItemName,
		Price: price,
		Size: order.Size,
		Status: "created",
	}, nil
}