package cart

import (
	"fmt"
	"testing"
	"time"
	"exchange.com/product"
	"exchange.com/user"
	"exchange.com/cart"
	"github.com/Rhymond/go-money"
	"github.com/stretchr/testify/assert"
)

func TestTotalPrice(t *testing.T) {
	items := []cart.Item{

		{
			Product: product.Product{
				ID:    "p-25",
				Name:  "banana",
				Price: money.New(10, "EUR"),
			},
			Quantity: 2,
		},
		{
			Product: product.Product{
				ID:    "p-26",
				Name:  "paddle",
				Price: money.New(800, "EUR"),
			},
			Quantity: 1,
		},
	}
	c := cart.Cart{
		ID:           "1254",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		User:         user.User{},
		Items:        items,
		CurrencyCode: "EUR",
	}
	
	actual, err := c.TotalPrice()
	assert.NoError(t, err)
	assert.Equal(t, money.New(820, "EUR"), actual)

}
