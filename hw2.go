// Homework 2: Object Oriented Programming
// Due February 7, 2017 at 11:59pm
package main

import (
	"fmt"
	"log"
)

func main() {
	// Feel free to use the main function for testing your functions
	world := struct {
		English string
		Spanish string
		French  string
	}{
		"world",
		"mundo",
		"monde",
	}
	fmt.Printf("Hello, %s\t%s\t%s!\n", world.English, world.Spanish, world.French)

	RegisterItem(Prices, "qurt", 100)
	RegisterItem(Prices, "baursaq", 50)
	RegisterItem(Prices, "közhe", 250)
	fmt.Printf("Prices: %v\n", Prices)

	c := new(Cart)
	c.AddItem("qurt")
	c.AddItem("baursaq")
	c.AddItem("közhe")
	fmt.Printf("c.hasMilk() = %v\n", c.hasMilk())
	fmt.Printf("c.HasItem(%v) = %v\n", "baursaq", c.HasItem("baursaq"))

	c.AddItem("milk")
	c.Checkout()

}

// Price is the cost of something in US cents.
type Price int64

// String is the string representation of a Price
// These should be represented in US Dollars
// Example: 2595 cents => $25.95
func (p Price) String() string {
	// TODO
	var num = float64(p) / 100
	var string = fmt.Sprintf("$%.2f", num)
	return string
}

// Prices is a map from an item to its price.
var Prices = map[string]Price{
	"eggs":          219,
	"bread":         199,
	"milk":          295,
	"peanut butter": 445,
	"chocolate":     150,
}

// RegisterItem adds the new item in the prices map.
// If the item is already in the prices map, a warning should be displayed to the user,
// but the value should be overwritten.
// Bonus (1pt) - Use the "log" package to print the error to the user
func RegisterItem(prices map[string]Price, item string, price Price) {
	if _, ok := Prices[item]; ok {
		log.Printf("item %v already registered\n", item)
	}
	prices[item] = price

}

// Cart is a struct representing a shopping cart of items.
type Cart struct {
	Items      []string
	TotalPrice Price
}

// hasMilk returns whether the shopping cart has "milk".
func (c *Cart) hasMilk() bool {
	for _, item := range c.Items {
		if item == "milk" {
			return true
		}
	}
	return false
}

// HasItem returns whether the shopping cart has the provided item name.
func (c *Cart) HasItem(item string) bool {
	// TODO
	for _, i := range c.Items {
		if i == item {
			return true
		}
	}
	return false
}

// AddItem adds the provided item to the cart and update the cart balance.
// If item is not found in the prices map, then do not add it and print an error.
// Bonus (1pt) - Use the "log" package to print the error to the user
func (c *Cart) AddItem(item string) {
	if price, ok := Prices[item]; ok {
		c.Items = append(c.Items, item)
		c.TotalPrice += price
	} else {
		log.Printf("item %v has no price\n", item)
	}
}

// Checkout displays the final cart balance and clears the cart completely.
func (c *Cart) Checkout() {
	fmt.Println("Cart list")
	for _, i := range c.Items {
		fmt.Println(i, Prices[i])
	}

	fmt.Printf("Total: %v\n", c.TotalPrice)

	c.Items = c.Items[:0]
	c.TotalPrice = 0
}
