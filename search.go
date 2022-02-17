package main

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type Ticket struct {
	ID         string    `json:"_id"`
	CreatedAt  time.Time `json:"created_at"`
	Type       string    `json:"type"`
	Subject    string    `json:"subject"`
	AssigneeID int       `json:"assignee_id"`
	Tags       []string  `json:"tags"`
}

type User struct {
	ID        int       `json:"_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Verified  bool      `json:"verified"`
}

var tickets []Ticket
var users []User

func loadTickets() {
	ticketsJson, _ := ioutil.ReadFile("./data/tickets.json")
	usersJson, _ := ioutil.ReadFile("./data/users.json")

	_ = json.Unmarshal(ticketsJson, &tickets)
	_ = json.Unmarshal(usersJson, &users)
}

func init() {
	loadTickets()
}

func search() string {
	return "Hello World"
}
