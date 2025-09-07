package main

type car struct {
	make  string
	model int
}

type truck struct {
	car
	bedsize int
}
