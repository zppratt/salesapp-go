package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"

	customerpb "prattlabs.com/salesapp/customer" // Import the generated protobuf package

	"google.golang.org/grpc"
)

type Person interface {
	fullName() string
}

type Salesperson struct {
	id        int
	firstName string
	lastName  string
}

func (sp *Salesperson) fullName() string {
	return sp.firstName + " " + sp.lastName
}

type Customer struct {
	id        int
	firstName string
	lastName  string
	prefs     Prefs
}

func NewCustomer(firstName string, lastName string) Customer {
	id := rand.Int()
	return Customer{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
	}
}

type Item struct {
	name        string
	description string
	department  string
}

func (i *Item) String() string {
	return i.name + "/" + i.department + "/" + i.description
}

type Prefs struct {
	colors []string
	items  []Item
}

func (p *Prefs) String() string {
	res := ""
	for i := 0; i < len(p.colors); i++ {
		res += p.colors[i] + " "
	}
	for i := 0; i < len(p.items); i++ {
		res += p.items[i].String()
	}
	return res
}

func (c *Customer) fullName() string {
	return c.firstName + " " + c.lastName
}

func (c *Customer) String() string {
	return fmt.Sprintf("%d", c.id) + ": " + c.fullName() + " - prefs: " + c.prefs.String()
}

type server struct {
	customerpb.UnimplementedCustomerServiceServer // Embed the Unimplemented
}

func (s *server) GetCustomerInfo(ctx context.Context, req *customerpb.CustomerRequest) (*customerpb.CustomerResponse, error) {
	// Simulate fetching customer information based on the customer_id from req
	customer := &customerpb.Customer{
		FirstName: "John",
		LastName:  "Doe",
	}
	return &customerpb.CustomerResponse{Customer: customer}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	customerpb.RegisterCustomerServiceServer(s, &server{})

	fmt.Println("Server listening on :50051")

	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
