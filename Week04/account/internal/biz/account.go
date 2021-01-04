package biz

type Account struct {
	Id int
	Name string
	Address Address
}

type Address struct {
	Id int
	City string
	Area string
}
