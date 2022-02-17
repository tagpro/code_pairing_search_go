package main

import (
	"io/ioutil"
)

var ticketsJson []byte
var usersJson []byte

func loadTickets() {
	ticketsJson, _ = ioutil.ReadFile("./data/tickets.json")
	usersJson, _ = ioutil.ReadFile("./data/users.json")
}

func init() {
	loadTickets()
}

func search() string {
	return "Hello World"
}
