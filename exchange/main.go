package main

import (
	"fmt"
	"log"
	"exchange.com/cart"
)

func main() {
	// load the cart... into variable cart
	newCart := cart.Cart{}
	fmt.Println("cool")

	totalPrice, err := newCart.TotalPrice()
	if err != nil {
		log.Printf("impossible to compute price of the cart: %s", err)
		return
	}
	log.Println("Total Price", totalPrice.Display())

	err = newCart.Lock()
	if err != nil {
		log.Printf("impossible to lock the cart: %s", err)
		return
	}

}
