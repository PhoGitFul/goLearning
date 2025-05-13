package main

import "fmt"

type Laptop struct {
	brand, ram, storage string
}

// Builder method to set Brand
func (l *Laptop) SetBrand(b string) *Laptop {
	l.brand = b
	return l // Returning pointer for method chaining
}

// Builder method to set RAM
func (l *Laptop) SetRAM(r string) *Laptop {
	l.ram = r
	return l // Returning pointer for method chaining
}

// Builder method to set Storage
func (l *Laptop) SetStorage(s string) *Laptop {
	l.storage = s
	return l // Returning pointer for method chaining
}

// Display the Laptop details
func (l Laptop) ShowSpecs() {
	fmt.Printf("Laptop Specs - Brand:%s, RAM:%s, Storage:%s\n", l.brand, l.ram, l.storage)
}

func main() {
	lap := &Laptop{} // Empty laptop instance

	// Using builder pattern
	lap.SetBrand("Acer").SetRAM("8GB").SetStorage("500GB HDD").ShowSpecs()
	lap.SetBrand("Apple").SetRAM("32GB").SetStorage("1TBB SSD").ShowSpecs()
	lap.SetBrand("ASUS").SetRAM("16GB").SetStorage("200GB HDD").ShowSpecs()
	lap.SetBrand("Samsung").SetRAM("64GB").SetStorage("2TB SSD").ShowSpecs()

}
