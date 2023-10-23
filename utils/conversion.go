package utils

import (
	"encoding/json"

	"github.com/haakaashs/antino-labs/models"
	"github.com/haakaashs/antino-labs/resources"
)

func ResourceToModel(cart resources.CartResource) models.Cart {
	output := models.Cart{
		ID:          cart.ID,
		UserID:      cart.UserID,
		Qty:         cart.Qty,
		TotalAmount: cart.TotalAmount,
	}
	output.CartProducts, _ = json.Marshal(cart.CartProducts)
	return output
}

func ModelToResource(cart models.Cart) (resources.CartResource, error) {
	output := resources.CartResource{
		ID:          cart.ID,
		UserID:      cart.UserID,
		Qty:         cart.Qty,
		TotalAmount: cart.TotalAmount,
	}
	err := json.Unmarshal(cart.CartProducts, &output.CartProducts)
	if err != nil {
		return output, err
	}
	return output, nil
}
