package services

import (
	"log"
	"strconv"

	"main.go/models"
)

type Order struct {
	PizzaID  int `json:"pizza_id"`
	Quantity int `json:"quantity"`
	Total    int `json:"total"`
}

type Orders []Order

func GetAllOrders(fileName string) (Orders, error) {
	var orders Orders
	// read data from csv
	records, err := models.ReadData(fileName)
	if err != nil {
		log.Println("[ERROR] Can't read orders data from csv")
		return orders, err

	}
	if len(records) == 0 {
		// log.Fatalln("Error: No orders found")
		log.Println("[WARN] No orders found")
		return orders, err
	}
	for _, record := range records {

		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Println(err)
		}
		quantity, err := strconv.Atoi(record[1])
		if err != nil {
			log.Println(err)
		}
		total_price, err := strconv.Atoi(record[2])
		if err != nil {
			log.Println(err)
		}

		order := Order{
			PizzaID:  id,
			Quantity: quantity,
			Total:    total_price,
		}
		orders = append(orders, order)
	}
	return orders, err
}

func GetOrderByID(id int) (Orders, error) {
	var orders Orders
	var os Orders

	orders, err := GetAllOrders("orders.csv")
	if err != nil {
		log.Println("[WARN] No any orders found")
		return os, err
	}

	os = orders.GetByID(id)
	if os == nil {
		log.Println("[WARN] Can't find order by PizzaID: " + strconv.Itoa(id))
		return os, err
	}
	return os, err
}

func (orders Orders) GetByID(ID int) Orders {
	var re_orders Orders
	for _, order := range orders {
		if order.PizzaID == ID {
			re_orders = append(re_orders, order)
		}
	}
	return re_orders
}

func PlaceOrder(o Order) error {
	pizzas, err := GetAllPizzas("pizzas.csv")
	if err != nil {
		log.Println("[WARN] No pizzas found, can't place order")
		return err
	}
	p, err := pizzas.FindByID(o.PizzaID)
	if err != nil {
		log.Println("[WARN] Not found this pizza, can't place order")
		return err
	}

	o.Total = p.Price * o.Quantity

	// store order data in csv
	order_new := []string{
		strconv.Itoa(o.PizzaID),
		strconv.Itoa(o.Quantity),
		strconv.Itoa(o.Total),
	}
	log.Println("[INFO] Write placed order record to csv")
	models.WriteData("orders.csv", order_new)
	return nil
}
