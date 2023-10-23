package utils

import (
	"encoding/json"

	"github.com/haakaashs/antino-labs/models"
	"github.com/haakaashs/antino-labs/resources"
)

func ModelToResource(c *models.Cart) (resources.CartResource, error) {
	output := resources.CartResource{
		ID:          c.ID,
		UserID:      c.UserID,
		Qty:         c.Qty,
		TotalAmount: c.TotalAmount,
	}
	err := json.Unmarshal(c.CartProducts, &output.CartProducts)
	if err != nil {
		return output, err
	}
	return output, nil
}

func ResourceToModel(c *resources.CartResource) (models.Cart, error) {
	var err error
	output := models.Cart{
		ID:          c.ID,
		UserID:      c.UserID,
		Qty:         c.Qty,
		TotalAmount: c.TotalAmount,
	}
	output.CartProducts, err = json.Marshal(c.CartProducts)
	if err != nil {
		return output, err
	}
	return output, nil
}
