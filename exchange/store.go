package exchange

import (

	"exchange.com/cart"
    "fmt"
	"github.com/go-sql-driver/mysql"


)

type CartTrans interface {

	GetCartID(c *cart.Cart)(ID string, err error)
	PutCart(c *cart.Cart)(err error)
}

func GetCartID(c *cart.Cart) (ID string, err error) {
	var err error
    if c.ID != nil {
		return c.ID, err
	} else {
	    return c.ID, nil
}
}

func PutCart(c *cart.Cart)(err error) {

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/cart")
	insert, err := db.Query("INSERT INTO Cart VALUES ('MariumAZ', '1127', '2020-01-01 15:10:10');")

}

