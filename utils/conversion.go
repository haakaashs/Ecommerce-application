package utils

import (
	"encoding/json"

	"github.com/haakaashs/antino-labs/models"
	"github.com/haakaashs/antino-labs/resources"
)

func CartResourceToModel(cart resources.CartResource) models.Cart {
	output := models.Cart{
		ID:          cart.ID,
		UserID:      cart.UserID,
		Qty:         cart.Qty,
		Discount:    cart.Discount,
		TotalAmount: cart.TotalAmount,
	}
	output.CartProducts, _ = json.Marshal(cart.CartProducts)
	return output
}

func CartModelToResource(cart models.Cart) (resources.CartResource, error) {
	output := resources.CartResource{
		ID:          cart.ID,
		UserID:      cart.UserID,
		Qty:         cart.Qty,
		Discount:    cart.Discount,
		TotalAmount: cart.TotalAmount,
	}
	err := json.Unmarshal(cart.CartProducts, &output.CartProducts)
	if err != nil {
		return output, err
	}
	return output, nil
}

func OrderModelToResource(order models.Order) (resources.OrderResource, error) {
	output := resources.OrderResource{
		ID:          order.ID,
		UserID:      order.UserID,
		OrderStatus: order.OrderStatus,
		OrderValue:  order.OrderValue,
		Discount:    order.Discount,
		Dispatched:  order.Dispatched,
		IsActive:    order.IsActive,
	}
	err := json.Unmarshal(order.OrderProducts, &output.OrderProducts)
	if err != nil {
		return output, err
	}
	return output, nil
}
