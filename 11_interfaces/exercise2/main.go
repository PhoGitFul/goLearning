package main

import "fmt"

// Define an interface
type Employee interface {
	GetSalary() float64
	GetDetails() string
}

// Permanent Employee
type Permie struct {
	name   string
	salary float64
}

// Contract Employee
type Contractor struct {
	name              string
	hours, hourlyRate float64
}

// Define the functions on each Struct
func (p Permie) GetSalary() float64 { return p.salary }

func (c Contractor) GetSalary() float64 {
	salary := c.hours * c.hourlyRate
	return salary
}

func (p Permie) GetDetails() string {
	msg1 := fmt.Sprintf("%s is a Permanent Employee.\n", p.name)
	return msg1
}

func (c Contractor) GetDetails() string {
	msg1 := fmt.Sprintf("%s is a Contract Employee.\n", c.name)
	return msg1
}

func printEmpDetails(e Employee) {
	fmt.Printf("%s", e.GetDetails())
	fmt.Printf("Their monthly salary is £%.2f\n", e.GetSalary())
	fmt.Println()
}

func main() {
	// Define a couple of new Permie Employee's along with their monthly salary
	p1 := Permie{
		name:   "Susan",
		salary: 500}

	p2 := Permie{
		name:   "Mal",
		salary: 750.32}

	// Define a couple of new Contract Employee's along with their hourly rate and hours worked
	c1 := Contractor{
		name:       "Natasha",
		hours:      20,
		hourlyRate: 45.20}

	c2 := Contractor{
		name:       "Kelvin",
		hours:      40,
		hourlyRate: 100.50}

	printEmpDetails(p1)
	printEmpDetails(p2)
	printEmpDetails(c1)
	printEmpDetails(c2)

}
