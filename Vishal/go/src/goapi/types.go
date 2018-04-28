package main

type Burgers struct {
	Id    int    `bson:"Id"`
	Name  string `bson:"Name"`
	Price int    `bson:"Price"`
	Path  string `bson:"Path"`
}

type Cart struct {
	Id       int64  `bson:"Id"`
	Name     string `bson:"Name"`
	Price    int64  `bson:"Price"`
	Path     string `bson:"Path"`
	UserName string `bson:"UserName"`
}

type order struct {
	OrderStatus string `bson:"OrderStatus"`
	Items       string `bson:"Items"`
	Price       string `bson:"Price"`
	UserName    string `bson:"UserName"`
}

type orderSave struct {
	Id          string `bson:"Id"`
	OrderStatus string `bson:"OrderStatus"`
	Items       string `bson:"Items"`
	Price       string `bson:"Price"`
	UserName    string `bson:"UserName"`
}