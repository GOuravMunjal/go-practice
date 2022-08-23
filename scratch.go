package main

import (
	"errors"
	"fmt"
	"log"
)

type employee struct {
	id     string
	name   string
	salary int
}

var employees = []employee{
	{
		id:     "4432d92c-e161-4555-bc8f-4e096f1e7312",
		name:   "Shikhar",
		salary: 5000,
	},
	{
		id:     "c409fd93-cbff-42fb-a433-a9d53b1df48e",
		name:   "Salman",
		salary: 10000,
	},
	{
		id:     "42b27581-927f-46f9-ba1c-1f6e2d36ff3a",
		name:   "Aamir",
		salary: 15000,
	},
}

func getNamelyID(idToSearch string) (employee, error) {
	for _, obj := range employees {
		if obj.id == idToSearch {
			return obj, nil
		}
	}
	return employee{}, errors.New("element with given Id is not present")
}

func main() {

	result, err := getNamelyID("42b27581-927f-46f9-ba1c-1f6e2d36ff3a")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

}
