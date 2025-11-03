package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck not found")
)

type Truck struct {
	id string
}

func (t *Truck) LoadCargo() error {
	return ErrTruckNotFound
}

func processTruck(truck Truck) error {
	fmt.Printf("Processing truck: %s\n", truck.id)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo: %w", err)
	}
	return ErrNotImplemented
}

func main() {
	trucks := []Truck{
		{id: "Truck-1"},
		{id: "Truck-2"},
		{id: "Truck-3"},
	}

	for _, truck := range trucks {
		fmt.Printf("Truck %s arrived.\n", truck.id)
		// err := processTruck(truck)
		// if err != nil {
		// 	log.Fatalf("Error processing trucks: %s", err)
		// }

		if err := processTruck(truck); err != nil {
			if errors.Is(err, ErrNotImplemented) {
				// We do this
			}

			if errors.Is(err, ErrTruckNotFound) {
				// We do this
			}
			log.Fatalf("Error prcessing truck: %s", err)
		}
	}
}
