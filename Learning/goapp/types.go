package main

import (
	"time"
 )
const (
	timeFmt              = "2006-01-02 15:04:05"
	customersBucket      = "Customers"
	ordersBucket         = "Orders"
	orderSummariesBucket = "OrderSummaries"
)

type Customer struct {
	Name        string
	Address     string
	City        string
	State       string
	Zip         string
	Phone       string
	CreatedDate time.Time
}

type Order struct {
	Id            string
	CustomerId    string
	SalespersonId string
	Items         []*OrderItem
	Total         float32
	Date          time.Time
}

type OrderItem struct {
	Id    string
	Title string
	Price float32
}

type OrderSummary struct {
	CustomerId string
	Summaries  []*OrderSummaryItem
}

type OrderSummaryItem struct {
	Id    string
	Total float32
	Date  time.Time
}