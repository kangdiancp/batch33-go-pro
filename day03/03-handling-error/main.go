package main

import (
	"errors"
	"fmt"
)

func addToCart(stock *int, total int) (int, error) {
	//check apakah jumlah total yg ditambahkan valid
	if total <= 0 {
		return 0, errors.New("jumlah item diinput > 0")
	}

	//chcek apakah stock cukup
	if *stock < total {
		return 0, errors.New("stock tidak cukup untuk ditambahkan ke keranjang")
	}

	//mengurangi stok dan return total item yg ditambahkan
	*stock -= total
	return total, nil
}

func main() {

	//1. initial stock & totalItem
	stok := 100
	totalItem := 0

	//add 3 item to cart
	newTotal, err := addToCart(&stok, 3)
	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		totalItem += newTotal
		fmt.Printf("Success add %d item, stock sisa: %d, total di cart :%d\n", newTotal, stok, totalItem)
	}

	//if total item value-nya negatif
	newTotal, err = addToCart(&stok, -1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		totalItem += newTotal
		fmt.Printf("Success add %d item, stock sisa: %d, total di cart :%d\n", newTotal, stok, totalItem)
	}

	//if total item > dari stok
	newTotal, err = addToCart(&stok, 200)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		totalItem += newTotal
		fmt.Printf("Success add %d item, stock sisa: %d, total di cart :%d\n", newTotal, stok, totalItem)
	}

}
