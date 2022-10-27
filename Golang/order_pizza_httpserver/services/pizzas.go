package services

import (
	"fmt"
	"log"
	"strconv"

	"main.go/models"
)

type Pizza struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Pizzas []Pizza

func GetAllPizzas(fileName string) (Pizzas, error) {
	var pizzas Pizzas
	// read data from csv
	records, err := models.ReadData(fileName)
	if err != nil {
		log.Println("[ERROR] Can't read pizzas data from csv")
		return pizzas, err

	}
	if len(records) == 0 {
		log.Println("[WARN] No pizzas found")
		return pizzas, err
	}
	for _, record := range records {

		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Println(err)
		}
		price, err := strconv.Atoi(record[2])
		if err != nil {
			log.Println(err)
		}

		pizza := Pizza{
			ID:    id,
			Name:  record[1],
			Price: price,
		}
		pizzas = append(pizzas, pizza)
	}
	return pizzas, err
}

func (ps Pizzas) FindByID(ID int) (Pizza, error) {
	for _, pizza := range ps {
		if pizza.ID == ID {
			return pizza, nil
		}
	}

	return Pizza{}, fmt.Errorf("[WARN] Couldn't find pizza with ID: %d", ID)
}

func AddPizza(p Pizza) {

	// store pizza data in csv
	record := []string{
		strconv.Itoa(p.ID),
		p.Name,
		strconv.Itoa(p.Price),
	}
	log.Println("[INFO] Write pizza record to csv")
	models.WriteData("pizzas.csv", record)

}
