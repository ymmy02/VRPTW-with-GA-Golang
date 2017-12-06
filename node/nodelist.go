package node

import (
	"fmt"
	"math"
)

type NodeList struct {
	list            []*Node
	capacity        float64
	depot           *Node
	customers       []*Node
	customersIDList []int
}

func (n *NodeList) Depot() *Node {
	return n.depot
}

func (n *NodeList) NodeFromID(targetID int) *Node {
	for i, node := range n.list {
		id := node.ID()
		if id == targetID {
			return n.list[i]
		}
	}
	return nil
}

func (n *NodeList) Cusotmers() []*Node {
	return n.customers
}

func (n *NodeList) CusotmersIDList() []int {
	return n.customersIDList
}

func (n *NodeList) Position(targetID int) (float64, float64) {
	for _, node := range n.list {
		id := node.ID()
		if id == targetID {
			return node.Position()
		}
	}
	return -1.0, -1.0
}

func (n *NodeList) PrintInfo() {
	fmt.Println("========================================")
	fmt.Println("List")
	for _, node := range n.list {
		node.PrintInfo()
	}
	fmt.Println("========================================")
	fmt.Println("Depot")
	n.Depot().PrintInfo()
	fmt.Println("========================================")
	fmt.Println("Customers")
	for _, node := range n.Cusotmers() {
		node.PrintInfo()
	}
	fmt.Println("========================================")
	fmt.Println("Customers ID List")
	fmt.Println(n.CusotmersIDList())
}

func (n *NodeList) IsFeasible(route []int) bool {
	// Capacity check
	var amount float64 = 0.0
	for _, nodeID := range route {
		node := n.NodeFromID(nodeID)
		demand := node.Demand()
		amount += demand
		if amount > n.capacity {
			return false
		}
	}

	// Time Window check
	var t float64 = 0.0
	for _, nodeID := range route {
		node := n.NodeFromID(nodeID)
		readyTime := node.ReadyTime()
		dueDate := node.DueDate()
		serviceTime := node.ServiceTime()
		if t > dueDate {
			return false
		}
		t = math.Max(t, readyTime) + serviceTime
	}
	depot := n.depot
	if t > depot.DueDate() {
		return false
	}

	return true
}
