package main

import "fmt"

func main() {
	fmt.Println("Just use enter to skip if you do not have a the var it will work around it")
	var price float32
	fmt.Println("What is the price of what you are selling: ")
	fmt.Scan(&price)
	var amount float32
	fmt.Println("What is the amount of what you are selling: ")
	fmt.Scan(&amount)
	var gain float32
	fmt.Println("What is the percent gain you want to make [ex 5% is 5:] ")
	fmt.Scan(&gain)
	var decimal_percentage float32
	decimal_percentage = gain*.01 + 1

	fmt.Println(decimal_percentage)
	target_buy_price := price * decimal_percentage
	target_buy_amount := amount * decimal_percentage
	sold_for := amount * price

	fmt.Println("New order price", target_buy_price)
	//how much you sold your order for
	fmt.Println("You sold your order for ", sold_for)

	fmt.Println("You need to buy ", target_buy_amount)

}
