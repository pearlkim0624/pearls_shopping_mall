package main

import (
	"errors"
	"fmt"
)

const MAXORDER = 7
const MAXITEM = 5
const INITPOINT = 1000

type Customer struct {
	point    int
	orderNum int
	cart     map[string]int
}

func newCustomer() *Customer {
	c := Customer{}
	c.point = INITPOINT
	c.cart = make(map[string]int)
	return &c
}

func (c Customer) getPoint() int {
	return c.point
}

func (c *Customer) updatePoint(d int) (int, error) {
	if (c.point + d) < 0 {
		err := errors.New("cannot update points. points not enough")
		return c.point, err
	}

	c.point += d
	return c.point, nil
}

func (c Customer) getOrderNum() int {
	return c.orderNum
}

func (c *Customer) addOrderNum() (int, error) {
	c.orderNum++
	return c.orderNum, nil
}

func (c *Customer) reduceOrderNum() (int, error) {
	c.orderNum--
	return c.orderNum, nil
}

func (c Customer) displayCart() {
	fmt.Println(c.cart)
}

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

func (c *Customer) resetCart() error {
	c.cart = map[string]int{}
	return nil
}

type ItemStr struct {
	name   string
	price  int
	amount int
}

type ItemList [MAXITEM]ItemStr

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

func (items ItemList) getItemIndex(name string) int { // return -1 with no item
	for i, v := range items {
		if v.name == name {
			return i
		}
	}
	return -1
}

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

func (items *ItemList) displayItems() error {
	for i, v := range items {
		fmt.Printf("[%d] Name: %s, Price: %d, Amount in stock: %d\n", i+1, v.name, v.price, v.amount)
	}
	fmt.Println()
	return nil
}

const ORDER_NONE = "NONE"
const ORDER_ACCEPTED = "ACCEPTED"
const ORDER_SHIPPED = "SHIPPED"
const ORDER_DEV = "OUT FOR DELIVERY"
const ORDER_ARRIVED = "ARRIVED"

type ITEMSTOBUY map[string]int

type OrderStr struct {
	status string
	items  ITEMSTOBUY
}
type OrderList [MAXORDER]OrderStr

func (order *OrderStr) resetOrderStr() {
	order.status = ORDER_NONE
	order.items = ITEMSTOBUY{}
}

func newOrderList() *OrderList {
	list := new(OrderList)
	for i := range list {
		list[i].status = ORDER_NONE
		list[i].items = make(ITEMSTOBUY)
	}
	return list
}

func (list OrderList) displayOrderList() {
	for i, v := range list {
		fmt.Printf("Order [%d]: status: %s ", i+1, v.status)
		fmt.Println(v.items)
	}
}
