package main

import (
	"fmt"
	"log"
	"time"
)

// main function of shopping mall.
// main creates basic data structures,
// one go channel for delivery and
// display main menu to a user.
func main() {
	cust := newCustomer()
	itemList := newItemList()
	orderList := newOrderList()
	orderChannel := make(chan ITEMSTOBUY)

	go orderMain(orderChannel, cust, orderList)

	mainMenu(cust, itemList, orderList, orderChannel)
}

func mainMenu(cust *Customer, itemList *ItemList, orderList *OrderList, oc chan ITEMSTOBUY) {
	var input int

	for {
		input = 0

		fmt.Println("Main Menu")
		fmt.Println("1. Shopping")
		fmt.Println("2. My Cart")
		fmt.Println("3. My Points")
		fmt.Println("4. My Order")
		fmt.Println("5. Display Items")
		fmt.Println("6. Exit")

		fmt.Scanln(&input)

		switch input {
		case 1: // Shopping
			shoppingMenu(cust, itemList, oc)
		case 2: // My Cart
			cust.displayCart()
			myCartMenu(cust, itemList, oc)
		case 3: // My Points
			point := cust.getPoint()
			fmt.Println("Your point is ", point)
		case 4: // My Order
			orderList.displayOrderList()
		case 5: // Display Items
			itemList.displayItems()
		case 6: // Exit
			fmt.Println("Bye")
			return
		default:
			fmt.Println("Wrong Input. Choose again")
		}
	}
}

func myCartMenu(cust *Customer, itemList *ItemList, oc chan ITEMSTOBUY) {
	var input int

	for {
		fmt.Println("MyCart Menu")
		fmt.Println("1. Buy Now")
		fmt.Println("2. Reset Cart")
		fmt.Println("3. Goto Previouse Menu")
		fmt.Scanln(&input)

		switch input {
		case 1:
			buyNowMenu(cust, itemList, cust.cart, oc)
			cust.resetCart()
			return
		case 2:
			cust.resetCart()
			fmt.Println("Your cart is empty now")
			return
		case 3:
			return
		default:
			fmt.Println("Wrong Input. Choose menu again")
		}
	}
}

func shoppingMenu(cust *Customer, itemList *ItemList, oc chan ITEMSTOBUY) {
	itemList.displayItems()

	var input int

	for {
		fmt.Println("Choose Item to buy")
		fmt.Println("Enter 0 to goto previous menu")
		fmt.Scanln(&input)

		if input == 0 { // Goto Previous Menu
			return
		}
		if input > 0 && input <= MAXITEM {
			chooseItemMenu(cust, itemList, input-1, oc)
			return
		} else {
			fmt.Println("Wrong Input. Choose again")
		}
	}
}

func chooseItemMenu(cust *Customer, itemList *ItemList, index int, oc chan ITEMSTOBUY) {
	var input int
	item := itemList[index]

	for {
		fmt.Println("Input amounts to buy. We have", item.amount, "in stock")
		fmt.Println("Enter 0 to goto previous menu")
		fmt.Scanln(&input)

		if input == 0 { // Goto Previous Menu
			return
		}

		if input > 0 && input <= item.amount {
			itemsToBuy := make(map[string]int, 1)
			itemsToBuy[item.name] = input
			inputCountMenu(cust, itemList, itemsToBuy, oc)
			return
		} else {
			fmt.Println("Wrong Input. Please input within", item.amount)
		}
	}
}

func inputCountMenu(cust *Customer, itemList *ItemList, itemsToBuy ITEMSTOBUY, oc chan ITEMSTOBUY) {
	var input int

	for {
		fmt.Println("InputCounts Menu")
		fmt.Println("1. Buy Now")
		fmt.Println("2. Add to Cart")
		fmt.Println("3. Goto Previouse Menu")
		fmt.Scanln(&input)

		switch input {
		case 1: // Buy Now
			buyNowMenu(cust, itemList, itemsToBuy, oc)
			return
		case 2: // Add to cart
			newCart, err := cust.addToCart(itemsToBuy)
			if err != nil {
				log.Println(err)
			}
			fmt.Println("Items added to the cart. The cart is ", newCart)
			return
		case 3: // Goto Previous Menu
			return
		default:
			fmt.Println("Wrong Input. Choose menu again")
		}
	}
}

func buyNowMenu(cust *Customer, itemList *ItemList, itemsToBuy map[string]int, oc chan ITEMSTOBUY) {
	fmt.Println("Buy Now")

	/* check points. item amount and order number
	update points, item amount
	send request to orderMain through orderChannel
	*/
	if cust.getOrderNum() >= MAXORDER {
		log.Println("Order request denied. Maximum order number is reached. Try later")
		return
	}

	total := 0
	for k, v := range itemsToBuy {
		i := itemList.getItemIndex(k)
		if i == -1 {
			log.Println("wrong item name", k)
			return
		}
		if itemList[i].amount < v {
			log.Println("Lack of stocks", k)
			return
		}
		total += itemList[i].price * v
	}

	if total > cust.point {
		log.Printf("Lack of points. Your point is %d and your request total is %d\n", cust.point, total)
		return
	}
	cust.updatePoint(total * (-1))

	for k, v := range itemsToBuy {
		itemList.updateItemAmount(k, v*(-1))
	}

	fmt.Printf("Point is %d and total cost is %d\n", cust.point, total)
	fmt.Println(itemsToBuy)

	oc <- itemsToBuy

}

func orderMain(oc <-chan ITEMSTOBUY, cust *Customer, orderList *OrderList) {
	fmt.Println("start orderMain")

	dc := make(chan int)
	index := 0

	for {
		select {
		case items := <-oc:

			// New order - make orderStr, increase orderNum and go delivery
			if items == nil {
				log.Println("nil type input in order channel")
				break // break select
			}

			if orderList[index].status != ORDER_NONE {
				log.Println("OrderList error")
				fmt.Printf("index [%d]\n", index)
				fmt.Println(orderList[index])
				break // break select
			}
			fmt.Println("Order accepted")
			orderList[index].status = ORDER_ACCEPTED
			for k, v := range items {
				orderList[index].items[k] = v
			}

			fmt.Println(orderList[index])
			go orderDelivery(dc, index, &orderList[index])

			cust.addOrderNum()
			index = (index + 1) % MAXORDER

		case id := <-dc:
			// One order delivery done.
			if orderList[id].status != ORDER_NONE {
				log.Println("Delivery status of [", id, "] is ", orderList[id].status)
				break // break select
			}

			fmt.Println("Delivery of", id, "is done", orderList[id])

			cust.reduceOrderNum()
			orderList[id].resetOrderStr()

			fmt.Println("Order Num is ", cust.getOrderNum())
			fmt.Println(orderList)
		}
	}
}

func orderDelivery(dc chan<- int, index int, order *OrderStr) {
	// update status
	fmt.Println("orderDelivery", order)

	time.Sleep(time.Second * 5)
	order.status = ORDER_SHIPPED
	time.Sleep(time.Second * 5)
	order.status = ORDER_DEV
	time.Sleep(time.Second * 5)
	order.status = ORDER_ARRIVED
	time.Sleep(time.Second * 5)
	order.status = ORDER_NONE

	dc <- index
}
