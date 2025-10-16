// this file defines a fake database layer called DatabaseHandler
// represents the "data source" for the REST API
// holds two slices in memory (list of drinks and orders)

package repository

import (
	"ordersystem/model"
	"time"
) // fetches from model folder

type DatabaseHandler struct {
	// drinks represent all available drinks
	drinks []model.Drink //slice; "danyamic Array" can grow and shrink, each element is one drink
	// orders serves as order history
	orders []model.Order //slice of all past orders
}

// todo
// function that (*) returns a pointer to the DatabaseHandler struct
func NewDatabaseHandler() *DatabaseHandler {
	// Init the drinks slice with some test data
	drinks := []model.Drink{
		{ID: 1, Name: "Espresso", Price: 2.5, Description: "Strong coffee shot"},
		{ID: 2, Name: "Cappuccino", Price: 4.0, Description: "Coffee with milk foam"},
		{ID: 3, Name: "Matcha Latte", Price: 4.5, Description: "Green tea powder with steamed milk"},
		{ID: 4, Name: "Chai Latte", Price: 3.75, Description: "Spiced black tea with milk"},
		{ID: 5, Name: "Iced Tea", Price: 3.0, Description: "Chilled tea served with lemon"},
		{ID: 6, Name: "Hot Chocolate", Price: 5.0, Description: "Rich and creamy chocolate drink"},
	}

	// Init orders slice with some test data
	orders := []model.Order{
		{DrinkID: 1, CreatedAt: time.Date(2025, time.September, 2, 15, 4, 5, 0, time.UTC), Amount: 2},
		{DrinkID: 2, CreatedAt: time.Date(2025, time.September, 3, 15, 4, 25, 0, time.UTC), Amount: 3},
		{DrinkID: 3, CreatedAt: time.Date(2025, time.September, 4, 15, 24, 5, 0, time.UTC), Amount: 1},
		{DrinkID: 4, CreatedAt: time.Date(2025, time.September, 5, 19, 4, 5, 0, time.UTC), Amount: 2},
		{DrinkID: 5, CreatedAt: time.Date(2025, time.September, 6, 15, 4, 5, 0, time.UTC), Amount: 5},
		{DrinkID: 6, CreatedAt: time.Date(2025, time.September, 7, 15, 4, 5, 0, time.UTC), Amount: 8},
	}

	return &DatabaseHandler{
		drinks: drinks,
		orders: orders,
	}
}

// this defines a method on the struct
// *DatabaseHandler is the receiver // GetDrinks is the name // []model.Drink is the return type
// when called it returns a slice of drinks
func (db *DatabaseHandler) GetDrinks() []model.Drink {
	return db.drinks
}

func (db *DatabaseHandler) GetOrders() []model.Order {
	return db.orders
}

// todo
func (db *DatabaseHandler) GetTotalledOrders() map[uint64]uint64 {
	// calculate total orders
	// key = DrinkID, value = Amount of orders
	// totalledOrders map[uint64]uint64
	totalledOrders := make(map[uint64]uint64) // initialize totals map
	for _, order := range db.orders {         //loop over each order
		totalledOrders[order.DrinkID] += order.Amount //updating the map
	}
	return totalledOrders
}

// mutator method that modifies the database by adding a new order
// parameter (order *model.Order) is a pointer, passing by reference not copy
func (db *DatabaseHandler) AddOrder(order *model.Order) {
	// todo
	// add order to db.orders slice
	db.orders = append(db.orders, *order)
}
