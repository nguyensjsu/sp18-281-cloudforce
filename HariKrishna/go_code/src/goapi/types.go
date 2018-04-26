package main

type Burgers struct {
	Id    int    `bson:"Id"`
	Name  string `bson:"Name"`
	Price int    `bson:"Price"`
	Path  string `bson:"Path"`
}
