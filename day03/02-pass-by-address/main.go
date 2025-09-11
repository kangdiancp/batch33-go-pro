package main

import "fmt"

func addToStock(stock *int, total int) int {
	*stock = *stock + total
	return *stock
}

func addToStockPtr(stock *int, total int) *int {
	*stock = *stock + total
	return stock
}

func main() {
	stock := 100
	cart1 := &stock
	cart2 := &cart1

	fmt.Printf("[%s]\taddr:[%p]\tvalue:[%v]\n", "stock", &stock, stock)
	fmt.Printf("[%s]\taddr:[%p]\tvalue:[%v]\n", "cart1", &cart1, cart1)
	fmt.Printf("[%s]\taddr:[%p]\tvalue:[%v]\n", "cart2", &cart2, cart2)

	//display output
	fmt.Printf("[%s] value:[%v]\n", "cart1", *cart1)
	fmt.Printf("[%s] value:[%v]\n", "cart2", **cart2)

	//modify value 100-90
	*cart1 = *cart1 - 10
	fmt.Printf("[%s]\taddr:[%p]\tvalue:[%v]\n", "cart1", &cart1, cart1)
	fmt.Printf("[%s]\taddr:[%p]\tvalue:[%v]\n", "stock", &stock, stock)

	// cart1 = 90+15
	*cart1 = *cart1 + 15
	fmt.Printf("[%s]\taddr:[%p]\tvalue:[%v]\n", "stock", &stock, stock)

	//call func addStock
	addStock := addToStock(&stock, 15)
	fmt.Printf("[%s]\taddr:[%p]\tvalue:[%v]\n", "addStock", &addStock, addStock)
	fmt.Printf("[%s]\taddr:[%p]\tvalue:[%v]\n", "stock", &stock, stock)

	//call func addStock
	addStockPtr := addToStockPtr(&stock, 50)
	fmt.Printf("[%s]\taddr:[%p]\tvalue:[%v]\n", "addStockPtr", &addStockPtr, addStockPtr)
	fmt.Printf("[%s]\taddr:[%p]\tvalue:[%v]\n", "stock", &stock, stock)

}
