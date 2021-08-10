package cart

import (
	"errors"
	"time"
	"exchange.com/product"
	"exchange.com/user"
	"github.com/Rhymond/go-money"
)

type Cart struct {
    ID        string
    CreatedAt time.Time
    UpdatedAt time.Time
    lockedAt  time.Time
    user.User
    Items        []Item
    CurrencyCode string
    isLocked     bool
}

type Item struct {
    product.Product
    Quantity uint8
}

//The methods TotalPrice and Lock (bound to Cart ) are exported
// money and error are the result of the function
// A method is a function with a receiver
//The receiver of a method is a special parameter
//The receiver is not listed in the parameter list but before the method name
//A method can have only one receiver
func (c *Cart) TotalPrice() (*money.Money, error) {
    //...
 
	// initialize total price value
	totalprice := money.New(0, c.CurrencyCode) // create variable of type *money.Money
	var err error 
	for _, v := range c.Items {
		subtotal := v.Product.Price.Multiply(int64(v.Quantity)) // v.Product.Price receiver parameter
		totalprice, err = totalprice.Add(subtotal)
		if err != nil {
			return nil, err
		}
	}
	return totalprice, nil
}

func (c *Cart) Lock() error {

    //first check that the cart is not already locked 
    if c.isLocked {
        return errors.New("Cart is already locked")
	}
	c.isLocked = true
	c.lockedAt = time.Now()
	return nil
}


//the method delete bound to Cart is not exported
func (c *Cart) delete() error {
    // to implement
    return nil
}

