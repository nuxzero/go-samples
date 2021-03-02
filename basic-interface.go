package main

import "fmt"

type Vehicle interface {
	Run() string
	Stop() string
}

// Car fields
type Car struct {
	Name string
}

// Run the car
func (m Car) Run() string {
	return m.Name + " is running"
}

// Stop the car
func (m Car) Stop() string {
	return m.Name + " is stopped"
}

func main() {
	var c Vehicle
	c = Car{"Black"}

	fmt.Printf("Car name: %s\n", "test")

	status := c.Run()
	fmt.Println(status)

	status = c.Stop()
	fmt.Println(status)
}
