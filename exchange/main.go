package main

import (
	"database/sql"
	"log"
	"math/rand"
	"strconv"
	"time"

	"exchange.com/cart"
	_ "github.com/go-sql-driver/mysql"
)

type CartTrans interface {

	GetCartID(c *cart.Cart)(ID string)
	PutCart(c *cart.Cart)(err error)
}

func GetCartID(c *cart.Cart) (ID string) {
	
    return c.ID
}

func ErrorCheck(err error) {
    if err != nil {
        panic(err.Error())
    }
}


func main() {
    rand.Seed(time.Now().UTC().UnixNano())
	Id := strconv.Itoa(rand.Intn(1000))
	newCart := cart.Cart{ID: Id , CurrencyCode:"129M", CreatedAt: time.Now()}
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/cart")
	ErrorCheck(err)
    log.Println("Connected!")
    
	//When a package is imported prefixed with a blank identifier, the init function of the package is called.
	// The function registers the driver. 
	//err = db.Query("insert into Cart values = ? ", 3).Scan(&newCart.ID, &newCart.CurrencyCode, &newCart.CreatedAt)
        // INSERT INTO DB
    // prepare
    stmt, e := db.Prepare("insert into Cart(ID, user, createdat) values (?, ?, ?)")
    ErrorCheck(e)
 
    //execute
    res, e := stmt.Exec(newCart.ID, newCart.CurrencyCode, newCart.CreatedAt)
    ErrorCheck(e)
 
    id, er := res.LastInsertId()
    ErrorCheck(er)
 
    log.Println("Insert id", id)

    defer db.Close()


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
