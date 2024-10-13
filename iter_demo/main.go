package main

import (
	"cmp"
	"fmt"
	"iter"
	"slices"
)

// Order represents a customer order in a food delivery service
type Order struct {
	OrderID      string
	CustomerName string
	Amount       float64
	Status       string // Status could be "pending", "delivered", or "canceled"
}

// filter orders and return an iterator
func filter[V any](it iter.Seq[V], keep func(V) bool) iter.Seq[V] {
	seq := func(yield func(V) bool) {
		for v := range it {
			if keep(v) {
				if !yield(v) {
					break
				}
			}
		}
	}
	return seq
}

// display prints the orders from an iterator
func display(it iter.Seq[Order]) {
	for order := range it {
		fmt.Printf("Order ID: %s, Customer: %s, Amount: $%.2f, Status: %s\n",
			order.OrderID, order.CustomerName, order.Amount, order.Status)
	}
}

func main() {
	// Example slice of orders
	orders := []Order{
		{OrderID: "1", CustomerName: "Alice", Amount: 75.50, Status: "delivered"},
		{OrderID: "2", CustomerName: "Bob", Amount: 45.00, Status: "delivered"},
		{OrderID: "3", CustomerName: "Charlie", Amount: 90.75, Status: "pending"},
		{OrderID: "4", CustomerName: "Dana", Amount: 30.99, Status: "canceled"},
		{OrderID: "5", CustomerName: "Eve", Amount: 60.10, Status: "delivered"},
	}

	// Filter orders with amount greater than $50
	highValueOrders := filter(slices.Values(orders), func(order Order) bool {
		return order.Amount > 50
	})
	fmt.Println("Orders with Amount > $50:")
	display(highValueOrders)

	// Filter orders with status "delivered"
	deliveredOrders := filter(slices.Values(orders), func(order Order) bool {
		return order.Status == "delivered"
	})
	fmt.Println("\nDelivered Orders:")
	display(deliveredOrders)

	// Explore Collect() function
	seq := func(yield func(Order) bool) {
		for _, o := range orders {
			if o.Amount < 50 {
				if !yield(o) {
					return
				}
			}
		}
	}
	filteredOrders := slices.Collect(seq)

	// Display collected orders
	fmt.Println("\nCollected orders with price < $50:")
	for _, o := range filteredOrders {
		fmt.Printf("%s: $%.2f\n", o.OrderID, o.Amount)
	}

	// Sort orders by amount in descending order using SortedFunc()
	sortFunc := func(a, b Order) int {
		return cmp.Compare(b.Amount, a.Amount) // sort in descending order
	}
	sortedOrders := slices.SortedFunc(slices.Values(orders), sortFunc)
	// Display sorted orders
	fmt.Println("\nSorted Orders by Amount (Descending):")
	for _, o := range sortedOrders {
		fmt.Printf("Order ID: %s, Customer: %s, Amount: $%.2f, Status: %s\n", o.OrderID, o.CustomerName, o.Amount, o.Status)
	}

	// Chunk orders into []Order 3 elements at a time.
	fmt.Println("\nChunks of orders of size 3:")
	for c := range slices.Chunk(orders, 3) {
		fmt.Println(c)
	}
}
