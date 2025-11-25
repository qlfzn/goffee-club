package orders

import (
	"errors"
	"strconv"
)

type NewOrder struct {
	ItemName string  `json:"itemName"`
	Size     string  `json:"size"`
	Price    string `json:"price"`
}

type OrderDetails struct {
	Name string
	Size string
	Price float64
	Status string
}

func (o *OrderDetails) CreateNewOrder(order *NewOrder) (OrderDetails, error) {
	if order == nil {
		return OrderDetails{}, errors.New("order is nil")
	}

	price, _ := strconv.ParseFloat(order.Price, 64)

	o.Name = order.ItemName
	o.Size = order.Size
	o.Price = price
	o.Status = "created"
	
	return *o, nil
}