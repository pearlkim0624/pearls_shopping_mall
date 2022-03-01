// datastr.go defines constant variables, data structures,
// methods and initial functions.

package main

import (
	"errors"
	"fmt"
)

// MAXORDER is a maximum number of orders
// a customer has at the same time.
// It is set as 7 in this version.
const MAXORDER = 7

// MAXITEM is the number of items this shopping mall has.
const MAXITEM = 5

// INITPOINT is the amount of points
// given to a customer to shop
const INITPOINT = 1000

// Customer contains information for a customer.
type Customer struct {
	point    int            // remaining points
	orderNum int            // number of orders the customer has
	cart     map[string]int // map of item name and quantity
}

// newCustomer create a new Customer structure,
// initialize the values and return the pointer.
func newCustomer() *Customer {
	c := Customer{}
	c.point = INITPOINT
	c.cart = make(map[string]int)
	return &c
}

// getPoint returns the remaining points value
// the customer can buy new items.
func (c Customer) getPoint() int {
	return c.point
}

// updatePoint updates the points value with given value.
func (c *Customer) updatePoint(d int) (int, error) {
	if (c.point + d) < 0 {
		err := errors.New("cannot update points. points not enough")
		return c.point, err
	}

	c.point += d
	return c.point, nil
}

// getOrderNum returns the number of orders
// the customer has at that time.
func (c Customer) getOrderNum() int {
	return c.orderNum
}

// addOrderNum increases the number of orders
func (c *Customer) addOrderNum() (int, error) {
	c.orderNum++
	return c.orderNum, nil
}

// reduceOrderNum decreases the number of orders
func (c *Customer) reduceOrderNum() (int, error) {
	c.orderNum--
	return c.orderNum, nil
}

// displayCart prints the items in the cart.
func (c Customer) displayCart() {
	fmt.Println(c.cart)
}

// addToCart add the new order, which is a map of item name
// and quantity, into the cart, and return new cart and error.
func (c *Customer) addToCart(orders map[string]int) (map[string]int, error) {
	for k, v := range orders {
		if _, ok := c.cart[k]; ok { // item in cart, then add count
			c.cart[k] += v
		} else {
			c.cart[k] += v
		}
	}
	return c.cart, nil
}

// resetCart reset the cart.
func (c *Customer) resetCart() error {
	c.cart = map[string]int{}
	return nil
}

// ItemStr contains information for an item,
// which is sold in the shopping mall.
type ItemStr struct {
	name   string // item name
	price  int    // amount of points to buy this item
	amount int    // amount of remaining stocks
}

// ItemList is a list of items in the shopping mall.
type ItemList [MAXITEM]ItemStr

// newItemList initialize the ItemList
// and return the pointer.
func newItemList() *ItemList {
	var newList = ItemList{
		{"cellphone", 700, 10},
		{"earphone", 30, 10},
		{"snack", 2, 100},
		{"coffee", 5, 100},
		{"meal", 10, 50},
	}

	return &newList
}

// getItemIndex returns the index of the item in the ItemList
// if it exists in the list. Otherwise, return -1.
func (items ItemList) getItemIndex(name string) int { // return -1 with no item
	for i, v := range items {
		if v.name == name {
			return i
		}
	}
	return -1
}

// updateItemAmount updates the number of remaining items.
// If the item does not exist in ItemList or
// the update number is more than the existing number, then return error.
func (items *ItemList) updateItemAmount(name string, diffcount int) error {
	id := items.getItemIndex(name)
	if id == -1 {
		err := errors.New("no item with the name")
		return err
	}
	if (items[id].amount + diffcount) < 0 {
		err := errors.New("lack of item amount")
		return err
	}

	items[id].amount += diffcount
	return nil
}

// displayItem prints the current item information.
func (items *ItemList) displayItems() error {
	for i, v := range items {
		fmt.Printf("[%d] Name: %s, Price: %d, Amount in stock: %d\n", i+1, v.name, v.price, v.amount)
	}
	fmt.Println()
	return nil
}

// ORDER_ values are delivery status of an order.
const ORDER_NONE = "NONE"
const ORDER_ACCEPTED = "ACCEPTED"
const ORDER_SHIPPED = "SHIPPED"
const ORDER_DEV = "OUT FOR DELIVERY"
const ORDER_ARRIVED = "ARRIVED"

// ITEMSTOBUY is list of items to buy,
// which is a map of item name and quantity
type ITEMSTOBUY map[string]int

// OrderStr is a strunct for an order.
type OrderStr struct {
	status string
	items  ITEMSTOBUY
}

// OrderList is a list of orders
type OrderList [MAXORDER]OrderStr

// resetOrderStr resets the OrderStr after delivery and
// release the struct for next order.
func (order *OrderStr) resetOrderStr() {
	order.status = ORDER_NONE
	order.items = ITEMSTOBUY{}
}

// newOrderList creates OrderList,
// initialize the orders and return the pointer.
func newOrderList() *OrderList {
	list := new(OrderList)
	for i := range list {
		list[i].status = ORDER_NONE
		list[i].items = make(ITEMSTOBUY)
	}
	return list
}

// displayOrderList prints the contents of the current orders.
func (list OrderList) displayOrderList() {
	for i, v := range list {
		fmt.Printf("Order [%d]: status: %s ", i+1, v.status)
		fmt.Println(v.items)
	}
}
