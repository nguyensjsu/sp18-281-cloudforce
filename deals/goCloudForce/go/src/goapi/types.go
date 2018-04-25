package main

type Burgers struct {
	Id    int    `bson:"Id"`
	Name  string `bson:"Name"`
	Price int    `bson:"Price"`
	Path  string `bson:"Path"`
}

type Cart struct {
	Id    int64  `bson:"Id"`
	Name  string `bson:"Name"`
	Price int64  `bson:"Price"`
	Path  string `bson:"Path"`
}

type order struct {
	OrderStatus string
	Items       string
	Price       string
}

var orders map[string]order
