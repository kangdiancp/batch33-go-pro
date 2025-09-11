package main

import "fmt"

func addStock(stock int, total int) int {
	fmt.Printf("[%s]\taddr:[%p]\tvalue:[%d]\n", "stock-func", &stock, stock)
	return stock + total
}

func main() {
	stock := 100
	fmt.Printf("[%s]\taddr:[%p]\tvalue:[%d]\n", "stock-before", &stock, stock)

	stockAdd := addStock(stock, 15)
	fmt.Printf("[%s]\taddr:[%p]\tvalue:[%d]\n", "stock-add", &stockAdd, stockAdd)

	fmt.Printf("[%s]\taddr:[%p]\tvalue:[%d]\n", "stock-after", &stock, stock)

	//create function sell stock
}
