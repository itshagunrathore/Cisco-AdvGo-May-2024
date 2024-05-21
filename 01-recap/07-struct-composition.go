package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float64
}

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func (p *Product) ApplyDiscount(discount float64) {
	p.Cost = p.Cost * ((100 - discount) / 100)
}

// struct composition

type Dummy struct {
	Id int
}

type PerishableProduct struct {
	// Dummy
	Product
	Expiry string
}

// method overriding
func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry = %s", pp.Product.Format(), pp.Expiry)
}

func NewPerishableProduct(id int, name string, cost float64, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:   id,
			Name: name,
			Cost: cost,
		},
		Expiry: expiry,
	}
}

func main() {
	pen := Product{Id: 100, Name: "Pen", Cost: 10}
	fmt.Println("Before Applying discount :")
	fmt.Println(pen.Format())
	fmt.Println("After Applying discount :")
	pen.ApplyDiscount(10)
	fmt.Println(pen.Format())

	// composition
	/*
		milk := PerishableProduct{
			Product: Product{Id: 200, Name: "Milk", Cost: 50},
			Expiry:  "2 Days",
		}
	*/
	// instance creation using factory function
	milk := NewPerishableProduct(200, "Milk", 50, "2 Days")
	fmt.Println("Struct Composition")
	fmt.Printf("%+v\n", milk)

	// Accessing Attributes
	fmt.Println("Expiry :", milk.Expiry)
	/*
		fmt.Println("Id :", milk.Product.Id)
		fmt.Println("Name :", milk.Product.Name)
		fmt.Println("Cost :", milk.Product.Cost)
	*/
	fmt.Println("Id :", milk.Id)
	fmt.Println("Name :", milk.Name)
	fmt.Println("Cost :", milk.Cost)

	// Accessing Methods
	fmt.Println("Before Applying discount :")
	fmt.Println(milk.Format())
	fmt.Println("After Applying discount :")
	milk.ApplyDiscount(10)
	fmt.Println(milk.Format())
}
