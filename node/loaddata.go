package node

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const NODENUM = 101

//*********//
// Private //
//*********//
func fromFile(filename string, capacity int) []string {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File %s could not read: %v\n", filename, err)
		os.Exit(1)
	}

	defer f.Close()

	lines := make([]string, 0, capacity)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func trim(elements []string) []string {
	values := make([]string, 0, 7)
	for _, element := range elements {
		if element != "" {
			values = append(values, element)
		}
	}
	return values
}

func createNode(values []string, type_ int) *Node {
	id, _ := strconv.Atoi(values[0])
	x, _ := strconv.ParseFloat(values[1], 64)
	y, _ := strconv.ParseFloat(values[2], 64)
	demand, _ := strconv.ParseFloat(values[3], 64)
	readyTime, _ := strconv.ParseFloat(values[4], 64)
	dueDate, _ := strconv.ParseFloat(values[5], 64)
	serviceTime, _ := strconv.ParseFloat(values[6], 64)

	node := &Node{
		id:          id,
		type_:       type_,
		x:           x,
		y:           y,
		demand:      demand,
		readyTime:   readyTime,
		dueDate:     dueDate,
		serviceTime: serviceTime,
	}
	return node
}

//********//
// Public //
//********//
func LoadData(vehicleCapacityFilename string, nodeFilename string) *NodeList {
	vcapLines := fromFile(vehicleCapacityFilename, 1)
	vehicleCapacity, _ := strconv.ParseFloat(vcapLines[0], 64)

	nodes := &NodeList{
		list:            make([]*Node, 0, NODENUM),
		capacity:        vehicleCapacity,
		customersIDList: make([]int, 0, NODENUM-1),
	}

	nodeLines := fromFile(nodeFilename, 105)
	// Skip header i = 0
	// Depot
	line := nodeLines[1]
	elements := strings.Split(line, " ")
	values := trim(elements)
	depot := createNode(values, DEPOT)
	nodes.list = append(nodes.list, depot)
	nodes.depot = depot
	// Customers
	for i := 2; i < len(nodeLines); i++ {
		line := nodeLines[i]
		elements := strings.Split(line, " ")
		values := trim(elements)
		node := createNode(values, CUSTOMER)
		nodes.list = append(nodes.list, node)
	}

	nodes.customers = nodes.list[1:]
	for _, customer := range nodes.customers {
		id := customer.ID()
		nodes.customersIDList = append(nodes.customersIDList, id)
	}
	return nodes
}
