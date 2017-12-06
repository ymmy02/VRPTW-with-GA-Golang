package node

import (
	"fmt"
)

const (
	WITHIN = iota
	EARY
	LATE
	DEPOT = iota
	CUSTOMER
)

type Node struct {
	id          int
	type_       int
	x           float64
	y           float64
	demand      float64
	readyTime   float64
	dueDate     float64
	serviceTime float64
}

func (n *Node) ID() int {
	return n.id
}

func (n *Node) Type() int {
	return n.type_
}

func (n *Node) Position() (float64, float64) {
	return n.x, n.y
}

func (n *Node) Demand() float64 {
	return n.demand
}

func (n *Node) ReadyTime() float64 {
	return n.readyTime
}

func (n *Node) DueDate() float64 {
	return n.dueDate
}

func (n *Node) ServiceTime() float64 {
	return n.serviceTime
}

func (n *Node) PrintInfo() {
	fmt.Println("id ", n.id, "type ", n.type_, "pos ", n.x, n.y,
		"dem ", n.demand, "ready ", n.readyTime,
		"due ", n.dueDate, "service ", n.serviceTime)
}

func (n *Node) IsAvailable(now float64) int {
	if now < n.readyTime {
		return EARY
	} else if now > n.dueDate {
		return LATE
	} else {
		return WITHIN
	}
}
