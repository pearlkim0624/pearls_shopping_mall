# Pearl's Shopping Mall

This is a text based shopping mall implemented in Go. It is implemented as a final project of "Go Tutorial" in goormedu, https://edu.goormedu.io.

You, a customer, is given 1,000 points and can buy some products in the shopping mall using the points. There are 5 products in shopping mall. Each product has a price in points and available number in stock. You can purchase products directly or put them in the cart and purchase them later. You can check your points, the current content of the cart, or the delivery status of your orders. 

The program is implemented in Go, including struct, slice, map, method, go routine, and channel. Main function initializes data structures for customer, orders, and items in the shopping mall. It also invokes orderMain go routine which manages the orders and deliveries. When there is new order of buying products, main sends the content of the new order to orderMain and orderMain invokes new go routine to handle the delivery of the order. When a delivery routine complete the delivery, it notifies to ordrMain and ends. 

For the future improvement, it may have multiple users, web-based user interface, etc.

Pearl
3/2/2022
